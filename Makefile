ACTUAL_VERSION=$(shell git describe --tag --abbrev=0)
TYPE?=patch
DOCKER_USER=$(shell docker info | sed '/Username:/!d;s/.* //')
KEY=1234

y=$(subst ., ,$(ACTUAL_VERSION))
major=$(word 1, $(y))
minor=$(word 2, $(y))
patch=$(word 3, $(y))
ifeq ($(TYPE),major)
	major=$(shell echo $major\+1 | bc)
endif
ifeq ($(TYPE),minor)
	minor=$(shell echo $minor\+1 | bc)
endif
ifeq ($(TYPE),patch)
	patch=$(shell echo $patch\+1 | bc)
endif

NEW_VERSION=$(major).$(minor).$(patch)

build: 
	yarn install --modules-folder./ui
	yarn build ./ui
	go build -o ./ingresso-watcher cmd/main.go

docker: 
	docker build -t ingresso-watcher:test 
	docker run -e API_KEY $(KEY) --name ingresso-watcher -p 8080 -d ingresso-watcher:test

release: new-tag release-git release-docker

release-git: 
	git add .
	git commit -m "new release v$(NEW_VERSION)"
	git tag -a $(NEW_VERSION) -m "release v$(NEW_VERSION)"
	git push

release-docker:
	docker build -t $(DOCKER_USER)/ingresso-watcher:$(NEW_VERSION) 
	docker push $(DOCKER_USER)/ingresso-watcher:$(NEW_VERSION) 

new-tag:
	@echo "$(ACTUAL_VERSION) -> $(NEW_VERSION)"
