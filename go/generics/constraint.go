package after

type MyConstraint interface {
	int | int8 | int16 | int32 | int64
}
