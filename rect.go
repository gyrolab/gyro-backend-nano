/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nano

import (
	"github.com/gyrolab/gyro"
	"github.com/gyrolab/nanovgo"
)

type Rect struct {
	*widgets

	x      float32
	y      float32
	width  float32
	height float32
	right  float32
	bottom float32

	radius float32
	border float32

	color  gyro.Color
	colorN nanovgo.Color

	borderColor  gyro.Color
	borderColorN nanovgo.Color
}

func (b *Backend) NewRect() gyro.Rect {
	r := &Rect{
		color:       gyro.RGBA(0, 0, 0, 255),
		borderColor: gyro.RGBA(0, 0, 0, 255),
	}
	r.widgets = newWidgets(r)
	r.colorN = colorToNano(r.color)
	r.borderColorN = colorToNano(r.borderColor)
	return r
}

func (r *Rect) X() int {
	return int(r.x)
}

func (r *Rect) SetX(x int) {
	r.x = float32(x)
	r.right = r.x + r.width
}

func (r *Rect) Y() int {
	return int(r.y)
}

func (r *Rect) SetY(y int) {
	r.y = float32(y)
	r.bottom = r.y + r.height
}

func (r *Rect) Width() int {
	return int(r.width)
}

func (r *Rect) SetWidth(w int) {
	r.width = float32(w)
	r.right = r.x + r.width
}

func (r *Rect) Height() int {
	return int(r.height)
}

func (r *Rect) SetHeight(h int) {
	r.height = float32(h)
	r.bottom = r.y + r.height
}

func (r *Rect) Right() int {
	return int(r.right)
}

func (r *Rect) Bottom() int {
	return int(r.bottom)
}

func (r *Rect) Radius() int {
	return int(r.radius)
}

func (r *Rect) SetRadius(rr int) {
	r.radius = float32(rr)
}

func (r *Rect) Border() int {
	return int(r.border)
}

func (r *Rect) SetBorder(b int) {
	r.border = float32(b)
}

func (r *Rect) Color() gyro.Color {
	return r.color
}

func (r *Rect) SetColor(c gyro.Color) {
	r.color = c
	r.colorN = colorToNano(c)
}

func (r *Rect) BorderColor() gyro.Color {
	return r.borderColor
}

func (r *Rect) SetBorderColor(c gyro.Color) {
	r.borderColor = c
	r.borderColorN = colorToNano(c)
}

//###############//
//### Private ###//
//###############//

func (r *Rect) xF() float32 {
	return r.x
}

func (r *Rect) yF() float32 {
	return r.y
}

func (r *Rect) widthF() float32 {
	return r.width
}

func (r *Rect) heightF() float32 {
	return r.height
}

func (r *Rect) render(ctx *context) {
	// TODO: add box shadow options.
	/*
	 // Drop shadow
	 shadowPaint := nanovgo.BoxGradient(x, y+2, w, h, cornerRadius*2, 10, nanovgo.RGBA(0, 0, 0, 128), nanovgo.RGBA(0, 0, 0, 0))
	 ctx.BeginPath()
	 ctx.Rect(x-10, y-10, w+20, h+30)
	 ctx.RoundedRect(x, y, w, h, cornerRadius)
	 ctx.PathWinding(nanovgo.Hole)
	 ctx.SetFillPaint(shadowPaint)
	 ctx.Fill()
	*/

	// Draw the border if set.
	if r.border > 0 {
		// TODO: Maybe there is a better way than filling a complete rect.
		ctx.nc.BeginPath()
		ctx.nc.RoundedRect(
			r.x-r.border, r.y-r.border,
			r.width+2*r.border, r.height+2*r.border,
			r.radius,
		)
		ctx.nc.SetFillColor(r.borderColorN)
		ctx.nc.Fill()
	}

	ctx.nc.BeginPath()
	ctx.nc.RoundedRect(r.x, r.y, r.width, r.height, r.radius)
	ctx.nc.SetFillColor(r.colorN)
	ctx.nc.Fill()
}
