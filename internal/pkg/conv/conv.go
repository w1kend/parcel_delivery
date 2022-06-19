package conv

func Pointer[V any](val V) *V {
	return &val
}
