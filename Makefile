TAG='projectName'

build-project:
	cd $$(pwd) && docker build -t ${TAG} .

run-docker:
	docker --rm -ti -p 8080:8080 \
		-v $$(pwd)/bin:/app/bin \
		-v $$(pwd)/src:/app/src \
		${TAG} bash


