ACTUAL_VERSION=$(shell git describe --tag --abbrev=0)
TYPE?=patch
KEY=1234

y=$(subst ., ,$(ACTUAL_VERSION))
tmp_major=$(word 1, $(y))
tmp_minor=$(word 2, $(y))
tmp_patch=$(word 3, $(y))
ifeq ($(TYPE),major)
	major=$(shell expr $(tmp_major) + 1)
else
	major=$(tmp_major)
endif
ifeq ($(TYPE),minor)
	minor=$(shell expr $(tmp_minor) + 1)
else
	minor=$(tmp_minor)
endif
ifeq ($(TYPE),patch)
	patch=$(shell expr $(tmp_patch) + 1)
else
	patch=$(tmp_patch)
endif

NEW_VERSION=$(major).$(minor).$(patch)

release: test new-tag docker-build docker-push release-git 

build: 
	yarn install --modules-folder./ui
	yarn build ./ui
	go build -o ./ingresso-watcher cmd/main.go

test:
	go test ./...

new-tag:
	@echo "$(ACTUAL_VERSION) -> $(NEW_VERSION)"

docker-build: 
	docker build . -t $(DOCKER_USER)/ingresso-watcher:$(NEW_VERSION) 

docker-push:
	docker push $(DOCKER_USER)/ingresso-watcher:$(NEW_VERSION) 

release-git: 
	git add .
	git tag -a $(NEW_VERSION) -m "new release v$(NEW_VERSION)"
	git commit -m "add changes"
	git push --follow-tags


