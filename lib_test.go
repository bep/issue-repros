package issuerepros

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestTxtTarNewlines(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testscripts/newlines",
	})
}

func TestMain(m *testing.M) {
	os.Exit(
		testscript.RunMain(m, map[string]func() int{
			// The main program.
			"sha256": func() int {
				filename := os.Args[1]
				sha, err := Sha256(filename)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("SHA256:", sha)
				return 0
			},
		}),
	)
}
