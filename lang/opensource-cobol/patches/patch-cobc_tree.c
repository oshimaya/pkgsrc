$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- cobc/tree.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ cobc/tree.c
@@ -1275,6 +1275,7 @@ cb_build_picture (const char *str)
 	unsigned char		lastonechar = 0;
 	unsigned char		lasttwochar = 0;
 	unsigned char		buff[COB_SMALL_BUFF];
+	COB_UNUSED(flg);
 
 	pic = make_tree (CB_TAG_PICTURE, CB_CATEGORY_UNKNOWN, sizeof (struct cb_picture));
 	if (strlen (str) > 50) {
