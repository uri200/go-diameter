// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatypes

import "fmt"

// IPFilterRule Diameter Type.
type IPFilterRule OctetString

func DecodeIPFilterRule(b []byte) (IPFilterRule, error) {
	return IPFilterRule(OctetString(b)), nil
}

func (s IPFilterRule) Serialize() []byte {
	return OctetString(s).Serialize()
}

func (s IPFilterRule) String() string {
	return fmt.Sprintf("IPFilterRule{%s}", string(s))
}
