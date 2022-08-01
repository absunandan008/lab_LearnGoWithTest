package main

import (
	"os"
	"time"

	"github.com/absunandan008/lab_LearnGoWithTest/math/v6/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
