GO = go

GOFILES = $(shell find . -name '*.go')

EXCLUDED = ./.git ./.git/% ./vendor ./vendor/% ./cmd ./cmd/%
FOLDERS = $(shell find . -mindepth 1 -type d)
PACKAGES = $(filter-out $(EXCLUDED), $(FOLDERS))

all: wswserv # client
	@echo ${PACKAGES}
	@echo $(FOLDERS)

wswserv: $(GOFILES)
	${GO} build cmd/server/wswserv.go

test: ${PACKAGES:%=%_test}

%_test:
	@${GO} test ./$*
