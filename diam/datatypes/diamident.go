// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatypes

import "fmt"

// DiameterIdentity Diameter Type.
type DiameterIdentity OctetString

func DecodeDiameterIdentity(b []byte) (DiameterIdentity, error) {
	return DiameterIdentity(OctetString(b)), nil
}

func (s DiameterIdentity) Serialize() []byte {
	return OctetString(s).Serialize()
}

func (s DiameterIdentity) String() string {
	return fmt.Sprintf("DiameterIdentity{%s}", string(s))
}
