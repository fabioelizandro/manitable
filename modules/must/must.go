package must

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func True(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func False(condition bool, msg string) {
	True(!condition, msg)
}

func Return[T any](value T, err error) T {
	NoErr(err)

	return value
}
