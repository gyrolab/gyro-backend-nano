/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package nano

import "github.com/gyrolab/nanovgo"

type context struct {
	nc *nanovgo.Context
}

func newContext(nc *nanovgo.Context) *context {
	return &context{
		nc: nc,
	}
}
