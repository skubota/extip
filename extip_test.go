package extip

import (
	"fmt"
	"testing"
)

func TestGetIP(t *testing.T) {
        address, err := GetIP()
        if err != nil {
                t.Errorf("Error: %s", err)
        }
        fmt.Printf("GetIP: %s\n", address)
}

func TestGetPort(t *testing.T) {
        port, err := GetPort()
        if err != nil {
                t.Errorf("Error: %s", err)
        }
        fmt.Printf("GetPort: %d\n", port)
}

func TestGetIP4(t *testing.T) {
	address, err := GetIP4()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Printf("GetIP4: %s\n", address)
}

func TestGetPort4(t *testing.T) {
        port4, err := GetPort4()
        if err != nil {
                t.Errorf("Error: %s", err)
        }
        fmt.Printf("GetPort4: %d\n", port4)
}

func TestGetIP6(t *testing.T) {
	address, err := GetIP6()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Printf("GetIP6: %s\n", address)
}

func TestGetPort6(t *testing.T) {
        port6, err := GetPort6()
        if err != nil {
                t.Errorf("Error: %s", err)
        }
        fmt.Printf("GetPort6: %d\n", port6)
}
