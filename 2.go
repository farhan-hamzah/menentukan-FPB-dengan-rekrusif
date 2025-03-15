package main
import "fmt"

func pers(a, b int)int{
	if b == 0{
		return a
	}else{
		return pers(b, a%b)
	}
}

func main(){
	var d1, d2 int
	fmt.Scan(&d1, &d2)
	fmt.Print(pers(d1, d2))
}