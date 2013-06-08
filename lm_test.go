
package lm

import "testing"
import "fmt"
import "math"

func TestWls(t *testing.T) {
    const n, p = 10, 2
    // Testing with y = x
    X := make([]float64, n*p)
    y := make([]float64, n)
    w := make([]float64, n)

    // Fill X matrix and simulate y's
    for i := 0; i < n; i++ {
        X[i] = 1.
        X[i + n] = (float64) (i)
        y[i] = X[i + n]
        w[i] = 1.
    }

    // Run regression
    coef, status := Wls(X, n, p, y, w)

    // Print results
    fmt.Printf("Status: %d\n", status)
    fmt.Printf("Beta hat: %v\n", coef)

    l2Error := math.Sqrt(math.Pow(coef[0], 2) + math.Pow(coef[1]-1, 2))
    maxError := math.Sqrt(2)*math.Sqrt(math.Nextafter(1., 2.) - 1.)
    if status > 0 || l2Error > maxError {
        t.Errorf("Status %d\tL2 error %g > %g", status, l2Error, maxError)
    } else {
        fmt.Printf("L2 error %g < %g\n", l2Error, maxError)
    }
}

