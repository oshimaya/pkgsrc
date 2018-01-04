$NetBSD$

--- src/cmd/link/internal/ld/elf.go.orig	2017-02-16 19:12:22.000000000 +0000
+++ src/cmd/link/internal/ld/elf.go
@@ -1263,6 +1263,8 @@ const (
 	ELF_NOTE_NETBSD_DESCSZ  = 4
 	ELF_NOTE_NETBSD_TAG     = 1
 	ELF_NOTE_NETBSD_VERSION = 599000000 /* NetBSD 5.99 */
+	ELF_NOTE_NETBSD_MARCH_NAMESZ	= 7
+	ELF_NOTE_NETBSD_MARCH_TAG	= 5
 )
 
 var ELF_NOTE_NETBSD_NAME = []byte("NetBSD\x00")
@@ -1271,6 +1273,19 @@ func elfnetbsdsig(sh *ElfShdr, startva u
 	n := int(Rnd(ELF_NOTE_NETBSD_NAMESZ, 4) + Rnd(ELF_NOTE_NETBSD_DESCSZ, 4))
 	return elfnote(sh, startva, resoff, n, true)
 }
+func elfnetbsdarmsig(sh *ElfShdr, startva uint64, resoff uint64) int {
+	mArch := []byte("earm\x00")
+	switch objabi.GOARM {
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
@@ -1280,12 +1295,32 @@ func elfwritenetbsdsig() int {
 		return 0
 	}
 
+
 	// Followed by NetBSD string and version.
 	Cwrite(ELF_NOTE_NETBSD_NAME)
 	Cput(0)
 
 	Thearch.Lput(ELF_NOTE_NETBSD_VERSION)
 
+	if SysArch.Family == sys.ARM {
+		mArch := []byte("earm\x00")
+		switch objabi.GOARM {
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
 
@@ -1905,6 +1940,9 @@ func (ctxt *Link) doelf() {
 	}
 	if Headtype == objabi.Hnetbsd {
 		Addstring(shstrtab, ".note.netbsd.ident")
+		if SysArch.Family == sys.ARM {
+			Addstring(shstrtab, ".note.netbsd.march")
+		}
 	}
 	if Headtype == objabi.Hopenbsd {
 		Addstring(shstrtab, ".note.openbsd.ident")
@@ -2332,6 +2370,10 @@ func Asmbelf(ctxt *Link, symo int64) {
 		case objabi.Hnetbsd:
 			sh = elfshname(".note.netbsd.ident")
 			resoff -= int64(elfnetbsdsig(sh, uint64(startva), uint64(resoff)))
+			if SysArch.Family == sys.ARM {
+				sh = elfshname(".note.netbsd.march")
+				resoff -= int64(elfnetbsdarmsig(sh, uint64(startva), uint64(resoff)))
+			}
 
 		case objabi.Hopenbsd:
 			sh = elfshname(".note.openbsd.ident")
