include $(GOROOT)/src/Make.inc

TARG=⚛perf

GOFILES=\
	src/perf.go\
	src/perf_$(GOARCH).go\
	src/types.$(O).go\

include $(GOROOT)/src/Make.pkg

src/types.${O}.go: src/types.c
	godefs -g perf src/types.c > src/types.$(O).go
	#gofmt -w src/types.$(O).go
