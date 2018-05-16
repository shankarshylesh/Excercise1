package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []string{"yak", "apple", "kannur", "calicut"}
    sort.Sort(sort.StringSlice(s))
    fmt.Println(s)

    num := []int{-15,898,45,87,5,6,-85}
    sort.Sort(sort.IntSlice(num))
    fmt.Println(num)
}
