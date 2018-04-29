/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nano

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/gyrolab/gyro"

	"github.com/desertbit/closer"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
)

const (
	FPS = time.Second / 30 // TODO: make this configurable through NewApp(opts)
)

type App struct {
	*closer.Closer

	mutex      sync.Mutex
	window     *Window // TODO: Currenlty only one window is supported.
	runStopped chan struct{}
	onInitFunc gyro.AppInitFunc
}

func (b *Backend) NewApp(f gyro.AppInitFunc) (gyro.App, error) {
	a := &App{
		runStopped: make(chan struct{}),
		onInitFunc: f,
	}
	a.Closer = closer.New(a.onClose)

	// Set the app to the global variable.
	b.app = a
	return a, nil
}

func (a *App) onClose() error {
	// Wait for the run loop to exit.
	select {
	case <-a.runStopped:
	case <-time.After(2 * time.Second):
	}

	if a.window != nil {
		a.window.Close()
	}

	glfw.Terminate()
	return nil
}

func (a *App) CloseChan() <-chan struct{} {
	return a.Closer.CloseChan
}

// Sync calls the function in a locked context and
// synchronizes with the main app rendering routine.
func (a *App) Sync(f func()) {
	a.mutex.Lock()
	defer a.mutex.Unlock() // Use defer for panics.

	f()
}

// Lock the app rendering routine.
// Prefer to use Sync.
func (a *App) Lock() {
	a.mutex.Lock()
}

// Unlock the app rendering routine and continue rendering.
// Prefer to use Sync.
func (a *App) Unlock() {
	a.mutex.Unlock()
}

func (a *App) Run() (err error) {
	// Lock to OS threadh for GLFW & OpenGl.
	runtime.LockOSThread()

	defer func() {
		// Recover panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}

		close(a.runStopped)
		a.Close()

		// If the close method was called from another goroutine,
		// then give a little time for cleanup.
		// This will occurr when main() exits before the cleanup is done.
		time.Sleep(100 * time.Millisecond)
	}()

	// Initialize GLFW.
	err = glfw.Init(gl.ContextWatcher)
	if err != nil {
		return err
	}

	// Call the init method.
	if a.onInitFunc == nil {
		return fmt.Errorf("app init function is not set")
	}
	err = a.onInitFunc()
	if err != nil {
		return err
	}

	w := a.window
	if w == nil {
		return fmt.Errorf("no window set")
	}

	fpsTicker := time.NewTicker(FPS)
	defer fpsTicker.Stop()

	ctx := newContext(w.nc)
	closeChan := a.CloseChan()

	for {
		select {
		case <-closeChan:
			return nil

		case <-fpsTicker.C:
			if w.win.ShouldClose() {
				return nil
			}

			// Process all events.
			glfw.PollEvents()

			// TODO: Lock
			// TODO: color to RGBA

			// TODO: don't do this every loop.
			fbWidth, fbHeight := w.win.GetFramebufferSize()
			winWidth, winHeight := w.win.GetSize()
			//mx, my := w.win.GetCursorPos()
			pixelRatio := float32(fbWidth) / float32(winWidth)

			// Ensure thread-safety.
			a.mutex.Lock()

			gl.Viewport(0, 0, fbWidth, fbHeight)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
			gl.ClearColor(w.bgColorR, w.bgColorG, w.bgColorB, w.bgColorA)
			gl.Enable(gl.BLEND)
			gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
			gl.Enable(gl.CULL_FACE)
			gl.Disable(gl.DEPTH_TEST)

			w.nc.BeginFrame(winWidth, winHeight, pixelRatio)
			w.renderWidgets(ctx)
			w.nc.EndFrame()

			// Release the lock again after one render step.
			a.mutex.Unlock()

			gl.Enable(gl.DEPTH_TEST)
			w.win.SwapBuffers()
		}
	}
}
