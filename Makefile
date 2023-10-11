build:
	docker buildx build -f Dockerfile . -t docker.io/moatorres/infrago-rotational

push: 
	docker push docker.io/moatorres/infrago-rotational:latest

run:
	docker run --name docker.io/moatorres/infrago-rotational -it -d -p 3000:3000 docker.io/moatorres/infrago-rotational

exec:
	docker exec -t -i docker.io/moatorres/infrago-rotational /bin/ash

explore:
	docker run --rm -it --entrypoint=/bin/ash docker.io/moatorres/infrago-rotational

stop:
	docker stop docker.io/moatorres/infrago-rotational
	
remove:
	docker remove docker.io/moatorres/infrago-rotational

deploy-prd:
	kubectl apply -f deploy/deployment.yaml

cleanup-prd:
	kubectl delete -f deploy/deployment.yaml

deploy-dev:
	kubectl apply -f deploy/deployment.dev.yaml

cleanup-dev:
	kubectl delete -f deploy/deployment.yaml