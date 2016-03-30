package imagick

import "testing"

func TestInitOpenCL(t *testint.T){
    success := InitializeOpenCL() 
    if !success {
        t.Fatal("Error initializing OpenCL")    
    }
}
