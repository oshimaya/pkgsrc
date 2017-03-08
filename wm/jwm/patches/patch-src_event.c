$NetBSD$

--- src/event.c.orig	2017-03-03 16:40:52.000000000 +0000
+++ src/event.c
@@ -389,7 +389,8 @@ void HandleButtonEvent(const XButtonEven
       np = FindClientByWindow(event->window);
       if(np) {
          const char move_resize = (np->state.status & STAT_DRAG)
-            || ((mask == Mod1Mask) && !(np->state.status & STAT_NODRAG));
+            || ((mask == settings.moveMask)
+		&& !(np->state.status & STAT_NODRAG));
          switch(event->button) {
          case Button1:
          case Button2:
