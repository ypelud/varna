
VARNA_ENVS := \
	-e OS_ARCH_ARG \
	-e OS_PLATFORM_ARG \
	-e TESTFLAGS \
	-e VERBOSE \
	-e VERSION \
	-e CODENAME \
	-e TESTDIRS \
	-e CI

BIND_DIR := "dist"
VARNA_MOUNT := -v "$(CURDIR)/$(BIND_DIR):/go/src/github.com/ypelud/varna/$(BIND_DIR)"

REPONAME := $(shell echo $(REPO) | tr '[:upper:]' '[:lower:]')
VARNA_IMAGE := $(if $(REPONAME),$(REPONAME),"ypelud/varna")

GIT_BRANCH := $(subst heads/,,$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))
VARNA_DEV_IMAGE := varna-dev$(if $(GIT_BRANCH),:$(subst /,-,$(GIT_BRANCH)))

DOCKER_RUN_OPTS := $(VARNA_ENVS) $(VARNA_MOUNT) "$(VARNA_DEV_IMAGE)"
DOCKER_RUN_VARNA := docker run $(INTEGRATION_OPTS) -it $(DOCKER_RUN_OPTS)


all: build ## validate all checks, build linux binary, run all tests\ncross non-linux binaries
	$(DOCKER_RUN_VARNA) ./script/make.sh

binary: build ## build the linux binary
	$(DOCKER_RUN_VARNA) ./script/make.sh generate binary

build: dist
	docker build $(DOCKER_BUILD_ARGS) -t "$(VARNA_DEV_IMAGE)" -f build.Dockerfile .

dist:
	mkdir dist  

image: binary ## build a docker varna image
	docker build -t $(VARNA_IMAGE) .