package main

import (
	"fmt"
	"strconv"
)

func main() {

	result := Sum(1, 2, 3)
	fmt.Println(result)

	result2 := Sum2(1, "2", 3)
	fmt.Println(result2)

	strSlc := []interface{}{"10","20",30}

	result3 := Sum2(strSlc...)
	fmt.Println(result3)
}


func Sum3(nums ...interface{}) int64 {
	var res int64
	for _, n := range nums {
		strval := fmt.Sprintf(`%v`,n)
		i,_:=strconv.ParseInt(strval, 10, 64)
		res += i
	}
	return res
}
func Sum2(nums ...interface{}) int64 {
	var res int64
	for _, n := range nums {
		strval := fmt.Sprintf(`%v`,n)
		i,_:=strconv.ParseInt(strval, 10, 64)
		res += i
	}
	return res
}
func Sum(nums ...int) int {
	var res int
	res = 0
	for _, n := range nums {
		res += n
	}
	return res
}