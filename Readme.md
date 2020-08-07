Package extip is external IPv4 or IPv6 address check
===================
[![GoDev][godev-image]][godev-url]
![Go](https://github.com/skubota/extip/workflows/Go/badge.svg)

[godev-image]: https://pkg.go.dev/badge/github.com/skubota/extip
[godev-url]: https://pkg.go.dev/github.com/skubota/extip


DESCRIPTION

extip returns the public facing IPv4 or IPv6 address of the requesting client by querying servers at STUN

usage

	$ go get github.com/skubota/extip

package extip 

<https://pkg.go.dev/github.com/skubota/extip>

```go
package main

import (
	"fmt"
	"github.com/skubota/extip"
)

func main() {
	// Any
	addr,err := extip.GetIP()
	if err != nil {
		fmt.Errorf("Address cannot get : %s\n", err)
	}else{
		fmt.Printf("Adress : %s\n",addr) 
	}


	// IPv4
	ipv4,err := extip.GetIP4()
	if err != nil {
		fmt.Errorf("IPv4 address cannot get : %s\n", err)
	}else{
		fmt.Printf("IPv4 address : %s\n", ipv4)
	}

	// IPv6
	ipv6,err := extip.GetIP6()
	if err != nil {
		fmt.Errorf("IPv6 address cannot get : %s\n", err)
	}else{
		fmt.Printf("IPv6 address : %s\n", ipv6)
	}
}
```



