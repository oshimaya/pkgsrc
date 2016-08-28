$NetBSD$

 * add physpagesz variable when netbsd

--- src/runtime/proc.c.orig	2015-09-23 04:37:37.000000000 +0000
+++ src/runtime/proc.c
@@ -44,6 +44,9 @@ P*	runtime·allp[MaxGomaxprocs+1];
 int8*	runtime·goos;
 int32	runtime·ncpu;
 int32	runtime·newprocs;
+#ifdef GOOS_netbsd
+uintptr	runtime·physpagesz;
+#endif
 
 Mutex runtime·allglock;	// the following vars are protected by this lock or by stoptheworld
 G**	runtime·allg;
