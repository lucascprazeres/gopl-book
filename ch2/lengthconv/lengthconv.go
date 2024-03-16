// Pacote Lengthconv realiza conversões entre metros e pés
package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (f Feet) String() string { return fmt.Sprintf("%g'", f) }
