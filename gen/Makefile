all:
	@echo "Determine your CPU architecture and rerun the Makefile with an approriate argument"
	@exit 1

types_386.go: types.go
	go tool cgo -godefs $< > $@  || (rm -f $@ && exit 1)
	gofmt -w $@

types_amd64.go: types.go
	go tool cgo -godefs $< > $@  || (rm -f $@ && exit 1)
	gofmt -w $@

clean:
	rm -rf _obj