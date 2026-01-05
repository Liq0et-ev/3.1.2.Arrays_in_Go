// 241RDB316
// Vladislav Ebert
package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("\nFirst sequence of numbers:")
	FirstSequence := EnterMas()

	fmt.Println("\nSecond sequence of numbers:")
	SecondSequence := EnterMas()

	fmt.Println("\nResult:")
	UniqExit(FirstSequence, SecondSequence)
}

func EnterMas() []int {
	var numbers []int
	var EnterNumbers string

	for {
		fmt.Print("Eneter number: ")
		fmt.Scan(&EnterNumbers)

		numberss, Mistake := strconv.Atoi(EnterNumbers)

		if Mistake != nil {
			fmt.Println("Pu pu pu. Inccorect input")
			continue
		}

		numbers = append(numbers, numberss)

		if numberss == 0 {
			break
		}
	}

	return numbers
}

func UniqExit(FirstSequence, SecondSequence []int) {
	var found bool

	for i := 0; i < len(FirstSequence)-1; i++ {
		numberss := FirstSequence[i]
		var InSecond bool

		for j := 0; j < len(SecondSequence); j++ {
			if SecondSequence[j] == numberss {
				InSecond = true
				break
			}
		}

		if !InSecond {
			fmt.Print(numberss, " ")
			found = true
		}
	}

	if !found {
		fmt.Println("Not found")
	} else {
		fmt.Println()
	}
}
