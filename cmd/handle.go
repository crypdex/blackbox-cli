package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func handle(err *error) {
	if r := recover(); r != nil {
		*err = errors.New(r.(error).Error())
		fatal(*err)
	}
}

func fatal(err error) {
	fmt.Println(aurora.Red(err.Error()))
	os.Exit(0)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
