TAG=""
# TAG=$(git rev-parse --short HEAD)
$(eval TAG := $(shell git rev-parse --short HEAD))

build:
	docker buildx build -f Dockerfile . -t docker.io/moatorres/infrago:$(value TAG)

push: 
	docker push docker.io/moatorres/infrago:$(value TAG)

run:
	docker run --name infrago -it -d -p 3000:3000 moatorres/infrago:$(value TAG)

exec:
	docker exec -t -i infrago /bin/ash

explore:
	docker run --rm -it --entrypoint=/bin/ash infrago

stop:
	docker stop infrago
	
remove:
	docker remove infrago

deploy-prd:
	kubectl apply -f k8s/production/deployment.yaml
	kubectl apply -f k8s/production/ingress.yaml

cleanup-prd:
	kubectl delete -f k8s/production/deployment.yaml
	kubectl delete -f k8s/production/ingress.yaml

deploy-dev:
	kubectl apply -f k8s/dev/deployment.yaml
	kubectl apply -f k8s/dev/ingress.yaml

cleanup-dev:
	kubectl delete -f k8s/dev/deployment.yaml
	kubectl delete -f k8s/dev/ingress.yaml