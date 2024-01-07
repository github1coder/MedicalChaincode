## Abstract
https://github.com/github1coder/MedicalChaincode
TODO: 完成基于区块链的EHR医疗系统，数据集参照MIMICⅢ https://mimic.mit.edu

## Update
### 将medical.go按模块分成了三个文件，方便编写
main.go:    链码入口
model.go:   医疗数据结构体定义
CURD.go:    基本的增删改查
### 稍微改了下EHR格式
直接采用师兄给的MIMICⅢ数据集，没有改动，数据类型有待商榷
``` 
Unnamed: 0                                                              0
index                                                                   0
SUBJECT_ID                                                            249
GENDER                                                                  F
DOB                                                   2075-03-13 00:00:00
DOD                                                                   NaN
DOD_HOSP                                                              NaN
DOD_SSN                                                               NaN
EXPIRE_FLAG                                                             0
HADM_ID                                                            149546
ADMITTIME                                             2155-02-03 20:16:00
DISCHTIME                                             2155-02-14 11:15:00
DEATHTIME                                                             NaN
ADMISSION_TYPE                                                  EMERGENCY
ADMISSION_LOCATION                                   EMERGENCY ROOM ADMIT
DISCHARGE_LOCATION                               REHAB/DISTINCT PART HOSP
INSURANCE                                                        Medicare
LANGUAGE                                                             ENGL
RELIGION                                                         CATHOLIC
MARITAL_STATUS                                                   DIVORCED
ETHNICITY                                                           WHITE
EDREGTIME                                             2155-02-03 17:43:00
EDOUTTIME                                             2155-02-03 21:26:00
DIAGNOSIS                                            GI BLEED/ CHEST PAIN
HOSPITAL_EXPIRE_FLAG                                                    0
HAS_CHARTEVENTS_DATA                                                    1
ICUSTAY_Details         ICUSTAY_ID: 269035, FIRST_CAREUNIT: MICU, LAST...
ICD9_CODE               ['56985', '41071', '43411', '5849', '2851', '3...
PRO_CODE                                           [3974, 40, 8841, 9910]
Drug_Details            ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
Input_Details           ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
Note_Details            1495461495461495461495461495461495461495461495...
```

## Result
测试时用的命令组，可以跑通：
``` shell
./network down
./network.sh up createChannel -ca
./network.sh deployCC -ccn Hao -ccp $MedPATH/main -ccl go
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
# Org1MSP 创建身份
export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.example.com/
fabric-ca-client register --caname ca-org1 --id.name owner --id.secret ownerpw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
fabric-ca-client enroll -u https://owner:ownerpw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/owner@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/owner@org1.example.com/msp/config.yaml"
# Org2MSP 创建身份
# export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org2.example.com/
# fabric-ca-client register --caname ca-org2 --id.name buyer --id.secret buyerpw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
# fabric-ca-client enroll -u https://buyer:buyerpw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.example.com/users/buyer@org2.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/tls-cert.pem"
# cp "${PWD}/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.example.com/users/buyer@org2.example.com/msp/config.yaml"
# 以Org1 owner用户操作peer CLI
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/owner@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
# 以Org2 owner用户操作peer CLI
# export CORE_PEER_TLS_ENABLED=true
# export CORE_PEER_LOCALMSPID="Org2MSP"
# export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
# export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/buyer@org2.example.com/msp
# export CORE_PEER_ADDRESS=localhost:9051
# 调用链码
peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"AddMedicalRecord","Args":["330881","1","fever"]}'
peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'
peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"UpdateMedicalRecord","Args":["330881","just cold"]}'
peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'
peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"DeleteMedicalRecord","Args":["330881"]}'
peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'

```

## Architecture