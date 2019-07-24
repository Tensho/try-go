package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string { return fmt.Sprintf("%gf", f) }

// MToF converts a Meter to Feet.
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// MToF converts a Meter to Feet.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }
