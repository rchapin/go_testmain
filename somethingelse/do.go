package somethingelse

import (
	"fmt"
	"log"
)

func DoSomethingElse(input string) string {
	log.Printf("From DoSomethingElse")
	return fmt.Sprintf("DoSomethingElse %s", input)
}
