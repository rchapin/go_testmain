package dosomethinginttest

import (
	"log"
	"os"
	"testing"

	"github.com/rchapin/go_testmain/something"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	runExitCode := m.Run()
	log.Println("DoSomething Integration tests complete")
	os.Exit(runExitCode)
}

func TestDoSomething(t *testing.T) {
	assert.Equal(t, "DoSomething moo", something.DoSomething("moo"))
}
