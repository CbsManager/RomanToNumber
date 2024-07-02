package main

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

func RomanToNumberV2(x string) int {
	curr := GetValue(x[0])
	sum := 0
	for i:= 1 ; i < len(x) ; i++{
		next := GetValue(x[i])
		if curr < next {
			sum -= curr
		}else{
			sum += curr
		}
		curr = next
	}
	sum += curr
	return sum	
}

func GetValue(s byte) int{
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}