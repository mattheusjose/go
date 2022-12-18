package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileName := "calculator.go"
	CreateFile(fileName)
	ReadFile(fileName)
	OpenFile(fileName)
	DeleteFile(fileName)
}

func ReadFile(fileName string) {

	bytes, error := os.ReadFile(fileName)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(bytes))
}

func OpenFile(fileName string) {

	file, error := os.Open(fileName)
	defer file.Close()

	if error != nil {
		panic(error)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 200)

	for {
		lastPosition, error := reader.Read(buffer)

		if error != nil {
			break
		}

		fmt.Printf("%v\n", string(buffer[:lastPosition]))
	}
}

func DeleteFile(fileName string) {

	error := os.Remove(fileName)

	if error != nil {
		panic(error)
	}
}

func CreateFile(fileName string) {

	file, error := os.Create(fileName)
	defer file.Close()

	if error != nil {
		panic(error)
	}

	file.Write([]byte(`
		package main

		import "fmt"

		type Number interface {
			~int | float32 | float64
		}

		func Calculator(action string, a int, b int) int {

			functions := map[string]func(a, b int) int{
				"SUM": func(a, b int) int {
					return a + b
				},
				"SUB": func(a, b int) int {
					return a - b 
				},
			}
			
			operation := functions[action]
			return operation(a, b)
		}

		func main() {

			var (
				a int = 10
				b int = 10
			)

			fmt.Println(Calculator("SUM", a, b))
			fmt.Println(Calculator("SUB", a, b))
		}
	`))

}
