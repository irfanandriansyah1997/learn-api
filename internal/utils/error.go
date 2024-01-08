//nolint:typecheck // disable lint since NewNotFoundError still on same package
package utils

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNotFoundError(err error) {
	if err != nil {
		panic(NewNotFoundError(err.Error()))
	}
}
