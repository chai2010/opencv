# Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(OpenCV)

#------------------------------------------------------------------------------

set(CXCORE_SRC)
set(CV_SRC)
set(CVAUX_SRC)
set(ML_SRC)
set(HIGHGUI_SRC)

#------------------------------------------------------------------------------

add_library(OpenCVLib SHARED
	${CXCORE_SRC}
	${CV_SRC}
	${CVAUX_SRC}
	${ML_SRC}
	${HIGHGUI_SRC}
	./opencv.def
)

set_target_properties(OpenCVLib
	PROPERTIES OUTPUT_NAME "opencv.go"
)

target_link_libraries(OpenCVLib
	Ole32 Strmiids OleAut32 Uuid
)

install(TARGETS OpenCVLib
	RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

add_executable(OpenCVApp
	./hello.c
)

set_target_properties(OpenCVApp
	PROPERTIES OUTPUT_NAME "OpenCVApp"
)

target_link_libraries(OpenCVApp
	OpenCVLib
)

install(TARGETS OpenCVApp
	RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
