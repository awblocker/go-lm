
package lm

// #cgo CFLAGS: -O2 -Wall
// #cgo LDFLAGS: -lf77blas -llapack -latlas -lm -lgfortran
// #include "wls.h"
import "C"
import "unsafe"

// X must be in column-major order
func Wls(X []float64, n int, p int, y []float64, w []float64) (
    []float64, int) {
    // Allocate memory for intermediate objects
    XTX := make([]float64, p*p)
    sqw := make([]float64, n)
    sqwX := make([]float64, n*p)
    sqwy := make([]float64, n)
    coef := make([]float64, p)

    // Call C function for WLS fit
    status := C.wls(
        (*C.double) (unsafe.Pointer(&X[0])), C.int(n), C.int(p),
        (*C.double) (unsafe.Pointer(&y[0])),
        (*C.double) (unsafe.Pointer(&w[0])),
        (*C.double) (unsafe.Pointer(&XTX[0])),
        (*C.double) (unsafe.Pointer(&sqw[0])),
        (*C.double) (unsafe.Pointer(&sqwX[0])),
        (*C.double) (unsafe.Pointer(&sqwy[0])),
        (*C.double) (unsafe.Pointer(&coef[0])))

    return coef, (int) (status)
}

