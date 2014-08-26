// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo CFLAGS: -I./opencv110/highgui/include

#include <highgui.h>
*/
import "C"

///export cvHaveImageReader_goproxy
func cvHaveImageReader_goproxy(filename *C.schar) C.int {
	println("cvHaveImageReader_goproxy: todo")
	return 0
}

///export cvHaveImageWriter_goproxy
func cvHaveImageWriter_goproxy(filename *C.schar) C.int {
	println("cvHaveImageWriter_goproxy: todo")
	return 0
}

///export cvLoadImage_goproxy
func cvLoadImage_goproxy(filename *C.schar, iscolor C.int) *C.IplImage {
	println("cvLoadImage_goproxy: todo")
	return nil
}

///export cvLoadImageM_goproxy
func cvLoadImageM_goproxy(filename *C.schar, iscolor C.int) *C.CvMat {
	println("cvLoadImageM_goproxy: todo")
	return nil
}

///export cvSaveImage_goproxy
func cvSaveImage_goproxy(filename *C.schar, arr *C.CvArr) C.int {
	println("cvSaveImage_goproxy: todo")
	return 0
}
