$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- cobc/pplex.l.m4.orig	2016-05-10 12:16:32.000000000 +0000
+++ cobc/pplex.l.m4
@@ -758,6 +758,7 @@ check_dollar_directive (char *buff, int 
 	char			sbuff[5][256];
 	int			isDEFINED, isNOT;
 	int			i;
+	COB_UNUSED(n);
 
 	if (cb_source_format == CB_FORMAT_FIXED) {
 		if (*line_size < 8) {
