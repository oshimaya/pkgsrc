$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/move.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/move.c
@@ -554,6 +554,7 @@ cob_move_fp_to_display (cob_field *f1, c
 	char		*x, *y;
 	char		buff[64];
 	char		buff2[64];
+	COB_UNUSED(frac);
 
 	memset ((ucharptr)buff, 0, sizeof (buff));
 	memset ((ucharptr)buff2, 0, sizeof (buff2));
