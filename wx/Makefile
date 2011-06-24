wx_top_builddir = ${WXHOME}
WXDEFS=-D__WXMSW__ ${WXDBG} -I${wx_top_builddir}/${WXBLD} -I${wx_top_builddir}/include -DWX_PRECOMP

SWIG_CXXFLAGS=${WXDEFS}

WXLIBS=$(WXLIBDIR) $(__WXLIB_CORE_p)  $(__WXLIB_BASE_p)  $(__WXLIB_MONO_p) $(__LIB_TIFF_p) $(__LIB_JPEG_p) $(__LIB_PNG_p) $(EXTRALIBS_FOR_GUI) $(__LIB_ZLIB_p) $(__LIB_ODBC_p) $(__LIB_REGEX_p) $(__LIB_EXPAT_p) $(EXTRALIBS_FOR_BASE)
EXTRALIBS =  -mthreads  -lrpcrt4 -loleaut32 -lole32 -luuid -lwinspool -lwinmm -lshell32 -lcomctl32 -lcomdlg32 -lctl3d32 -ladvapi32 -lwsock32 -lgdi32 
SWIG_LDFLAGS=${WXLIBS} ${EXTRALIBS}

CXXFLAGS=-c -g -O0 ${WXDEFS}

include ${GOROOT}/src/Make.inc

SWIGOPT+=${WXDEFS}

TARG = wx
	
SWIGFILES=\
	wx.swigcxx

CXXFILES=\
	wxwrap.cxx

GOFILES=wxwrap.go
	
CXXOBJS=$(patsubst %.cxx,_obj/%.o,$(CXXFILES))
	
include ${GOROOT}/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) $(SWIG_RPATH) -o $@ $*.$0
	
_obj/%.o: %.cxx $(HFILES)
	@mkdir -p _obj
	$(HOST_CXX) $(CXXFLAGS) -o $@ $*.cxx

	
