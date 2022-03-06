#!/bin/bash

function createOrg1() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org1.example.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.example.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering peer1"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient 1"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name patient1 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient 2"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name patient2 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient 3"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name patient3 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
  
  infoln "Registering doctor 1"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name doctor1 --id.secret 123 --id.type client --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering doctor 2"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name doctor2 --id.secret 123 --id.type client --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering doctor 3"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name university1 --id.secret 123 --id.type client --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name org1admin --id.secret org1adminpw --id.type admin --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp" --csr.hosts peer0.org1.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1 msp"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/msp" --csr.hosts peer1.org1.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls" --enrollment.profile tls --csr.hosts peer0.org1.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls" --enrollment.profile tls --csr.hosts peer1.org1.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key"

  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/server.key"


  mkdir -p "${PWD}/organizations/peerOrganizations/org1.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/msp/tlscacerts/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/msp/tlscacerts/ca.crt"

  mkdir -p "${PWD}/organizations/peerOrganizations/org1.example.com/tlsca"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem"


  mkdir -p "${PWD}/organizations/peerOrganizations/org1.example.com/ca"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem"


  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://patient1:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient1@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
    fabric-ca-client enroll -u https://patient2:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient2@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
    fabric-ca-client enroll -u https://patient3:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient3@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
    fabric-ca-client enroll -u https://doctor1:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/doctor1@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
    fabric-ca-client enroll -u https://doctor2:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/doctor2@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null
    fabric-ca-client enroll -u https://university1:123@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/university1@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient1@org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient2@org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/patient3@org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/doctor1@org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/doctor2@org1.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/university1@org1.example.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/config.yaml"
}

function createOrg2() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org2.example.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org2.example.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-org2 --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering peer1"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name patient11 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null
  
  infoln "Registering doctor"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name doctor11 --id.secret 123 --id.type client --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name org2admin --id.secret org2adminpw --id.type admin --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/msp" --csr.hosts peer0.org2.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1 msp"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/msp" --csr.hosts peer1.org2.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls" --enrollment.profile tls --csr.hosts peer0.org2.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls" --enrollment.profile tls --csr.hosts peer1.org2.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.key"

  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/server.key"


  mkdir -p "${PWD}/organizations/peerOrganizations/org2.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/msp/tlscacerts/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/msp/tlscacerts/ca.crt"

  mkdir -p "${PWD}/organizations/peerOrganizations/org2.example.com/tlsca"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem"


  mkdir -p "${PWD}/organizations/peerOrganizations/org2.example.com/ca"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem"

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://patient11:123@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/users/patient11@org2.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null
  fabric-ca-client enroll -u https://doctor11:123@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/users/doctor11@org2.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/users/patient11@org2.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/users/doctor11@org2.example.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org2admin:org2adminpw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/config.yaml"
}

function createOrderer2() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/ordererOrganizations/example2.com

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/example2.com

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:29054 --caname ca-orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-29054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-29054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-29054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-29054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/ordererOrganizations/example2.com/msp/config.yaml"

  infoln "Registering orderer"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the orderer admin"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the orderer msp"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:29054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/msp" --csr.hosts orderer.example2.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example2.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/msp/config.yaml"

  infoln "Generating the orderer-tls certificates"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:29054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls" --enrollment.profile tls --csr.hosts orderer.example2.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/ca.crt"
  cp "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/signcerts/"* "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/server.crt"
  cp "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/keystore/"* "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/server.key"

  mkdir -p "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/msp/tlscacerts/tlsca.example2.com-cert.pem"

  mkdir -p "${PWD}/organizations/ordererOrganizations/example2.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/example2.com/orderers/orderer.example2.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example2.com/msp/tlscacerts/tlsca.example2.com-cert.pem"

  infoln "Generating the admin msp"
  set -x
  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:29054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example2.com/users/Admin@example2.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg2/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example2.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/example2.com/users/Admin@example2.com/msp/config.yaml"
}

