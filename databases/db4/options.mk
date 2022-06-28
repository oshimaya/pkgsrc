# $NetBSD: options.mk,v 1.6 2013/11/30 16:25:12 bsiegert Exp $

PKG_OPTIONS_VAR=	PKG_OPTIONS.db4

PKG_SUPPORTED_OPTIONS=	doc posixmutexs tcl
PKG_SUGGESTED_OPTIONS=	doc

.include "../../mk/bsd.prefs.mk"

.include "../../mk/bsd.prefs.mk"
.if (${MACHINE_ARCH} == "sh3el" || ${MACHINE_ARCH} == "sh3eb" || \
     ${MACHINE_ARCH} == "vax" ) && (${OPSYS} == "NetBSD")
PKG_SUGGESTED_OPTIONS+=        posixmutexes
.endif

.if ${OPSYS} == "MirBSD"
PLIST_SRC+=		PLIST.${OPSYS}
.else
PLIST_SRC+=		PLIST
.endif

.include "../../mk/bsd.options.mk"

###
### Install documentation files
###
.if empty(PKG_OPTIONS:Mdoc)
SUBST_CLASSES+=		docs
SUBST_STAGE.docs=	pre-configure
SUBST_FILES.docs=	dist/Makefile.in
SUBST_SED.docs=		-e '/^library_install:/s, install_docs,,'
.else
PLIST_SRC+=		PLIST.docs
.endif

.if !empty(PKG_OPTIONS:Mposixmutexes)
CONFIGURE_ARGS+=       --enable-posixmutexes
.endif

.if !empty(PKG_OPTIONS:Mtcl)
.include "../../lang/tcl85/buildlink3.mk"
CONFIGURE_ARGS+=       --with-tcl=${BUILDLINK_PREFIX.tcl}/lib
.endif
