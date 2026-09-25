# $NetBSD$

DISTNAME=	krita-6.0.4
PKGNAME=        ${DISTNAME:S/krita-6.0/krita5-5.3/}
PKGREVISION=	
CATEGORIES=	graphics
MASTER_SITES=	${MASTER_SITE_KDE:=krita/${DISTNAME:S/krita-//}/}

MAINTAINER=	oshima@localhost
HOMEPAGE=	https://www.krita.org/
COMMENT=	Raster graphics editor for digital painting
LICENSE=	gnu-gpl-v3

STILL_QT5=	yes
.include "../../meta-pkgs/kde/applications.mk"

USE_TOOLS+=		msgfmt pkg-config perl
# Uses targets with spaces in the names
USE_TOOLS+=		gmake

CMAKE_CONFIGURE_ARGS+=	-DCMAKE_DISABLE_FIND_PACKAGE_PyQt5:BOOL=true
CMAKE_CONFIGURE_ARGS+=	-DCMAKE_DISABLE_FIND_PACKAGE_KF5KDcraw=ON
CMAKE_CONFIGURE_ARGS+=	-DKRITA_ENABLE_QT6=OFF
CMAKE_CONFIGURE_ARGS+=	-DCMAKE_INSTALL_PREFIX=${LOCALBASE}/krita5
CMAKE_CONFIGURE_ARGS+=	-DBUILD_TESTING=OFF
CMAKE_CONFIGURE_ARGS+=	-DKRITA_BUILD_TESTING=OFF

DEPENDS+=	kitemmodels-[0-9]*:../../devel/kitemmodels
DEPENDS+=	mlt-[0-9]*:../../multimedia/mlt

REPLACE_PYTHON+=	libs/ui/tests/data/AppImageUpdateDummy

SUBST_CLASSES+=		fix-moc
SUBST_STAGE.fix-moc=	post-configure
SUBST_MESSAGE.fix-moc=	Cleaning up MOC includes
SUBST_FILES.fix-moc=	_KDE_build/plugins/*/*/CMakeFiles/*_autogen.dir/AutogenInfo.json
SUBST_NOOP_OK.fix-moc=	yes
SUBST_SED.fix-moc=	-e 's;"${PREFIX}/include",;;g'
SUBST_CLASSES+=		prefix
SUBST_STAGE.prefix=	pre-configure
SUBST_MESSAGE.prefix=	Forcing custom install prefix to krita5 subdirectory.
SUBST_FILES.prefix=	CMakeLists.txt
SUBST_SED.prefix=	-e 's|^project(|set(CMAKE_INSTALL_PREFIX "${LOCALBASE}/krita5" CACHE PATH "Forced" FORCE)\nproject(|'

post-install:
	${RUN} ${MKDIR} ${DESTDIR}${LOCALBASE}/bin
	${INSTALL_SCRIPT} ${FILESDIR}/krita5.sh ${DESTDIR}${LOCALBASE}/bin/krita5
	${RUN} ${MKDIR} ${DESTDIR}${LOCALBASE}/share
	${LN} -sf ${LOCALBASE}/share/mime ${DESTDIR}${LOCALBASE}/krita5/share/mime

.include "../../lang/python/application.mk"
.include "../../archivers/karchive/buildlink3.mk"
.include "../../archivers/quazip/buildlink3.mk"
.include "../../converters/fribidi/buildlink3.mk"
.include "../../devel/libexecinfo/buildlink3.mk"
BUILDLINK_API_DEPENDS.boost-header+=	boost-headers>=1.73.0nb1
.include "../../devel/boost-headers/buildlink3.mk"
.include "../../devel/boost-libs/buildlink3.mk"
.include "../../devel/immer/buildlink3.mk"
.include "../../devel/lager/buildlink3.mk"
.include "../../devel/zug/buildlink3.mk"
.include "../../devel/ki18n/buildlink3.mk"
.include "../../devel/kio/buildlink3.mk"
.include "../../devel/kconfig/buildlink3.mk"
.include "../../devel/kcoreaddons/buildlink3.mk"
.include "../../devel/kcrash/buildlink3.mk"
.include "../../devel/libatomic/buildlink3.mk"
.include "../../devel/SDL2/buildlink3.mk"
.include "../../devel/xsimd/buildlink3.mk"
.include "../../graphics/exiv2/buildlink3.mk"
.include "../../graphics/hicolor-icon-theme/buildlink3.mk"
.include "../../graphics/giflib/buildlink3.mk"
.include "../../graphics/lcms2/buildlink3.mk"
.include "../../graphics/libheif/buildlink3.mk"
.include "../../graphics/libjxl/buildlink3.mk"
.include "../../graphics/libmypaint/buildlink3.mk"
.include "../../graphics/libraw/buildlink3.mk"
.include "../../graphics/libwebp/buildlink3.mk"
.include "../../graphics/opencolorio/buildlink3.mk"
.include "../../graphics/openexr/buildlink3.mk"
.include "../../graphics/openjpeg/buildlink3.mk"
.include "../../graphics/tiff/buildlink3.mk"
.include "../../math/eigen3/buildlink3.mk"
.include "../../math/fftw/buildlink3.mk"
.include "../../math/gsl/buildlink3.mk"
.include "../../multimedia/mlt/buildlink3.mk"
.include "../../print/poppler-qt5/buildlink3.mk"
.include "../../textproc/kcompletion/buildlink3.mk"
.include "../../textproc/libunibreak/buildlink3.mk"
.include "../../x11/kitemviews/buildlink3.mk"
.include "../../x11/kguiaddons/buildlink3.mk"
.include "../../x11/kwidgetsaddons/buildlink3.mk"
.include "../../x11/kwindowsystem/buildlink3.mk"
.include "../../x11/qt5-qtbase/buildlink3.mk"
.include "../../x11/qt5-qtdeclarative/buildlink3.mk"
.include "../../x11/qt5-qtmultimedia/buildlink3.mk"
.include "../../x11/qt5-qtquickcontrols2/buildlink3.mk"
.include "../../x11/qt5-qtsvg/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"

