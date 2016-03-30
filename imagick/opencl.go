// Copyright 2016 Nicholas Saarela. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <magick/MagickCore.h>
*/
import "C"

// Initializes opencl...
func InitializeOpenCL () C.MagickBooleanType {
    var e *C.ExceptionInfo
	return C.InitImageMagickOpenCL(C.MAGICK_OPENCL_DEVICE_SELECT_AUTO, nil, nil, e)
}

