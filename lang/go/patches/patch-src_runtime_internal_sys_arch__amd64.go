$NetBSD$

  * Split NetBSD from common cpu arch source.

--- src/runtime/internal/sys/arch_amd64.go.orig
+++ src/runtime/internal/sys/arch_amd64.go
@@ -1,6 +1,7 @@
 // Copyright 2014 The Go Authors. All rights reserved.
 // Use of this source code is governed by a BSD-style
 // license that can be found in the LICENSE file.
+// +build !netbsd
 
 package sys
 
