$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/screenio.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/screenio.c
@@ -1611,6 +1611,8 @@ CBL_OC_KEISEN (unsigned char * cmd, unsi
 	int	k_lng2;
 	int	k_color;
 	int	k_prn;
+	COB_UNUSED(k_color);
+	COB_UNUSED(k_prn);
 
 	COB_CHK_PARMS (CBL_OC_ATTRIBUTE, 5);
 
