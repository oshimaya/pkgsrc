$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/common.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/common.c
@@ -2095,6 +2095,9 @@ cobcommandline (int flags, int *pargc, c
 	char		**spenvp;
 	char		*spname;
 	int		sflags;
+	COB_UNUSED(spenvp);
+	COB_UNUSED(spname);
+	COB_UNUSED(sflags);
 
 	if (!cob_initialized) {
 		cob_runtime_error ("'cobcommandline' - Runtime has not been initialized");
@@ -2685,7 +2688,7 @@ cob_acuw_calledby (unsigned char *data)
 			memset (f1->data, ' ', (int)f1->size);
 			return 0;
 		} else {
-			called_program_name = (const char *)cob_current_module->next->program_id;
+			called_program_name = (char *)cob_current_module->next->program_id;
 			if (called_program_name == NULL) {
 				return -1;
 			}
