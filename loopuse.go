package main
import "fmt"

func main(){
	for i := 0; i <=100; i++{
		if(i%2 ==0){ 
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}
}