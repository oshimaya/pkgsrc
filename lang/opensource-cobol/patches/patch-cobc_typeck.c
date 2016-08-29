$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- cobc/typeck.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ cobc/typeck.c
@@ -832,6 +832,7 @@ cb_reference_type_check (cb_tree ref, cb
 	char			strbuf[256];
 	int			offset = 0 ;
 	int			ret = 0;
+	COB_UNUSED(r);
 
 	r = CB_REFERENCE (ref);
 	switch (CB_TREE_TAG (x)) {
