package main
import "fmt"

func main(){
	a, b, c := 10, 33, 499
	fmt.Println("A = ",a,"\nB = ",b,"\nC = ",c)
	if (a> b && a >c){
		fmt.Println("\nA is Large Number")
	} else if(b>a && b>c){
		fmt.Println("\nB is Large Number")
	} else {
		fmt.Println("\nC is Large Number")
	}
}