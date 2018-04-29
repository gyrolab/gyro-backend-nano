/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nanovgo

import (
	"github.com/gyrolab/gyro"
	"github.com/gyrolab/nanovgo"
)

func colorToNano(c gyro.Color) nanovgo.Color {
	return nanovgo.RGBA(c.R, c.G, c.B, c.A)
}
