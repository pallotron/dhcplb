/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package dhcplb

import (
	"fmt"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv6"
	"github.com/insomniacslk/dhcp/iana"
)

// GetDUIDLL generates a DUID-LL based on the MAC address of the first
// available, up, non-loopback network interface. This provides a stable,
// predictable DUID for the server.
func GetDUIDLL() (*dhcpv6.DUIDLL, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}
		// Skip interfaces without a MAC address
		if len(iface.HardwareAddr) == 0 {
			continue
		}

		// Found a suitable interface
		return &dhcpv6.DUIDLL{
			HWType:        iana.HWTypeEthernet,
			LinkLayerAddr: iface.HardwareAddr,
		}, nil
	}

	return nil, fmt.Errorf("no suitable network interface found to generate a DUID")
}
