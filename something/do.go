package something

import (
	"fmt"
	"log"
)

func DoSomething(input string) string {
	log.Printf("From DoSomething")
	return fmt.Sprintf("DoSomething %s", input)
}
