export workingDir=$PWD
cd ./crypto-config/peerOrganizations/public.secc.com/ca && export SECC_CA_KEY=$(find *_sk)
cd $workingDir
docker-compose -f docker-compose-cli.yaml up -d
echo "########################################"
echo "           CREATING NETWORK             "
echo "########################################"
sleep 30
docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer channel create -o orderer0.secc.com:7050 -c public.secc -f ./channel-artifacts/publicchannel.tx
sleep 2
docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer channel join -b public.secc.block
sleep 2
docker exec -e "CORE_PEER_ADDRESS=peer1.public.secc.com:7051" cli peer channel join -b public.secc.block
sleep 2

docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer channel update -o orderer0.secc.com:7050 -c public.secc -f ./channel-artifacts/SeccPeerOrgAnchor.tx
sleep 2

docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer chaincode install -n secc -v 1.0 -p secc/private/chaincode/
sleep 1
docker exec -e "CORE_PEER_ADDRESS=peer1.public.secc.com:7051" cli peer chaincode install -n secc -v 1.0 -p secc/private/chaincode/
sleep 1
docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer chaincode instantiate -o orderer0.secc.com:7050 -C public.secc -n secc -v 1.0 -c '{"Args":[]}' -P "AND ('SECCPeerOrgMSP.peer')"


echo "########################################"
echo "         CHAINCODE INSTALLED            "
echo "########################################"
sleep 2

# docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer chaincode invoke -o orderer0.secc.com:7050 -C public.secc -n secc -c '{"Args":["initLedger"]}'

# ./bin/figlet "Sample function invoked. Check couchdb"



