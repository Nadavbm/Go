TAG='envars'

build-project:
	cd $$(pwd) && docker build -t ${TAG} .

run-docker:
	docker run -ti -p 8080:8080 \
		${TAG} /bin/sh