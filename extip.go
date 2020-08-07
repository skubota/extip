// Package extip
package extip

import (
	"gortc.io/stun"
	"net"
)

// STUN servers
var stunsrv string = "stun.l.google.com:19302"

// GetIP is get external address
func GetIP() (net.IP, error) {
        c, err := stun.Dial("udp", stunsrv)
        var xorAddr stun.XORMappedAddress
        if err != nil {
                // no route to host
                return net.IP(nil), nil
        }
        message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
        if err := c.Do(message, func(res stun.Event) {
                if res.Error != nil {
                        err = res.Error
                }
                if err := xorAddr.GetFrom(res.Message); err != nil {
                        return
                }
        }); err != nil {
                return net.IP(nil), err
        }
        return xorAddr.IP, nil
}

// GetIP4 is get external ipv4 address
func GetIP4() (net.IP, error) {
	c, err := stun.Dial("udp4", stunsrv)
	var xorAddr stun.XORMappedAddress
	if err != nil {
		// no route to host
		return net.IP(nil), nil
	}
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			err = res.Error
		}
		if err := xorAddr.GetFrom(res.Message); err != nil {
			return
		}
	}); err != nil {
		return net.IP(nil), err
	}
	return xorAddr.IP, nil
}

// GetIP6 is get external ipv6 address
func GetIP6() (net.IP, error) {
	c, err := stun.Dial("udp6", stunsrv)
	var xorAddr stun.XORMappedAddress
	if err != nil {
		// no route to host
		return net.IP(nil), nil
	}
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			err = res.Error
		}
		if err := xorAddr.GetFrom(res.Message); err != nil {
			return
		}
	}); err != nil {
		return net.IP(nil), err
	}
	return xorAddr.IP, nil
}
