#!/bin/bash

set -e
set -x

export $(./setOrgEnv.sh Org1 | xargs)

export FABRIC_CFG_PATH=$PWD/../config/
export PATH=${PWD}/../bin:$PATH 

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ORG1_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ORG2_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_ORG1_ADDRESS=localhost:7051
export CORE_PEER_ORG2_ADDRESS=localhost:9051
export ORDERER_ADDRESS=orderer.example.com

CHANNEL_NAME=medicalsystem
MEDICALDATA_CHAINCODE_NAME=medicaldata
COMPUTETOKEN_CHAINCODE_NAME=computationtoken
EXAMPLEALGHORYTMM_CHAINCODE_NAME=examplealgorithm
MEDICALDATA_CHAINCODE_LOCATION=../chaincode-sources/chaincode-medicaldata
COMPUTETOKEN_CHAINCODE_LOCATION=../chaincode-sources/chaincode-computationtoken
EXAMPLEALGHORYTMM_CHAINCODE_LOCATION=../chaincode-sources/chaincode-examplealgorithm

./network.sh createChannel -c $CHANNEL_NAME
./network.sh deployCC -ccn $MEDICALDATA_CHAINCODE_NAME -ccp $MEDICALDATA_CHAINCODE_LOCATION -ccl go -c $CHANNEL_NAME
./network.sh deployCC -ccn $COMPUTETOKEN_CHAINCODE_NAME -ccp $COMPUTETOKEN_CHAINCODE_LOCATION -ccl go -c $CHANNEL_NAME
./network.sh deployCC -ccn $EXAMPLEALGHORYTMM_CHAINCODE_NAME -ccp $EXAMPLEALGHORYTMM_CHAINCODE_LOCATION -ccl go -c $CHANNEL_NAME

../bin/peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $MEDICALDATA_CHAINCODE_NAME --peerAddresses $CORE_PEER_ORG1_ADDRESS --tlsRootCertFiles $CORE_PEER_ORG1_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ORG2_ADDRESS --tlsRootCertFiles $CORE_PEER_ORG2_TLS_ROOTCERT_FILE -c '{"function":"MedicalDataSmartContract:InitLedger","Args":[]}'

../bin/peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $MEDICALDATA_CHAINCODE_NAME --peerAddresses $CORE_PEER_ORG1_ADDRESS --tlsRootCertFiles $CORE_PEER_ORG1_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ORG2_ADDRESS --tlsRootCertFiles $CORE_PEER_ORG2_TLS_ROOTCERT_FILE -c '{"function":"PatientSmartContract:InitLedger","Args":[]}'

