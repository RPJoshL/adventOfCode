package utils

import "git.rpjosh.de/RPJosh/go-logger"

func Panic(err error) {
	if err != nil {
		logger.Fatal("Panic: %s", err)
	}
}
func PanicTwo[T any](val T, err error) {
	Panic(err)
}
