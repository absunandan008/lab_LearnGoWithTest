package main

import (
	"os"
	"time"

	"github.com/absunandan008/lab_LearnGoWithTest/math/vFinal/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
