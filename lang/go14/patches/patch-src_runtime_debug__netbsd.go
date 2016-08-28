$NetBSD$

 * add function return 'Real' pagesize from runtime package

--- src/runtime/debug_netbsd.go.orig	2016-08-27 17:50:31.000000000 +0000
+++ src/runtime/debug_netbsd.go
@@ -0,0 +1,10 @@
+// Copyright 2009 The Go Authors. All rights reserved.
+// Use of this source code is governed by a BSD-style
+// license that can be found in the LICENSE file.
+// +build netbsd
+
+package runtime
+
+func Physpagesize() int {
+	return int(physpagesz)
+}
