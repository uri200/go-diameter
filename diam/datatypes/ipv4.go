// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatypes

import (
	"fmt"
	"net"
)

// IPv4 Diameter Type for Framed-IP-Address AVP.
type IPv4 net.IP

func DecodeIPv4(b []byte) (IPv4, error) {
	return IPv4(b), nil
}

func (ip IPv4) Serialize() []byte {
	if ip4 := net.IP(ip).To4(); ip4 != nil {
		return ip4
	}
	return ip
}

func (ip IPv4) String() string {
	return fmt.Sprintf("IPv4{%s}", net.IP(ip))
}
