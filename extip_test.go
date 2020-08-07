package extip

import (
	"fmt"
	"testing"
)

func TestGetIP(t *testing.T) {
        address, err := GetIP()
        fmt.Printf("GetIP: %s\n", address)
        if err != nil {
                t.Errorf("Error: %s", err)
        }
}

func TestGetIP4(t *testing.T) {
	address, err := GetIP4()
	fmt.Printf("GetIP4: %s\n", address)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGetIP6(t *testing.T) {
	address, err := GetIP6()
	fmt.Printf("GetIP6: %s\n", address)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
