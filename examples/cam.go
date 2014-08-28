// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"flag"
	"fmt"
	"log"

	opencv "github.com/chai2010/opencv.go"
)

var (
	flagCamNum = flag.Int("cam-num", 1, "set camera num")
)

func main() {
	flag.Parse()

	cameras := make([]*opencv.Capture, *flagCamNum)
	for i := 0; i < len(cameras); i++ {
		cameras[i] = opencv.NewCameraCapture(i)
		if cameras[i] == nil {
			log.Fatalf("can not open camera %d", i)
		}
		defer cameras[i].Release()
	}

	windows := make([]*opencv.Window, *flagCamNum)
	for i := 0; i < len(windows); i++ {
		windows[i] = opencv.NewWindow(fmt.Sprintf("Camera: %d", i))
		defer windows[i].Destroy()
	}

	for i := 0; ; i++ {
		for idx := 0; idx < *flagCamNum; idx++ {
			if m := cameras[idx].QueryFrame(); m != nil {
				fmt.Printf("camera(%d): Frame %d\n", idx, i)
				windows[idx].ShowImage(m)
			}
		}
		if key := opencv.WaitKey(30); key == 27 {
			break
		}
	}
	fmt.Println("Done")
}
