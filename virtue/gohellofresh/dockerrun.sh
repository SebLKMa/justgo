# This is a sample script to run Docker container with input and output files.

echo removing exited containers...
docker rm $(docker ps -a -q)

echo start docker app and stay in background...
docker run --name=hellofresh -d hellofresh-hello-alpine tail -f /dev/null
echo start docker app up and running
# NOTE: For investigation, go into container shell with:
#   docker run --name=hellofresh --rm -it --entrypoint=/bin/sh hellofresh-hello-alpine

echo copying input file to container...
docker cp input.json hellofresh:/app/input.json

echo executing program...
docker exec -it hellofresh /app/hello

echo copying output file from container...
docker cp hellofresh:/app/output.json output.json
echo done!

echo stopping docker app...
docker stop hellofresh
echo docker app stopped

