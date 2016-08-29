$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/call.h.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/call.h
@@ -36,7 +36,7 @@ struct call_stack_list {
 	struct call_stack_list	*parent;
 	struct call_stack_list	*children;
 	struct call_stack_list	*sister;
-	const char		*name;
+	char			*name;
 };
 
 DECLNORET COB_EXPIMP void	cob_call_error		(void) COB_A_NORETURN;
