# rm -rf ./channel-artifacts
# rm -rf ./crypto-config
# ./bin/figlet "Old Certificates Removed"

docker kill $(docker ps -q)
docker rm $(docker ps -qa)
echo y | docker volume prune
echo y | docker network prune
docker rmi $(docker images dev-* -q)
echo "########################################"
echo "             DOCKER CLEANED             "
echo "########################################"