function createOrderer() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/ordererOrganizations/example.com

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/example.com

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml"

  infoln "Registering orderer"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

    infoln "Registering orderer 2"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer2 --id.secret ordererpw --id.type orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the orderer admin"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the orderer msp"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp" --csr.hosts orderer.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

   infoln "Generating the orderer 2 msp"
  set -x
  fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp" --csr.hosts orderer2.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/config.yaml"
  cp "${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp/config.yaml"


  infoln "Generating the orderer-tls certificates"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls" --enrollment.profile tls --csr.hosts orderer.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the orderer2-tls certificates"
  set -x
  fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls" --enrollment.profile tls --csr.hosts orderer2.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/signcerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/keystore/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key"

  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/ca.crt"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/signcerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.crt"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/keystore/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.key"

  mkdir -p "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

  mkdir -p "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

  mkdir -p "${PWD}/organizations/ordererOrganizations/example.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

  infoln "Generating the admin msp"
  set -x
  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/example.com/users/Admin@example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/example.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/example.com/users/Admin@example.com/msp/config.yaml"
}

function createOrg3 {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org3.example.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org3.example.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:11054 --caname ca-org3 --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-11054-ca-org3.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-11054-ca-org3.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-11054-ca-org3.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-11054-ca-org3.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org3 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering peer1"
  set -x
  fabric-ca-client register --caname ca-org3 --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient"
  set -x
  fabric-ca-client register --caname ca-org3 --id.name patient11 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null
  
  infoln "Registering doctor"
  set -x
  fabric-ca-client register --caname ca-org3 --id.name doctor11 --id.secret 123 --id.type client --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org3 --id.name org3admin --id.secret org3adminpw --id.type admin --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/msp" --csr.hosts peer0.org3.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1 msp"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/msp" --csr.hosts peer1.org3.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls" --enrollment.profile tls --csr.hosts peer0.org3.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls" --enrollment.profile tls --csr.hosts peer1.org3.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/server.key"

  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/server.key"


  mkdir -p "${PWD}/organizations/peerOrganizations/org3.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/msp/tlscacerts/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/msp/tlscacerts/ca.crt"

  mkdir -p "${PWD}/organizations/peerOrganizations/org3.example.com/tlsca"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem"


  mkdir -p "${PWD}/organizations/peerOrganizations/org3.example.com/ca"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/ca/ca.org3.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org3.example.com/ca/ca.org3.example.com-cert.pem"

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://patient11:123@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/users/patient11@org3.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null
  fabric-ca-client enroll -u https://doctor11:123@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/users/doctor11@org3.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org3.example.com/users/patient11@org3.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org3.example.com/users/doctor11@org3.example.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org3admin:org3adminpw@localhost:11054 --caname ca-org3 -M "${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org3/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org3.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp/config.yaml"
}

function createOrg4 {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org4.example.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org4.example.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:12054 --caname ca-org4 --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-12054-ca-org4.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-12054-ca-org4.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-12054-ca-org4.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-12054-ca-org4.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org4 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering peer1"
  set -x
  fabric-ca-client register --caname ca-org4 --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering patient"
  set -x
  fabric-ca-client register --caname ca-org4 --id.name patient11 --id.secret 123 --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering doctor"
  set -x
  fabric-ca-client register --caname ca-org4 --id.name doctor11 --id.secret 123 --id.type client --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org4 --id.name org4admin --id.secret org4adminpw --id.type admin --id.attrs ReadOthersData=1:ecert --id.attrs RequestTokenRole=1:ecert --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/msp" --csr.hosts peer0.org4.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1 msp"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/msp" --csr.hosts peer1.org4.example.com --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls" --enrollment.profile tls --csr.hosts peer0.org4.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer1-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls" --enrollment.profile tls --csr.hosts peer1.org4.example.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/server.key"

  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/server.key"


  mkdir -p "${PWD}/organizations/peerOrganizations/org4.example.com/msp/tlscacerts"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/msp/tlscacerts/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/msp/tlscacerts/ca.crt"

  mkdir -p "${PWD}/organizations/peerOrganizations/org4.example.com/tlsca"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/tlsca/tlsca.org4.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/tlsca/tlsca.org4.example.com-cert.pem"


  mkdir -p "${PWD}/organizations/peerOrganizations/org4.example.com/ca"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/ca/ca.org4.example.com-cert.pem"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/msp/cacerts/"* "${PWD}/organizations/peerOrganizations/org4.example.com/ca/ca.org4.example.com-cert.pem"

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://patient11:123@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/users/patient11@org4.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null
  fabric-ca-client enroll -u https://doctor11:123@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/users/doctor11@org4.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org4.example.com/users/patient11@org4.example.com/msp/config.yaml"
  cp "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org4.example.com/users/doctor11@org4.example.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org4admin:org4adminpw@localhost:12054 --caname ca-org4 -M "${PWD}/organizations/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org4/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org4.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp/config.yaml"
}
