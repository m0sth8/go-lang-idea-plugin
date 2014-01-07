// Taken from: src/pkg/strconv/atob.go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main
// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ParseBool(str string)(value bool,err error) {
switch str {
case "1","t","T","true","TRUE","True":
return true, nil
case "0","f","F","false","FALSE","False":
return false, nil
}
return false, nil
}
