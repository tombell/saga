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

	fmt.Println("Session")
	fmt.Printf("VRSN: %s\n", session.Vrsn.Version())
	fmt.Printf("OENT chunks: %d\n", len(session.Oent))
	fmt.Printf("OREN chunks: %d\n", len(session.Oren))
	fmt.Println("--------")

	for _, oent := range session.Oent {
		fmt.Printf("%s\n", oent.Adat.Fields)
	}

	return nil
}
