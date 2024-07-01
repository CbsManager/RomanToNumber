package roman

import "fmt"

func main(){
	fmt.Println(RomanToNumber("XX"))
}

func RomanToNumber(x string) int {
	valueMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	value := 0

	for i := 0; i < len(x); i++ {
		if i+1 < len(x) && valueMap[x[i+1]] > valueMap[x[i]] {
			value += (valueMap[x[i+1]] - valueMap[x[i]])
			i++
		} else {
			value += valueMap[x[i]]
		}
	}
	return value
}
