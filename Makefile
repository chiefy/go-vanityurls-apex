APEXUP=$(shell up version 2> /dev/null)
AWS_PROFILE_NAME=

.PHONY: check-apex
check-apex:
ifeq "$(APEXUP)" ""
	$(error "You don't have apex/up installed, check here https://up.docs.apex.sh/#installation")
endif

up.json:
	@up

