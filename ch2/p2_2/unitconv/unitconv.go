package unitconv

import (
	"fmt"
)

type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

//format foot
func (ft Foot) String() string {
	return fmt.Sprintf("%.2fft", ft)
}

//format meter
func (m Meter) String() string {
	return fmt.Sprintf("%.2fm", m)
}

//format pound
func (p Pound) String() string {
	return fmt.Sprintf("%.2flb", p)
}

//format kilogram
func (kg Kilogram) String() string {
	return fmt.Sprintf("%.2fkg", kg)
}

//foot to meter
func FTtoM(ft Foot) Meter {
	return Meter(ft / 3.28)
}

//meter to foot
func MtoFT(m Meter) Foot {
	return Foot(m * 3.28)
}

//pound to kilogram
func PtoKG(p Pound) Kilogram {
	return Kilogram(p * 0.45)
}

//kilogram to pound
func KGtoP(kg Kilogram) Pound {
	return Pound(kg / 0.45)
}
