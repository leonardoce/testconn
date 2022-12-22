.PHONY: run
run: build
	./scripts/run.sh

.PHONY: build
build:
	./scripts/build.sh

.PHONY: check
check:
	./scripts/check.sh

.PHONY: image
image:
	./scripts/image.sh

.PHONY: test
test:
	./scripts/test.sh
