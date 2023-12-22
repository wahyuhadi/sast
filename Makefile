docker-build: 
	docker buildx build --file Dockerfile . -t code-scanning
