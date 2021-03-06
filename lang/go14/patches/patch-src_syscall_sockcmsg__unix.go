$NetBSD$

--- src/syscall/sockcmsg_unix.go.orig	2017-07-23 05:37:09.000000000 +0000
+++ src/syscall/sockcmsg_unix.go
@@ -18,6 +18,9 @@ func cmsgAlignOf(salen int) int {
 	if darwin64Bit || dragonfly64Bit {
 		salign = 4
 	}
+	if netbsd32BitArm {
+		salign = 8
+	}
 	return (salen + salign - 1) & ^(salign - 1)
 }
 
