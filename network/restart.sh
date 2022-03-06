#!/bin/bash

./network.sh down && ./network.sh up && ./installChaincode.sh
cp ./organizations/peerOrganizations/org1.example.com/connection-org1.yaml ../web/network1.yaml && chmod 644 ../web/network1.yaml
cp ./organizations/peerOrganizations/org2.example.com/connection-org2.yaml ../web/network2.yaml && chmod 644 ../web/network2.yaml
cp ./organizations/peerOrganizations/org3.example.com/connection-org3.yaml ../web/network3.yaml && chmod 644 ../web/network3.yaml
cp ./organizations/peerOrganizations/org4.example.com/connection-org4.yaml ../web/network4.yaml && chmod 644 ../web/network4.yaml