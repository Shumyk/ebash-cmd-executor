package util

import (
	"log"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%v took %v", name, time.Since(start))
	}
}

func Cautiosly[T any](result T, err error) func(string) T {
	return func(action string) T {
		Panically(err, action)
		return result
	}
}

func Panically(err error, action string) {
	if err != nil {
		log.Panicf("could not %v ahh!! [%v]", action, err)
	}
}
