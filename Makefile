# This should match the username used on your registry
USERNAME=moatorres

# You can rename this to whatever you want
PROJECT=infrago

# TAG will be filled by the output of next command
TAG=

# Set TAG to the shorter SHA256 version of the latest commit
$(eval TAG := $(shell git rev-parse --short HEAD))

# 
# Docker
# 

build:
	docker buildx build -f Dockerfile . -t ghcr.io/${USERNAME}/${PROJECT}:$(value TAG)
	docker buildx build -f Dockerfile . -t docker.io/${USERNAME}/${PROJECT}:$(value TAG)

push: 
	docker push docker.io/${USERNAME}/${PROJECT}:$(value TAG)
	docker push ghcr.io/${USERNAME}/${PROJECT}:$(value TAG)

run:
	docker run --name ${PROJECT} -it -d -p 3000:3000 ${USERNAME}/${PROJECT}:$(value TAG)

exec:
	docker exec -t -i ${PROJECT} /bin/ash

explore:
	docker run --rm -it --entrypoint=/bin/ash ${PROJECT}

stop:
	docker stop ${PROJECT}
	
remove:
	docker remove ${PROJECT}

# 
# Kubernetes
# 

DEV_MANIFESTS=$(wildcard k8s/dev/*.yaml)

deploy-dev:
	for file in $(DEV_MANIFESTS); do kubectl apply -f $$file; done

cleanup-dev:
	for file in $(DEV_MANIFESTS); do kubectl delete -f $$file; done

PRD_MANIFESTS=$(wildcard k8s/production/*.yaml)

deploy-prd:
	for file in $(PRD_MANIFESTS); do kubectl apply -f $$file; done

cleanup-prd:
	for file in $(PRD_MANIFESTS); do kubectl delete -f $$file; done

