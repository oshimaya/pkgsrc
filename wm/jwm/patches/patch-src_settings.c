$NetBSD$

--- src/settings.c.orig	2017-03-03 16:47:02.000000000 +0000
+++ src/settings.c
@@ -21,6 +21,7 @@ static void FixRange(unsigned int *value
 /** Initialize settings. */
 void InitializeSettings(void)
 {
+   settings.moveMask = (1 << Mod1MapIndex);
    settings.doubleClickSpeed = 400;
    settings.doubleClickDelta = 2;
    settings.snapMode = SNAP_BORDER;
