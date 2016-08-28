$NetBSD$

 * Change LOAD Alignment in ELF Program Headers to the same as system standard.

--- src/cmd/link/internal/arm/obj.go.orig	2016-08-15 22:47:58.000000000 +0000
+++ src/cmd/link/internal/arm/obj.go
@@ -100,6 +100,7 @@ func archinit() {
 
 	case obj.Hlinux,
 		obj.Hfreebsd,
+		obj.Hnetbsd,
 		obj.Hnacl,
 		obj.Hdarwin:
 		break
@@ -124,7 +125,6 @@ func archinit() {
 
 	case obj.Hlinux, /* arm elf */
 		obj.Hfreebsd,
-		obj.Hnetbsd,
 		obj.Hopenbsd:
 		ld.Debug['d'] = 0
 		// with dynamic linking
@@ -140,6 +140,21 @@ func archinit() {
 			ld.INITRND = 4096
 		}
 
+	case obj.Hnetbsd:
+		ld.Debug['d'] = 0
+		// with dynamic linking
+		ld.Elfinit()
+		ld.HEADR = ld.ELFRESERVE
+		if ld.INITTEXT == -1 {
+			ld.INITTEXT = 0x10000 + int64(ld.HEADR)
+		}
+		if ld.INITDAT == -1 {
+			ld.INITDAT = 0
+		}
+		if ld.INITRND == -1 {
+			ld.INITRND = 0x10000
+		}
+
 	case obj.Hnacl:
 		ld.Elfinit()
 		ld.HEADR = 0x10000
