// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"log"
	"os"

	opencv "github.com/chai2010/opencv.go"
)

func main() {
	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()

	cam := opencv.NewCameraCapture(0)
	if cam == nil {
		log.Fatalf("can not open camera")
	}
	defer cam.Release()

	win.CreateTrackbar("Thresh", 1, 100, func(pos int, param ...interface{}) {
		for {
			if cam.GrabFrame() {
				img := cam.RetrieveFrame()
				if img != nil {
					ProcessImage(img, win, pos)
				} else {
					log.Println("pos(%d) is nil", pos)
				}
			}

			if key := opencv.WaitKey(10); key == 27 {
				os.Exit(0)
			}
		}
	})
	opencv.WaitKey(0)
}

func ProcessImage(img *opencv.IplImage, win *opencv.Window, pos int) error {
	w := img.GetWidth()
	h := img.GetHeight()

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
	defer cedge.Release()

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()
	defer edge.Release()

	opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY)

	opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
	opencv.Not(gray, edge)

	// Run the edge detector on grayscale
	opencv.Canny(gray, edge, float64(pos), float64(pos*3), 3)

	opencv.Zero(cedge)
	// copy edge points
	opencv.Copy(img, cedge, edge)

	win.ShowImage(cedge)
	return nil
}
