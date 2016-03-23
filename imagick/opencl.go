// Copyright 2016 Nicholas Saarela. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <magick/MagickCore.h>
*/
import "C"

import (
	"unsafe"
)


type ExceptionInfo struct {
	except* C.ExceptionInfo
}

// Inicializes the opencl...
func InitializeOpenCL () {
	
	C.InitImageMagickOpenCL(C.ImageMagickOpenCLMode, unsafe.Pointer, unsafe.Pointer, except)
}

}