package main

import "fmt"
import "strings"
import "strconv"

func printRec(arr [9]string,next string) int {
  fmt.Println(arr)


  var flag bool
  flag=false
  var i int
  i=0
  var countx,counto,count int
  countx =0
  counto =0
  count=0
  for i<9{
    if(arr[i][:1]=="X"){
      countx=countx+1
    }
    if(arr[i][:1]=="O"){
      counto=counto+1
    }
    if(arr[i]=="*"){
      flag=true
    }
    i=i+1
  }
  if(flag){
    i=0
    for i<9{
      if(arr[i]=="*"){

      if(next=="O"){
      arr[i]=next

      var g=[]string{arr[i],strconv.Itoa(counto+1)}
      arr[i]=strings.Join(g,"")

      count=count + printRec(arr,"X")
      }
       if(next=="X"){
      arr[i]=next
      var g=[]string{arr[i],strconv.Itoa(countx+1)}
      arr[i]=strings.Join(g,"")
      count = count + printRec(arr,"O")
      }
      arr[i]="*"
      }
       i=i+1
  }
  }
  return count+1
}

func main() {
  var i int
  i=0
  var arr =[9]string{"*","*","*","*","*","*","*","*","*"}
  var count int
  count=0
  for i<9{
    arr[i]="X1"
    count = count + printRec(arr,"O")
    arr[i]="*"
    i=i+1
  }
  fmt.Println(count)
}
