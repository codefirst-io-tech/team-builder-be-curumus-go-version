package go101

import (
	"fmt"
)

func main() {
	abidin := [2]string{"merhaba", "dünya"}

	var abidin2 [2]string
	abidin2[0] = "hello"
	abidin2[1] = "world"

	abidin3 := [...]string{"koniçiva", "moniçiva"}

	fmt.Println(abidin)
	fmt.Println(abidin2)
	fmt.Println(abidin3)
	fmt.Println("******************")

	var numbers []int
	numbers = append(numbers, 4, 2, 6, 5, 4, 9, 8, 7)
	fmt.Println(numbers)
	fmt.Println(numbers[1:4])
	fmt.Println(numbers[:4])
	fmt.Println(numbers[:3])
	fmt.Println("******************")

	if numbers != nil {
		fmt.Println("array has values")
	}

	for i := range numbers {
		fmt.Println(numbers[i])
	}
	fmt.Println("******************")

	//*(<< üst alma ile birlikte çarpma) <<
	fmt.Println(8 << uint(3))

	//*(üst alma ile birlikte bölme işareti) >>
	fmt.Println(8 >> uint(3))

	fmt.Println("******************")

	maps := map[string]int{
		"tuzla":    40000,
		"kadıköy":  2345,
		"bağcılar": 232556,
	}
	fmt.Println(maps)

	for key, value := range maps {
		fmt.Printf("key %s : value %d\n", key, value)
	}
	fmt.Println("******************")

	fmt.Println(float64(1 << 4))
	fmt.Println(float64(1 * 3))

}
