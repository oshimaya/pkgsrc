$NetBSD$

--- src/settings.h.orig	2017-03-03 16:48:13.000000000 +0000
+++ src/settings.h
@@ -76,6 +76,7 @@ typedef struct {
    unsigned int menuOpacity;
    unsigned int desktopDelay;
    unsigned int cornerRadius;
+   unsigned int moveMask;
    AlignmentType titleTextAlignment;
    SnapModeType snapMode;
    MoveModeType moveMode;
