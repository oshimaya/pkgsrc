$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/strings.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/strings.c
@@ -268,6 +268,7 @@ cob_inspect_before (const cob_field *str
 	unsigned char	*p2;
 	unsigned int	n;
 	int		fig;
+	COB_UNUSED(sign);
 
 	switch (COB_FIELD_TYPE (str)) {
 	case COB_TYPE_NUMERIC_DISPLAY:
