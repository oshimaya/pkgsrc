$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- cobc/codegen.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ cobc/codegen.c
@@ -998,6 +998,7 @@ output_param (cb_tree x, int id)
 	int			extrefs;
 	int			sav_stack_id;
 	char			fname[12];
+	COB_UNUSED(extrefs);
 
 	param_id = id;
 
