package main

import "fmt"

func main() {
  var n, small, big int
  x := []int{
  1,2,3,78,67,5,7,89,65,78,87,
  }

  for _,i:=range x {
    if i>n {
      fmt.Println(i,">",n)
      n = i
      big = n
    } else {
      fmt.Println(i,"<",n)
    }
  }

  fmt.Println("The biggest number is ", big)
  for _,i:=range x {
    if i>n {
      fmt.Println(i,">",n)
    } else {
      fmt.Println(i,"<",n)
      n = i
      small = n
    }
  }
  fmt.Println("The smallest number is ", small)
}
