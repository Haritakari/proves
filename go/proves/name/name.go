package name

import "fmt"

type geeneral interface {
	FunctionTest()
}
type platzi struct {
	Name   string
	Slug   string
	Skills []string
}

func (platzi platzi) FunctionTest() {
	fmt.Printf("Noooooo %s \n", platzi.Name)
}

type platziA struct {
	platzi
}

func (platziA platziA) FunctionTest() {
	fmt.Printf("las %s \n", platziA.Name)
}

func finalprocess() {
	fmt.Println("Acabado Todo En Name")
}

func GetName(name string) int {
	defer finalprocess()

	mapTest := make(map[string]int)

	mapTest["hola"] = 3
	mapTest["adios"] = 2

	fmt.Println("probando package name")

	/*platzi := platzi{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}

	fmt.Println(platzi)*/

	platzi1 := new(platzi)

	platzi1.Name = "Otra"
	platzi1.Slug = "yes"
	platzi1.Skills = []string{"no", "viene"}

	fmt.Println(platzi1)

	platziA2 := new(platziA)
	platziA2.Name = "esta"
	platziA2.Slug = "no"
	platziA2.Skills = []string{"cosa", "va"}

	TryInterface(platzi1)
	TryInterface(platzi1)
	TryInterface(platziA2)
	TryInterface(platziA2)

	return mapTest[name]
}

func TryInterface(gener geeneral) {
	geeneral.FunctionTest(gener)
}
