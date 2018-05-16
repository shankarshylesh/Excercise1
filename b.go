package main
import "fmt"
func main() {

   num :=[]int{2,3,4,5,6,7,8,45,76,}
   sum :=0

   for _, num :=range num {
      sum += num
   }
   fmt.Println("total",sum)

   for i, num :=range num {
      if num == 76 {
        fmt.Println("index:",i)
      }
   }
}
