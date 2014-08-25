// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// https://github.com/Itseez/opencv/tree/master/3rdparty/ffmpeg
// https://github.com/Itseez/opencv/blob/master/3rdparty/ffmpeg/opencv_ffmpeg.dll?raw=true
// https://github.com/Itseez/opencv/blob/master/3rdparty/ffmpeg/opencv_ffmpeg_64.dll?raw=true
const baseURL = "https://github.com/Itseez/opencv/blob/master/3rdparty/ffmpeg/"

func main() {
	var err error

	if err = GetFfmpegDll("opencv_ffmpeg.dll"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download opencv_ffmpeg.dll ok")

	if err = GetFfmpegDll("opencv_ffmpeg_64.dll"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download opencv_ffmpeg_64.dll ok")

	fmt.Println("Done")
}

func GetFfmpegDll(filename string) (errRet error) {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create %s: %s", filename, err)
	}
	defer f.Close()
	defer func() {
		if errRet != nil {
			os.Remove(filename)
		}
	}()
	resp, err := http.Get(baseURL + filename + "?raw=true")
	if err != nil {
		return fmt.Errorf("failed to download %s: %s", baseURL+filename, err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write %s: %s", filename, err)
	}
	return nil
}
