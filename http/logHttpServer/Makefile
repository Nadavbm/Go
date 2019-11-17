TAG=loghttpserver

build-http:
	cd $$(pwd) && docker build -t ${TAG} .

run-http:
	docker run -d -p 8080:8080 ${TAG}:latest


