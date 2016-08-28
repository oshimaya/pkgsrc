$NetBSD$

 * Set default GOARM to 6 for NetBSD/arm bucause some sources of golang
   use armv6 or higher instruction always, so real armv5 cpu cannot work.

--- src/cmd/dist/util.go.orig	2016-08-15 22:47:58.000000000 +0000
+++ src/cmd/dist/util.go
@@ -527,6 +527,13 @@ func xgetgoarm() string {
 		// OpenBSD currently only supports softfloat.
 		return "5"
 	}
+	if goos == "netbsd"{
+		// NetBSD supports many arm machines with armv5 cpu,
+		// But some current sources for go/arm support 
+		// only armv6 or lator
+		// (See runtime/internal/atomic/asm_arm.s)
+		return "6"
+	}
 
 	// Try to exec ourselves in a mode to detect VFP support.
 	// Seeing how far it gets determines which instructions failed.
