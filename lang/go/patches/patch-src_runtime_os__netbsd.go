$NetBSD$

 * Change const to variable PhysPageSize, initialize from sysctl 'HW_PAGESIZE'.
 * Change sysctl 'HW_NCPU' to 'HW_NCPUONLINE' as actual number of cpu.

--- src/runtime/os_netbsd.go.orig	2016-08-15 22:48:00.000000000 +0000
+++ src/runtime/os_netbsd.go
@@ -6,6 +6,7 @@ package runtime
 
 import (
 	"runtime/internal/atomic"
+	"runtime/internal/sys"
 	"unsafe"
 )
 
@@ -80,11 +81,12 @@ var sigset_all = sigset{[4]uint32{^uint3
 // From NetBSD's <sys/sysctl.h>
 const (
 	_CTL_HW  = 6
-	_HW_NCPU = 3
+	_HW_NCPUONLINE = 16
+	_HW_PAGESIZE = 7
 )
 
 func getncpu() int32 {
-	mib := [2]uint32{_CTL_HW, _HW_NCPU}
+	mib := [2]uint32{_CTL_HW, _HW_NCPUONLINE}
 	out := uint32(0)
 	nout := unsafe.Sizeof(out)
 	ret := sysctl(&mib[0], 2, (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
@@ -94,6 +96,17 @@ func getncpu() int32 {
 	return 1
 }
 
+func getPhysPageSize()  uintptr {
+	mib := [2]uint32{_CTL_HW, _HW_PAGESIZE}
+	out := uintptr(0)
+	nout := unsafe.Sizeof(out)
+	ret := sysctl(&mib[0], 2, (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
+	if ret >= 0 {
+		return out
+	}
+        return 0
+}
+
 //go:nosplit
 func semacreate(mp *m) {
 }
@@ -186,6 +199,7 @@ func netbsdMstart() {
 
 func osinit() {
 	ncpu = getncpu()
+	sys.PhysPageSize = getPhysPageSize()
 }
 
 var urandom_dev = []byte("/dev/urandom\x00")
