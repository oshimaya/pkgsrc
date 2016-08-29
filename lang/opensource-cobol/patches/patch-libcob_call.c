$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/30

--- libcob/call.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/call.c
@@ -56,10 +56,6 @@
 #include <windows.h>
 /* Prototype */
 static char *	lt_dlerror (void);
-void			init_call_stack_list (void);
-struct call_stack_list
-			cob_create_call_stack_list (const char *);
-void			cob_cancel_call_stack_list (struct call_stack_list *);
 
 static HMODULE
 lt_dlopen (const char *x)
@@ -285,6 +281,48 @@ lookup (const char *name)
 	return NULL;
 }
 
+static void
+init_call_stack_list (void)
+{
+	if (!call_stack_list_head) {
+		call_stack_list_head = cob_malloc (sizeof (struct call_stack_list));
+		memset (call_stack_list_head, 0, sizeof (struct call_stack_list));
+	}
+	current_call_stack_list = call_stack_list_head;
+}
+
+static struct call_stack_list *
+cob_create_call_stack_list (char *name)
+{
+	struct call_stack_list *new_list = cob_malloc (sizeof (struct call_stack_list));
+	memset (new_list, 0, sizeof (struct call_stack_list));
+	new_list->parent = current_call_stack_list;
+	new_list->name = cob_malloc (strlen (name) + 1);
+	strcpy (new_list->name, name);
+	current_call_stack_list = new_list;
+	return new_list;
+}
+
+static void
+cob_cancel_call_stack_list (struct call_stack_list *p)
+{
+	if (!p) {
+		/*No program*/
+		return;
+	}
+	static cob_field_attr a_2 = {33, 0, 0, 0, NULL};
+	cob_field f = {strlen (p->name), (unsigned char *) p->name, &a_2};
+	cob_field_cancel (&f);
+	if (p->children) {
+		cob_cancel_call_stack_list (p->children);
+	}
+	struct call_stack_list *s = p->sister;
+	while (s != NULL) {
+		cob_cancel_call_stack_list (s);
+		s = s->sister;
+	}
+}
+
 const char *
 cob_resolve_error (void)
 {
@@ -676,6 +714,7 @@ void *
 cobsavenv2 (struct cobjmp_buf *jbuf, const int jsize)
 {
 	int	jtemp;
+	COB_UNUSED(jtemp);
 
 	/* Shut up compiler */
 	jtemp = jsize;
@@ -698,28 +737,6 @@ coblongjmp (struct cobjmp_buf *jbuf)
 }
 
 void
-init_call_stack_list ()
-{
-	if (!call_stack_list_head) {
-		call_stack_list_head = cob_malloc (sizeof (struct call_stack_list));
-		memset (call_stack_list_head, 0, sizeof (struct call_stack_list));
-	}
-	current_call_stack_list = call_stack_list_head;
-}
-
-struct call_stack_list *
-cob_create_call_stack_list (char *name)
-{
-	struct call_stack_list *new_list = cob_malloc (sizeof (struct call_stack_list));
-	memset (new_list, 0, sizeof (struct call_stack_list));
-	new_list->parent = current_call_stack_list;
-	new_list->name = cob_malloc (strlen (name) + 1);
-	strcpy (new_list->name, name);
-	current_call_stack_list = new_list;
-	return new_list;
-}
-
-void
 cob_push_call_stack_list (char *name)
 {
 	if (!current_call_stack_list) {
@@ -762,26 +779,6 @@ cob_pop_call_stack_list ()
 }
 
 void
-cob_cancel_call_stack_list (struct call_stack_list *p)
-{
-	if (!p) {
-		/*No program*/
-		return;
-	}
-	static cob_field_attr a_2 = {33, 0, 0, 0, NULL};
-	cob_field f = {strlen (p->name), (unsigned char *) p->name, &a_2};
-	cob_field_cancel (&f);
-	if (p->children) {
-		cob_cancel_call_stack_list (p->children);
-	}
-	struct call_stack_list *s = p->sister;
-	while (s != NULL) {
-		cob_cancel_call_stack_list (s);
-		s = s->sister;
-	}
-}
-
-void
 cob_cancel_all ()
 {
 	if (!current_call_stack_list) {
