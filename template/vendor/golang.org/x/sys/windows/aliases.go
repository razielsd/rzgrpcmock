// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows && go1.9
// +build windows,go1.9

package windows

import "syscall"

type (
	Errno       = syscall.Errno
	SysProcAttr = syscall.SysProcAttr
)
