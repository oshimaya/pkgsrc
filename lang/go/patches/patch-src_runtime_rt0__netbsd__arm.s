$NetBSD$

 * Support external linker for netbsd/arm

--- src/runtime/rt0_netbsd_arm.s.orig	2016-08-15 22:48:01.000000000 +0000
+++ src/runtime/rt0_netbsd_arm.s
@@ -9,3 +9,8 @@ TEXT _rt0_arm_netbsd(SB),NOSPLIT,$-4
 	MOVW	$4(R13), R1		// argv
 	MOVM.DB.W [R0-R1], (R13)
 	B runtime·rt0_go(SB)
+
+TEXT main(SB),NOSPLIT,$-4
+	MOVM.DB.W [R0-R1], (R13)
+	MOVW	$runtime·rt0_go(SB), R4
+	B	(R4)
