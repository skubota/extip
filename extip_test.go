package extip

import (
	"fmt"
	"testing"
)

func TestGet4(t *testing.T) {
	address, err := Get4()
	fmt.Printf("Get4: %s\n", address)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGet6(t *testing.T) {
	address, err := Get6()
	fmt.Printf("Get6: %s\n", address)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
