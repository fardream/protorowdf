package protorowdf

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func getOrPanic[T any](a T, err error) T {
	orPanic(err)
	return a
}
