echo removing exited containers...
docker rm $(docker ps -a -q)

IMAGE=hellofresh-hello-scratch
# for investigation, go into container shell
#docker run --name=hellofresh --rm -it --entrypoint=/bin/sh $IMAGE

echo start docker app and stay in background...
docker run --name=hellofresh -d $IMAGE tail -f /dev/null
echo start docker app up and running

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

