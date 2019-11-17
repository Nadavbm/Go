TAG=simpleapiserver

build-api:
	cd $$(pwd) && docker build -t ${TAG} .

run-api:
	docker run -d -p 8080:8080 ${TAG}:latest


