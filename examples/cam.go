// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"log"
	"runtime"

	opencv "github.com/chai2010/opencv.go"
)

func main() {
	runtime.LockOSThread()
	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()

	cam := opencv.NewCameraCapture(0)
	if cam == nil {
		log.Fatalf("can not open camera")
	}
	defer cam.Release()

	for i := 0; ; i++ {
		if m := cam.QueryFrame(); m != nil {
			win.ShowImage(m)
		}
		fmt.Println(i)

		if key := opencv.WaitKey(30); key == 27 {
			break
		}
	}
	fmt.Println("Done")
}
