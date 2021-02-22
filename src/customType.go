package main
import ("fmt")
//var x int = 12
var y string = "James Bond"
var z bool = true
type soumendra int
var x soumendra

func main() {
	//s := fmt.Sprintf("%v %v %v", x, y, z)
	//fmt.Print(s)
	fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Println("value of x:", x)
	fmt.Printf("Type of x: %T", x)
	x = 42
	fmt.Println("value of x:", x)
}