package main

import (
	"fmt"
	"sort"
	)

func main() {

    m := make(map[string] int)

    m["k1"] = 7
    m["k2"] = 13
    m["k3"] = 15
    m["k4"] = 13

    fmt.Println("map:", m)

    //delete(m, "k2")
    //fmt.Println("map:", m)

    values := make([]int, 0, len(m))

    for _,value := range m {
        values = append(values, value)
    }

    sort.Ints(values)
    fmt.Println(values)

}

