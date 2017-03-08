$NetBSD$

--- src/parse.c.orig	2017-03-03 16:45:17.000000000 +0000
+++ src/parse.c
@@ -425,6 +425,10 @@ void ParseMoveMode(const TokenNode *tp)
    if(str) {
       settings.desktopDelay = ParseUnsigned(tp, str);
    }
+   str = FindAttribute(tp->attributes, "mask");
+   if(str && *str) {
+      settings.moveMask = ParseModifierString(str);
+   }
 
    settings.moveStatusType = ParseStatusWindowType(tp);
    settings.moveMode = ParseTokenValue(mapping, ARRAY_LENGTH(mapping), tp,
