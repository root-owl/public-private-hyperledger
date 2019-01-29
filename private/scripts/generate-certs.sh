# F O R   C E R T I F I C A T E S
rm -rf ./crypto-config
mkdir ./crypto-config
cryptogen generate --config=./crypto-config.yaml --output=./crypto-config

# F O R   C H A N N E L   A R T I F A C T S
mkdir ./channel-artifacts
export FABRIC_CFG_PATH=$PWD
configtxgen -profile SECCPublicOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
export CHANNEL_NAME=public.secc  && configtxgen -profile SECCPublicChannel -outputCreateChannelTx ./channel-artifacts/publicchannel.tx -channelID $CHANNEL_NAME
configtxgen -profile SECCPublicChannel -outputAnchorPeersUpdate ./channel-artifacts/SeccPeerOrgAnchor.tx -channelID $CHANNEL_NAME -asOrg SECCPeerOrg
echo "########################################"
echo "        CERTIFICATES GENERATED          "
echo "########################################"
