$NetBSD$

--- src/key.h.orig	2017-03-03 16:42:59.000000000 +0000
+++ src/key.h
@@ -71,6 +71,12 @@ extern unsigned int lockMask;
  */
 KeyType GetKey(const XKeyEvent *event);
 
+/** Parse a modifier string.
+ * @param str The modifier string.
+ * @return The modifier mask.
+ */
+unsigned int ParseModifierString(const char *str);
+
 /** Insert a key binding.
  * @param key The key binding type.
  * @param modifiers The modifier mask.
