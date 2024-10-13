// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package ldap

// SecurityProtocol protocol type
type SecurityProtocol int

// Note: new type must be added at the end of list to maintain compatibility.
const (
	SecurityProtocolUnencrypted SecurityProtocol = iota
	SecurityProtocolLDAPS
	SecurityProtocolStartTLS
)

// String returns the name of the SecurityProtocol
func (s SecurityProtocol) String() string {
	return SecurityProtocolNames[s]
}

// Int returns the int value of the SecurityProtocol
func (s SecurityProtocol) Int() int {
	return int(s)
}

// SecurityProtocolNames contains the name of SecurityProtocol values.
var SecurityProtocolNames = map[SecurityProtocol]string{
	SecurityProtocolUnencrypted: "Unencrypted",
	SecurityProtocolLDAPS:       "LDAPS",
	SecurityProtocolStartTLS:    "StartTLS",
}
