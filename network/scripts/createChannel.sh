#!/bin/bash

# imports  
. scripts/envVar.sh
. scripts/utils.sh

CHANNEL_NAME="$1"
DELAY="$2"
MAX_RETRY="$3"
VERBOSE="$4"
: ${CHANNEL_NAME:="medicalsystem"}
: ${DELAY:="3"}
: ${MAX_RETRY:="5"}
: ${VERBOSE:="false"}

if [ ! -d "channel-artifacts" ]; then
	mkdir channel-artifacts
fi

createChannelGenesisBlock() {
	which configtxgen
	if [ "$?" -ne 0 ]; then
		fatalln "configtxgen tool not found."
	fi
	set -x
	configtxgen -profile TwoOrgsApplicationGenesis -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
	res=$?
	{ set +x; } 2>/dev/null
  verifyResult $res "Failed to generate channel configuration transaction..."
}

createChannel() {
	setGlobals 1
	# Poll in case the raft leader is not set yet
	local rc=1
	local COUNTER=1
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
		sleep $DELAY
		set -x
		osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" >&log.txt
		osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7153 --ca-file "$ORDERER_CA" --client-cert "$ORDERER2_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER2_ADMIN_TLS_PRIVATE_KEY" >&log.txt
		osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:27053 --ca-file "$ORDERER2_CA" --client-cert "$ORDEREREXAMPLE2_ADMIN_TLS_SIGN_CERT" --client-key "$ORDEREREXAMPLE2_ADMIN_TLS_PRIVATE_KEY" >&log.txt
		osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:27153 --ca-file "$ORDERER2_CA" --client-cert "$ORDERER2EXAMPLE2_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER2EXAMPLE2_ADMIN_TLS_PRIVATE_KEY" >&log.txt
		osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:27253 --ca-file "$ORDERER2_CA" --client-cert "$ORDERER3EXAMPLE2_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER3EXAMPLE2_ADMIN_TLS_PRIVATE_KEY" >&log.txt
		res=$?
		{ set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "Channel creation failed"
}

# joinChannel ORG
joinChannel() {
  FABRIC_CFG_PATH=$PWD/../config/
  ORG=$1
  PEER=$2
  setGlobals $ORG $PEER
	local rc=1
	local COUNTER=1
	## Sometimes Join takes time, hence retry
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
    sleep $DELAY
    set -x
    peer channel join -b $BLOCKFILE >&log.txt
    res=$?
    { set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "After $MAX_RETRY attempts, peer0.org${ORG} has failed to join channel '$CHANNEL_NAME' "
}


setAnchorPeer() {
  ORG=$1
  docker exec $(docker ps -q -f name=cli) ./scripts/setAnchorPeer.sh $ORG $CHANNEL_NAME 
}

FABRIC_CFG_PATH=${PWD}/configtx

## Create channel genesis block
infoln "Generating channel genesis block '${CHANNEL_NAME}.block'"
createChannelGenesisBlock

FABRIC_CFG_PATH=$PWD/../config/
BLOCKFILE="./channel-artifacts/${CHANNEL_NAME}.block"

## Create channel
infoln "Creating channel ${CHANNEL_NAME}"
createChannel
successln "Channel '$CHANNEL_NAME' created"

## Join all the peers to the channel
infoln "Joining org1 peer to the channel..."
joinChannel 1 0
infoln "Joining org2 peer to the channel..."
joinChannel 2 0
infoln "Joining org3 peer to the channel..."
joinChannel 3 0
infoln "Joining org4 peer to the channel..."
joinChannel 4 0

infoln "Joining org1 peer 1 to the channel..."
joinChannel 1 1
infoln "Joining org2 peer 1 to the channel..."
joinChannel 2 1
infoln "Joining org3 peer 1 to the channel..."
joinChannel 3 1
infoln "Joining org4 peer 1 to the channel..."
joinChannel 4 1

## Set the anchor peers for each org in the channel
infoln "Setting anchor peer for org1..."
setAnchorPeer 1
infoln "Setting anchor peer for org2..."
setAnchorPeer 2
infoln "Setting anchor peer for org3..."
setAnchorPeer 3
infoln "Setting anchor peer for org4..."
setAnchorPeer 4


successln "Channel '$CHANNEL_NAME' joined"
