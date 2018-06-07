APEXUP=$(shell up version 2> /dev/null)
DEP=$(shell dep version 2> /dev/null)

.PHONY: check-apex
check-apex:
ifeq "$(APEXUP)" ""
	$(error "You don't have apex/up installed, check here https://up.docs.apex.sh/#installation")
endif

.PHONY: check-dep
check-dep:
ifeq "$(DEP)" ""
	$(error "You don't have dep installed, installation: https://golang.github.io/dep/")
endif

.PHONY: init
init: check-apex check-dep vendor up.json

vendor:
	@dep ensure

up.json:
	@up

