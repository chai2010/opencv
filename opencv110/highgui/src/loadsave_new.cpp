// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_highgui.h"
// #include "_cgo_export.h"

//extern "C" {
	//int cvHaveImageReader_goproxy(const char* filename);
	//int cvHaveImageWriter_goproxy(const char* filename);
	//IplImage* cvLoadImage_goproxy(const char* filename, int iscolor);
	//CvMat* cvLoadImageM_goproxy(const char* filename, int iscolor);
	//int cvSaveImage_goproxy(const char* filename, const CvArr* arr);
//}

CV_IMPL
int cvHaveImageReader( const char* filename )
{
	return 0;//cvHaveImageReader_goproxy(filename);
}

CV_IMPL
int cvHaveImageWriter( const char* filename )
{
	return 0;//cvHaveImageWriter_goproxy(filename);
}

CV_IMPL
IplImage* cvLoadImage( const char* filename, int iscolor )
{
	return NULL;//cvLoadImage_goproxy(filename, iscolor);
}

CV_IMPL
CvMat* cvLoadImageM( const char* filename, int iscolor )
{
	return NULL;//cvLoadImageM_goproxy(filename, iscolor);
}

CV_IMPL
int cvSaveImage( const char* filename, const CvArr* arr )
{
	return 0;//cvSaveImage_goproxy(filename, arr);
}
