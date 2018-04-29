/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nano

import "github.com/gyrolab/gyro"

type widgets struct {
	parent  widget
	widgets []widget
}

func newWidgets(parent widget) *widgets {
	return &widgets{
		parent: parent,
	}
}

func (ws *widgets) AddWidget(gw gyro.Widget) {
	w := gw.(widget)
	//w.setParent(ws.parent) // TODO
	ws.widgets = append(ws.widgets, w)
}

func (ws *widgets) hasWidgets() bool {
	return len(ws.widgets) != 0
}

func (ws *widgets) renderWidgets(ctx *context) {
	for _, w := range ws.widgets {
		w.render(ctx)

		// Render the children widgets in a limited drawing area.
		if w.hasWidgets() {
			ctx.nc.Save()
			ctx.nc.Scissor(w.xF(), w.yF(), w.widthF(), w.heightF())
			ctx.nc.Translate(0, -20) // TODO
			w.renderWidgets(ctx)
			ctx.nc.Restore()
		}
	}
}
