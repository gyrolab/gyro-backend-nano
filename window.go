/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nanovgo

import (
	"github.com/desertbit/closer"
	"github.com/gyrolab/gyro"
	"github.com/gyrolab/nanovgo"

	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
)

type Window struct {
	*closer.Closer
	*widgets

	//ctx *Context
	win *glfw.Window
	nc  *nanovgo.Context

	title string

	bgColor  gyro.Color
	bgColorR float32
	bgColorG float32
	bgColorB float32
	bgColorA float32
}

func (b *Backend) NewWindow(
	title string,
	width, height int,
) (gyro.Window, error) {
	// MSAA: multisample antialiasing.
	glfw.WindowHint(glfw.Samples, 4)

	//glfw.WindowHint(glfw.Visible, 0) // TODO
	glfw.WindowHint(glfw.Resizable, 1)

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	win.MakeContextCurrent()

	width, height = win.GetSize()
	gl.Viewport(0, 0, width, height)

	nc, err := nanovgo.NewContext(nanovgo.AntiAlias | nanovgo.StencilStrokes)
	if err != nil {
		return nil, err
	}

	w := &Window{
		win:   win,
		nc:    nc,
		title: title,
	}
	w.widgets = newWidgets(w)
	w.Closer = closer.New(w.onClose)

	// Register the window to the app.
	b.app.window = w

	return w, nil
}

func (w *Window) onClose() error {
	w.nc.Delete()
	w.win.Destroy()
	return nil
}

func (w *Window) CloseChan() <-chan struct{} {
	return w.Closer.CloseChan
}
