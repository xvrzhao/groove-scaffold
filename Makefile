registry = your.registry.host/your_repo_name
project = groove-app
image = ${registry}/${project}

version ?=
env ?= production
arch = amd64
tag = ${version}-${env}-${arch}

table ?=
model ?=

.PHONY: check image push publish api run-http run-cron

check:
ifeq ($(version),)
	$(error image version not specified)
endif
	@echo image tag: ${tag}

image: check
	docker buildx build --platform=linux/amd64 --build-arg PUBLISH_MODE=${env} -t ${image}:${tag} .

push: check
	docker push ${image}:${tag}

publish: image push
	@echo "\n-------------------------\n\nComplete! Copy image name: \n\033[3;32m${image}:${tag}\033[0m"

api:
	bin/go run ./cmd/gencode -t ${table} -m ${model}

run-http:
	bin/go run ./cmd/http

run-cron:
	bin/go run ./cmd/cron