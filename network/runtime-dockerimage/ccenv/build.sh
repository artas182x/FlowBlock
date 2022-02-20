#!/bin/bash
docker rmi hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3 hyperledger/fabric-ccenv:latest
docker build -t hyperledger/fabric-ccenv:2.3.3 .
docker tag hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3
docker tag hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:latest
