package ADT

type Integer interface {
	int8 | int16 | int32 | int64 | int
}

type Float interface {
	float32 | float64
}

type Numbers interface {
	Integer | Float
}

type AlNum interface {
	Integer | Float | string
}
