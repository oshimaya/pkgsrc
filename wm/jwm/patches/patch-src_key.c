$NetBSD$

--- src/key.c.orig	2017-03-03 16:42:07.000000000 +0000
+++ src/key.c
@@ -75,7 +75,6 @@ static KeyNode *bindings;
 unsigned int lockMask;
 
 static unsigned int GetModifierMask(XModifierKeymap *modmap, KeySym key);
-static unsigned int ParseModifierString(const char *str);
 static KeySym ParseKeyString(const char *str);
 static char ShouldGrab(KeyType key);
 static void GrabKey(KeyNode *np, Window win);
