$NetBSD$

 * Change LOAD Alignment in ELF Program Headers to the same as system standard.

--- src/cmd/link/internal/amd64/obj.go.orig	2016-08-15 22:47:58.000000000 +0000
+++ src/cmd/link/internal/amd64/obj.go
@@ -147,7 +147,6 @@ func archinit() {
 
 	case obj.Hlinux, /* elf64 executable */
 		obj.Hfreebsd,   /* freebsd */
-		obj.Hnetbsd,    /* netbsd */
 		obj.Hopenbsd,   /* openbsd */
 		obj.Hdragonfly, /* dragonfly */
 		obj.Hsolaris:   /* solaris */
@@ -164,6 +163,20 @@ func archinit() {
 			ld.INITRND = 4096
 		}
 
+	case obj.Hnetbsd: /* netbsd */
+		ld.Elfinit()
+
+		ld.HEADR = ld.ELFRESERVE
+		if ld.INITTEXT == -1 {
+			ld.INITTEXT = (1 << 22) + int64(ld.HEADR)
+		}
+		if ld.INITDAT == -1 {
+			ld.INITDAT = 0
+		}
+		if ld.INITRND == -1 {
+			ld.INITRND = 0x200000
+		}
+
 	case obj.Hnacl:
 		ld.Elfinit()
 		ld.Debug['w']++ // disable dwarf, which gets confused and is useless anyway
