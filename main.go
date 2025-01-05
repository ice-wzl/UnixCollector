package main
import (
	"fmt"

	"github.com/ice-wzl/UnixCollector"
)


func main() {
	fmt.Println("[+] UnixCollector Started")
	
	internals.getUsersHomedir()
}
