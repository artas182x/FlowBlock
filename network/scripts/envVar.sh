#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

# imports
. scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER2_CA=${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/msp/tlscacerts/tlsca.example2.com-cert.pem
export PEER0_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export PEER0_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export PEER0_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export PEER0_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt
export PEER1_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
export PEER1_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
export PEER1_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/ca.crt
export PEER1_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/ca.crt
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key
export ORDERER2_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.crt
export ORDERER2_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.key
export ORDEREREXAMPLE2_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/server.crt
export ORDEREREXAMPLE2_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/server.key

# Set environment variables for the peer org
setGlobals() {
  local USING_ORG=""
  local USING_PEER=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi

  if [ -z "$2" ]; then
    USING_PEER=0
  else
    USING_PEER=$2
  fi

  infoln "Using organization ${USING_ORG}"
  infoln "Using peer ${USING_PEER}"

  if [ $USING_PEER -eq 0 ]; then
    if [ $USING_ORG -eq 1 ]; then
      export CORE_PEER_LOCALMSPID="Org1MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
      export CORE_PEER_ADDRESS=localhost:7051
    elif [ $USING_ORG -eq 2 ]; then
      export CORE_PEER_LOCALMSPID="Org2MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
      export CORE_PEER_ADDRESS=localhost:9051

    elif [ $USING_ORG -eq 3 ]; then
      export CORE_PEER_LOCALMSPID="Org3MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
      export CORE_PEER_ADDRESS=localhost:11051
    elif [ $USING_ORG -eq 4 ]; then
      export CORE_PEER_LOCALMSPID="Org4MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG4_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp
      export CORE_PEER_ADDRESS=localhost:12051
    else
      errorln "ORG Unknown"
    fi
  elif [ $USING_PEER -eq 1 ]; then
    if [ $USING_ORG -eq 1 ]; then
      export CORE_PEER_LOCALMSPID="Org1MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG1_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
      export CORE_PEER_ADDRESS=localhost:7151
    elif [ $USING_ORG -eq 2 ]; then
      export CORE_PEER_LOCALMSPID="Org2MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG2_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
      export CORE_PEER_ADDRESS=localhost:9151

    elif [ $USING_ORG -eq 3 ]; then
      export CORE_PEER_LOCALMSPID="Org3MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG3_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
      export CORE_PEER_ADDRESS=localhost:11151
    elif [ $USING_ORG -eq 4 ]; then
      export CORE_PEER_LOCALMSPID="Org4MSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG4_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp
      export CORE_PEER_ADDRESS=localhost:12151
    else
      errorln "ORG Unknown"
    fi
  else
    errorln "PEER Unknown"
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# Set environment variables for use in the CLI container 
setGlobalsCLI() {
  setGlobals $1

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  if [ $USING_ORG -eq 1 ]; then
    export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_ADDRESS=peer0.org3.example.com:11051
  elif [ $USING_ORG -eq 4 ]; then
    export CORE_PEER_ADDRESS=peer0.org4.example.com:12051
  else
    errorln "ORG Unknown"
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  PEER=$1
  shift
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    PEER="peer${PEER}.org$1"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	PEERS="$PEER"
    else
	PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    CA=PEER0_ORG$1_CA
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
    # shift by one to get to the next organization
    shift
  done
}


verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}
