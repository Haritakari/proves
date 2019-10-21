package main

import (
	"errors"
	"fmt"
	"time"

	"./name"
)

func main() {

	number, error := testErrors(42)
	fmt.Println(error)
	if error != nil {
		fmt.Println("failed: ", error)
	} else {
		fmt.Println("Goooood: ", number)
	}
	numberName := name.GetName("hola")
	fmt.Println(numberName)

	a := 100

	var b *int

	b = &a

	fmt.Println(a, &a)
	fmt.Println(b, *b)

	*b = 45

	fmt.Println(a, &a)
	fmt.Println(b, *b)

	a = 38

	fmt.Println(a, &a)
	fmt.Println(b, *b)

	go testGorutinesFor(500, "1")
	go testGorutinesFor(600, "-------------------2")
	time.Sleep(1000 * time.Millisecond)
}

func testGorutines(index int, text string) {
	fmt.Println("valor : ", index, "del: ", text)
}

func testGorutinesFor(number int, text string) {
	for i := 0; i < number; i++ {
		testGorutines(i, text)
	}
}

func testErrors(number int) (int, error) {
	if number == 42 {
		return -1, errors.New("cant work with 42")
	}

	return number, nil
}
