package main

import "fmt"
import "bufio"
import "os"

func main(){
 var mat =[5][5]string{{"A","B","C","D","E"},{"F","G","H","I","J"},{"K","L","M","N","O"},
           {"P","Q","R","S","T"},{"U","V","W","X","Y"}}

  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter current :")
  scanner.Scan()
  curr:=scanner.Text()
  fmt.Println("Block posn :")
  scanner.Scan()
  bloc:=scanner.Text()

  var mov int
  mov=0
  var ax,bx,ay,by int
  ax=-2
  ay=-2
  bx=6
  by=6

  for p:=0;p<5;p++{
    for q:=0;q<5;q++{
      if mat[p][q]==curr{
        ax=p
        ay=q
      }
      if mat[p][q]==bloc {
        bx=p
        by=q
      }
    }
  }

  if(ax==-2){
    return
  }
  if(ax>0 && ay>0 && ax<4 && ay<4){
    mov=8
  }else if((ax+ay)%4==0 && (ax==0 || ay==0 || ax==4 || ay==4)){
    mov=3
  }else{
    mov=5
  }
  if(((ax-bx)==1||(bx-ax)==1)&&((ay-by)==0||(ay-by)==1||(by-ay)==1)){
    mov=mov-1
  }else if(((ay-by)==1||(by-ay)==1)&&((ax-bx)==0||(ax-bx)==1||(bx-ax)==1)){
    mov=mov-1
  }
  fmt.Println(mov)
}