// Copyright 2016 Nicholas Saarela. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <CL/cl.h>
#include <magick/MagickCore.h>
*/
import "C"
import "unsafe"

type DeviceID C.cl_device_id

// Initializes opencl...
func InitializeOpenCL () (int, int) {
	var device DeviceID
	devicePointer := unsafe.Pointer(&device)
    	var e *C.ExceptionInfo
		return int(C.InitImageMagickOpenCL(C.MAGICK_OPENCL_DEVICE_SELECT_AUTO, nil, devicePointer, e)), int(device)
}

