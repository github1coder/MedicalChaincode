# 测试网络下-测试命令组-网络与身份配置

# export MedPATH="/mnt/d/-MYDISK-/go/src/hyperledger/fabric-samples/MedicalChaincode"
# export MedPATH="/Users/duxiaotu/GolandProjects/MedicalChaincode/MedicalChaincode"
export MedPATH="/mnt/d/-MYDISK-/go/src/hyperledgerfabric/fabric-samples/MedicalChaincode/MedicalChaincode"
# export testnetwork="/Users/duxiaotu/go/src/github.com/hyperledger/fabric-release-2.5/scripts/fabric-samples/test-network"
export testnetwork="/mnt/d/-MYDISK-/go/src/hyperledgerfabric/fabric-samples/test-network"
cd $testnetwork

# 启动网络
./network.sh down
./network.sh up createChannel -ca -s couchdb
# ./network.sh up createChannel -ca
# 部署链码
./network.sh deployCC -ccn MedicalChaincode -ccp $MedPATH/main -ccl go -cccg $MedPATH/collections_config.json -ccep "OR('Org1MSP.peer','Org2MSP.peer')"
# ./network.sh deployCC -ccn MedicalChaincode -ccp $MedPATH/main -ccl go

