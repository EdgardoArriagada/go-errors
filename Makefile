define build_dev
	go build -o ./bin/go-errors ./cmd/main.go
endef

define build
	go build -ldflags "-s -w" -o ./bin/go-errors ./cmd/main.go
endef

define dev
	ls ./{lib,cmd}/*.go | entr -c $(call build_dev)
endef

.run:
	$(call $(ACTION));

build:
	$(MAKE) ACTION='build' .run

dev:
	$(MAKE) ACTION='dev' .run
