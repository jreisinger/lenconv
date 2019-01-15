// Package lenconv performs length conversions.
package lenconv

import (
	"fmt"
)

// Type declarations.
type Inch float64
type Cm float64

// Conversion functions.
func IToC(i Inch) Cm { return Cm(i / 0.39370) }
func CToI(c Cm) Inch { return Inch(c * 0.39370) }

// String method controls how values are printed as strings.
func (c Cm) String() string {
	return fmt.Sprintf("%.2fcm", c)
}

func (i Inch) String() string {
	return fmt.Sprintf("%.2fin", i)
}
