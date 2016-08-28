$NetBSD$

 * Add EABI flag for NetBSD/earm in ELF Header
 * Add ELF Section 'note.netbsd.march' which includes MARCH for NetBSD/earm
   (earmv6hf or earmv7hf)

--- src/cmd/link/internal/ld/elf.go.orig	2016-08-15 22:47:58.000000000 +0000
+++ src/cmd/link/internal/ld/elf.go
@@ -947,7 +947,8 @@ func Elfinit() {
 	// 32-bit architectures
 	case sys.ARM:
 		// we use EABI on both linux/arm and freebsd/arm.
-		if HEADTYPE == obj.Hlinux || HEADTYPE == obj.Hfreebsd {
+		if HEADTYPE == obj.Hlinux || HEADTYPE == obj.Hfreebsd || 
+		   HEADTYPE == obj.Hnetbsd {
 			// We set a value here that makes no indication of which
 			// float ABI the object uses, because this is information
 			// used by the dynamic linker to compare executables and
@@ -1257,6 +1258,8 @@ const (
 	ELF_NOTE_NETBSD_DESCSZ  = 4
 	ELF_NOTE_NETBSD_TAG     = 1
 	ELF_NOTE_NETBSD_VERSION = 599000000 /* NetBSD 5.99 */
+	ELF_NOTE_NETBSD_MARCH_NAMESZ	= 7
+	ELF_NOTE_NETBSD_MARCH_TAG	= 5
 )
 
 var ELF_NOTE_NETBSD_NAME = []byte("NetBSD\x00")
@@ -1265,6 +1268,19 @@ func elfnetbsdsig(sh *ElfShdr, startva u
 	n := int(Rnd(ELF_NOTE_NETBSD_NAMESZ, 4) + Rnd(ELF_NOTE_NETBSD_DESCSZ, 4))
 	return elfnote(sh, startva, resoff, n, true)
 }
+func elfnetbsdarmsig(sh *ElfShdr, startva uint64, resoff uint64) int {
+	mArch := []byte("earm\x00")
+	switch obj.Getgoarm() {
+	case 6:
+		mArch = []byte("earmv6hf\x00")
+	case 7:
+		mArch = []byte("earmv7hf\x00")
+	}
+	descsz := len(mArch)
+	n := int(Rnd(ELF_NOTE_NETBSD_MARCH_NAMESZ, 4) + Rnd(int64(descsz), 4))
+	return elfnote(sh, startva, resoff, n, true)
+}
+
 
 func elfwritenetbsdsig() int {
 	// Write Elf_Note header.
@@ -1274,12 +1290,32 @@ func elfwritenetbsdsig() int {
 		return 0
 	}
 
+
 	// Followed by NetBSD string and version.
 	Cwrite(ELF_NOTE_NETBSD_NAME)
 	Cput(0)
 
 	Thearch.Lput(ELF_NOTE_NETBSD_VERSION)
 
+	if SysArch.Family == sys.ARM {
+		mArch := []byte("earm\x00")
+		switch obj.Getgoarm() {
+		case 6:
+			mArch = []byte("earmv6hf\x00")
+		case 7:
+			mArch = []byte("earmv7hf\x00")
+		}
+		descsz := len(mArch)
+		sh2 := elfwritenotehdr(".note.netbsd.march", ELF_NOTE_NETBSD_MARCH_NAMESZ,
+			uint32(descsz), ELF_NOTE_NETBSD_MARCH_TAG)
+		if sh2 == nil {
+			return 0
+		}
+		Cwrite(ELF_NOTE_NETBSD_NAME)
+		Cput(0)
+		Cwrite(mArch)
+		return int(sh.size) + int(sh2.size)
+	}
 	return int(sh.size)
 }
 
@@ -1843,6 +1879,9 @@ func doelf() {
 	}
 	if HEADTYPE == obj.Hnetbsd {
 		Addstring(shstrtab, ".note.netbsd.ident")
+		if SysArch.Family == sys.ARM {
+			Addstring(shstrtab, ".note.netbsd.march")
+		}
 	}
 	if HEADTYPE == obj.Hopenbsd {
 		Addstring(shstrtab, ".note.openbsd.ident")
@@ -2242,6 +2281,10 @@ func Asmbelf(symo int64) {
 		case obj.Hnetbsd:
 			sh = elfshname(".note.netbsd.ident")
 			resoff -= int64(elfnetbsdsig(sh, uint64(startva), uint64(resoff)))
+			if SysArch.Family == sys.ARM {
+				sh = elfshname(".note.netbsd.march")
+				resoff -= int64(elfnetbsdarmsig(sh, uint64(startva), uint64(resoff)))
+			}
 
 		case obj.Hopenbsd:
 			sh = elfshname(".note.openbsd.ident")
