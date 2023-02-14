package read

import (
	"fmt"
)

func read() {
	var val1, val2 string
	var oprt string
	//var age int
	fmt.Print("Введите (число операнд число): ")
	fmt.Scan(&val1, &oprt, &val2)
	fmt.Println(val1, oprt, val2)
}
