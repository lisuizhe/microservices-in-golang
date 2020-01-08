default: build

SERVICES = \
	consignment-cli \
	consignment-service \
	vessel-service 

BUILD_SERVICES_TARGETS = $(foreach service, $(SERVICES), build-service-$(service))
CLEAN_SERVICES_TARGETS = $(foreach service, $(SERVICES), clean-service-$(service))

.PHONY: default build-docker build

build-service-%:
	echo Build start: $*
	cd ./$* && $(MAKE) build
	echo Build end: $*

clean-service-%:
	echo Clean start: $*
	cd ./$* && $(MAKE) clean
	echo Clean end: $*

build-docker:
	docker-compose build

build: $(BUILD_SERVICES_TARGETS) build-docker $(CLEAN_SERVICES_TARGETS)
