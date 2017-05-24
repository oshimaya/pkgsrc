$NetBSD$

--- src/runtime/os_netbsd.go.orig	2017-02-16 19:12:24.000000000 +0000
+++ src/runtime/os_netbsd.go
@@ -79,13 +79,13 @@ var sigset_all = sigset{[4]uint32{^uint3
 
 // From NetBSD's <sys/sysctl.h>
 const (
-	_CTL_HW      = 6
-	_HW_NCPU     = 3
+	_CTL_HW  = 6
+	_HW_NCPUONLINE = 16
 	_HW_PAGESIZE = 7
 )
 
 func getncpu() int32 {
-	mib := [2]uint32{_CTL_HW, _HW_NCPU}
+	mib := [2]uint32{_CTL_HW, _HW_NCPUONLINE}
 	out := uint32(0)
 	nout := unsafe.Sizeof(out)
 	ret := sysctl(&mib[0], 2, (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
@@ -167,13 +167,23 @@
 	var uc ucontextt
 	getcontext(unsafe.Pointer(&uc))
 
+	// XXX: _UC_SIGMASK does not seem to work here.
+	// It would be nice if _UC_SIGMASK and _UC_STACK
+	// worked so that we could do all the work setting
+	// the sigmask and the stack here, instead of setting
+	// the mask here and the stack in netbsdMstart
+	// For now do the blocking manually.
 	uc.uc_flags = _UC_SIGMASK | _UC_CPU
 	uc.uc_link = nil
 	uc.uc_sigmask = sigset_all
 
+	var oset sigset
+	sigprocmask(_SIG_SETMASK, &sigset_all, &oset)
+
 	lwp_mcontext_init(&uc.uc_mcontext, stk, mp, mp.g0, funcPC(netbsdMstart))
 
 	ret := lwp_create(unsafe.Pointer(&uc), 0, unsafe.Pointer(&mp.procid))
+	sigprocmask(_SIG_SETMASK, &oset, nil)
 	if ret < 0 {
 		print("runtime: failed to create new OS thread (have ", mcount()-1, " already; errno=", -ret, ")\n")
 		if ret == -_EAGAIN {
