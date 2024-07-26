package somethingelseinttest

import (
	"log"
	"os"
	"testing"

	"github.com/rchapin/go_testmain/somethingelse"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	runExitCode := m.Run()
	log.Println("DoSomethingElse Integration tests complete")
	os.Exit(runExitCode)
}

func TestDoSomethingElse(t *testing.T) {
	assert.Equal(t, "DoSomethingElse foo", somethingelse.DoSomethingElse("foo"))
}
