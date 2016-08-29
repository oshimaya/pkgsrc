$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/31

--- libcob.h.orig	2016-05-10 12:16:32.000000000 +0000
+++ libcob.h
@@ -37,6 +37,7 @@ extern "C" {
 #include <libcob/strings.h>
 #include <libcob/termio.h>
 #include <libcob/intrinsic.h>
+#include <libcob/codegen.h>
 
 #ifdef __cplusplus
 }
