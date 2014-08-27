// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo CFLAGS: -I./opencv110/cxcore/include

#include <cxcore.h>
*/
import "C"
import (
	"image"
	"image/color"
)

func (p *IplImage) ColorModel() color.Model {
	switch p.GetChannels() {
	case 1:
	case 3:
	case 4:
	}
	return nil
}

func (p *IplImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.GetWidth(), p.GetHeight())
}

func (p *IplImage) At(x, y int) color.Color {
	return nil
}
