package main
import 
(
	"fmt"
	"crypto/rand"
	"math/big"
)
func main() {
	fmt.Println("generate random");

	randInt, err := rand.Int(rand.Reader, big.NewInt(10));

	if err != nil {
		fmt.Println(err);
		return;
	}
	fmt.Println(randInt);
	


}