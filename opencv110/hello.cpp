// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdio.h>
#include <highgui.h>

int main(int argc, char* argv[]) {
	const char* filename = "../testdata/lena.jpg";
	if(argc >= 2) {
		filename = argv[1];
	}

	IplImage* img = cvLoadImage(filename, 0);
	if(img == NULL) {
		printf("Load %s failed!\n", filename);
		return -1;
	}

	const char* winName = "Go-OpenCV";
	cvNamedWindow(winName, 1);
	cvShowImage(winName, img);
	cvWaitKey(0);

	return 0;
}
