package saga

import (
	"fmt"

	"github.com/tombell/saga/serato"
)

// Run ...
func Run(filepath string) error {
	fmt.Printf("Reading %s...\n", filepath)

	session, err := serato.ReadSession(filepath)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", session)

	return nil
}
