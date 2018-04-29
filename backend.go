/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nanovgo

import (
	"github.com/gyrolab/gyro"
)

type Backend struct {
	app *App // There can be only one app.
}

func New() *Backend {
	return &Backend{}
}

func (b *Backend) App() gyro.App {
	return b.app
}

func (b *Backend) NewText() gyro.Text {
	return nil
}
