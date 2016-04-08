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
type MagickBool C.MagickBooleanType

const (

MagickTrue MagickBool = C.MagickTrue
MagickFalse MagickBool = C.MagickFalse
	
)

// Initializes opencl...
func InitializeOpenCL () (MagickBool, DeviceID) {
	var device DeviceID
	devicePointer := unsafe.Pointer(&device)
    	var e *C.ExceptionInfo
		return C.InitImageMagickOpenCL(C.MAGICK_OPENCL_DEVICE_SELECT_AUTO, nil, devicePointer, e), device
}

func (mw *MagickWand) AccelerateResizeImage (width uint, height uint, filter FIlterType) error {
		var e *C.ExceptionInfo
		ok := C.AccelerateResizeImage(mw.GetImageFromMagickWand(), C.size_t(cols), C.size_t(rows), C.FilterTypes(filter), e)
		return mw.getLastErrorIfFailed(ok)
}
