#Dockerfile Vars
go_ver=1.18
ver=v0.0.1

#vars
IMAGENAME=tga
REPO=nexus.oots.dev/homelab
VERSION=${ver}
IMAGEFULLNAME=${REPO}/${IMAGENAME}:${VERSION}


.PHONY: help build build_local push run_test fmt all

help:
	@echo "Makefile arguements:"
	@echo ""
	@echo "go_ver   - Go Version"
	@echo "key_name - Name of the ssh key"
	@echo "ver      - Container Version"
	@echo ""
	@echo "Makefile commands:"
	@echo "build"
	@echo "local_build"
	@echo "push"
	@echo "run_test"
	@echo "all"


.DEFAULT_GOAL := all

build:
	@docker build --pull --build-arg go_ver=${go_ver} -t ${IMAGEFULLNAME} .

build_local:
	@docker build --pull --build-arg go_ver=${go_ver} -t ${IMAGENAME}:${VERSION} .

push: 
	@docker push ${IMAGEFULLNAME}

run_test: 
	@docker run -it ${IMAGEFULLNAME}

fmt:
  @gofmt -w ./cmd

all: build push
