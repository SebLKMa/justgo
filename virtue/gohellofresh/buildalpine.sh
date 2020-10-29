CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello hello.go

# remove exited containers and images
docker rm $(docker ps -a -q)
docker rmi $(docker images --filter "dangling=true" -q)

IMAGE=hellofresh-hello-alpine
# remove SPECIFIC image by EXACT tag name
docker rmi $(docker images $IMAGE -q)

docker build -t $IMAGE -f Dockerfile.alpine .