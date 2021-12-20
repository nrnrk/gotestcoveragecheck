package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(feeling(r.Int() % 3))
}

func feeling(typ int) string {
	switch typ {
	case 0:
		return `so-so`
	case 1:
		return `fine`
	case 2:
		return `great`
	case 3:
		return `super`
	default:
		panic(errors.New(`not supported type`))
	}
}
