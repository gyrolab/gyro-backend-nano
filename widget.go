/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nanovgo

import "github.com/gyrolab/gyro"

type widget interface {
	gyro.Widget

	setParent(p widget)

	render(ctx *context)
	renderWidgets(ctx *context)
	hasWidgets() bool

	xF() float32
	yF() float32
	widthF() float32
	heightF() float32
}
