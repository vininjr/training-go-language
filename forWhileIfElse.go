package main

import "fmt"

func main(){
	for_()
	while_()
	if_else_(0)
}

func for_() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while_() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func if_else_(x int) string {
	if x < 0 {
		return "negativo"
	} else {
		return "positivo"
	}
}