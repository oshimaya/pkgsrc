$NetBSD$

 * Change EABI for NetBSD/arm, Add Padding for aliment for 64bit valiable.

--- src/runtime/defs1_netbsd_arm.go.orig	2016-08-15 22:48:00.000000000 +0000
+++ src/runtime/defs1_netbsd_arm.go
@@ -110,6 +110,7 @@ type stackt struct {
 type timespec struct {
 	tv_sec  int64
 	tv_nsec int32
+	Pad_cgo_0 [4]byte
 }
 
 func (ts *timespec) set_sec(x int32) {
@@ -123,6 +124,7 @@ func (ts *timespec) set_nsec(x int32) {
 type timeval struct {
 	tv_sec  int64
 	tv_usec int32
+	Pad_cgo_0 [4]byte
 }
 
 func (tv *timeval) set_usec(x int32) {
@@ -136,9 +138,10 @@ type itimerval struct {
 
 type mcontextt struct {
 	__gregs [17]uint32
-	__fpu   [4 + 8*32 + 4]byte // EABI
-	// __fpu [4+4*33+4]byte // not EABI
+	Pad_cgo_0 [4]byte
+	__fpu	[272]byte
 	_mc_tlsbase uint32
+	Pad_cgo_1 [4]byte
 }
 
 type ucontextt struct {
@@ -146,6 +149,7 @@ type ucontextt struct {
 	uc_link     *ucontextt
 	uc_sigmask  sigset
 	uc_stack    stackt
+	Pad_cgo_0 [4]byte
 	uc_mcontext mcontextt
 	__uc_pad    [2]int32
 }
@@ -157,6 +161,7 @@ type keventt struct {
 	fflags uint32
 	data   int64
 	udata  *byte
+	Pad_cgo_0 [4]byte
 }
 
 // created by cgo -cdefs and then converted to Go
