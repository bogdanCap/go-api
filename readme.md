docker compose up --build
docker build --target production -t my-go-app:prod .

start:
    docker-compose up -d --build
delete:
    docker-compose down
    docker system prune
bash:
    docker exec -ti 7f8c4fcad43f /bin/bash
live logs:
    docker logs -f 7f8c4fcad43f

go dependency:
    go get github.com/gin-gonic/gin


http://localhost:8080/
http://localhost:8080/products

k9s:
    minikube delete

    minikube start --mount --mount-string="/var/www/projects/:/projects" --memory 4096 --cpus 2

    eval $(minikube -p minikube docker-env)


    minikube ssh
    cd /projects/go
    docker build -t go/test:1.0 -f Dockerfile .
    docker build --target development -t go/test:1.0 -f Dockerfile .

    if need to use docker hub:
        docker build -t yourusername/go-app:latest .
        docker push yourusername/go-app:latest


    kubectl apply -f postgres.yaml -f deployment.yaml -f service.yaml 
    if all yaml in folde:
        kubectl apply -f k8s/ 

    kubectl get deployments 

    or

    kubectl create deployment go-test --image=go/test:1.0 (this do not need)
    kubectl expose deployment go-test --type=NodePort --port=8080


    restart pod:
        kubectl rollout restart deployment/go-test



Payload structure layer:
                JSON

                  |
                  v

          Controller DTO

                  |
                  v

              Mapper

                  |
                  v

          Domain Order
        (business rules)

                  |
                  v

        Application Service

                  |
                  v

        Repository Interface

                  |
                  v

       PostgreSQL Implementation

                  |
                  v

              Database

List endpoint structure:
        Database rows
        |
        v
        Repository
        |
        v
        Domain objects
        |
        v
        Application service
        |
        v
        Mapper
        |
        v
        Response DTO
        |
        v
        HTTP JSON response