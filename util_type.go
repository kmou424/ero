package ero

func as[T any](err error) (e T, ok bool) {
	e, ok = err.(T)
	return
}

func mustAs[T any](err error) (e T) {
	e, ok := as[T](err)
	if !ok {
		panic("ero can only handle errors created by itself")
	}
	return e
}
