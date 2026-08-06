DB_URL=postgres://postgres:123456@localhost:5433/chatapp?sslmode=disable
MIGRATION_DIR=./db/migrations
NAMESPACE=chat-app

up:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up

down:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down

migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force 1

run:
	go fmt ./...
	golangci-lint run ./...
	go run main.go
	
# -------------------------
# Kubernetes
# -------------------------

k8s-deploy:
	kubectl apply -f k8s/

k8s-delete:
	kubectl delete -f k8s/

pods:
	kubectl get pods -n $(NAMESPACE)

services:
	kubectl get svc -n $(NAMESPACE)

deployments:
	kubectl get deployments -n $(NAMESPACE)

ingress:
	kubectl get ingress -n $(NAMESPACE)

logs:
	kubectl logs deployment/backend -n $(NAMESPACE)

describe:
	kubectl describe pod -n $(NAMESPACE)

# -------------------------
# Minikube
# -------------------------

minikube-start:
	minikube start

minikube-stop:
	minikube stop

minikube-ip:
	minikube ip

ingress-enable:
	minikube addons enable ingress


# 	kubectl port-forward -n chat-app svc/postgres 15432:5432
# kubectl get svc -n chat-app
# kubectl get pods -n chat-app