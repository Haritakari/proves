package main

import "fmt"

const HELLO string = "hola %s %s \n"

func main() {
	var name string
	lastName := "sanchez"
	f32, f64 := getFloat()
	name = getName(name)
	fmt.Printf(HELLO, name, lastName)
	fmt.Println(getUnicode())
	fmt.Println(f32, f64)
	fmt.Println(getArray())
	fmt.Println(getSly())
	fmt.Println(associative())
	fmt.Println(ifTest(0))
	fmt.Println(ifTest(4))
	forTest()
	forTest2()
	infinite()
}

func getName(name string) string {
	fmt.Println("ingresa tu name")
	fmt.Scanf("%s", &name)

	return name
}

func getFloat() (float32, float64) {

	return float32(0.1), float64(float32(0.1))
}

func getUnicode() string {

	return "ャラク"
}

func getArray() [2]string {
	var arr [2]string
	arr[0] = "stringu"
	arr[1] = "stringu2"

	return arr
}

func getSly() []string {
	var sly []string
	sly = append(sly, "hola", "cosas")

	return sly
}

func associative() map[string]string {
	elements := make(map[string]string)

	elements["la"] = "asocia"
	elements["lo"] = "tambien"

	return elements

}

func ifTest(number int) int {
	if number != 0 {
		return 100
	}

	return number
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forTest2() {
	i := 3

	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func infinite() {
	condition := 6

	for {
		fmt.Println(condition)

		if condition == 3 {
			break
		}
		condition--
	}
}
