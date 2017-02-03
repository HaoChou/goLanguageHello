package main
import(
	"fmt"
	"math/rand"
	"math"
)
func main(){
	fmt.Println("My favorite number is ",rand.Intn(10))
	fmt.Println("My favorite number is ",math.Pi) //Pi is the unexported name the first letter must be a Uppercase letter
}