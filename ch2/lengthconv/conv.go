package lengthconv

func MToF(m Meter) Feet { return Feet(m * 3.281) }

func FToM(f Feet) Meter { return Meter(f / 3.281) }
