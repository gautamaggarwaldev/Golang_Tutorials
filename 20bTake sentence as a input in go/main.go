package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence : ")
	sentence,_ := reader.ReadString('\n')
	fmt.Println(sentence)
}