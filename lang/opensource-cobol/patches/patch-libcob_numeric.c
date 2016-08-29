$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/31

--- libcob/numeric.c.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob/numeric.c
@@ -30,12 +30,10 @@
 /* Force symbol exports */
 #define	COB_LIB_EXPIMP
 
+#define	COB_LIB_INCLUDE
 #include "libcob.h"
 #include "coblocal.h"
 
-#define	COB_LIB_INCLUDE
-#include "codegen.h"
-
 #define DECIMAL_NAN	-128
 #define DECIMAL_CHECK(d1,d2) \
   if (unlikely(d1->scale == DECIMAL_NAN || d2->scale == DECIMAL_NAN)) { \
