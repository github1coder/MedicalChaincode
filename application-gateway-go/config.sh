# 测试网络下-测试命令组-网络与身份配置
# 两个对等节点的端口分别分配为7054和8054

# export MedPATH="/mnt/d/-MYDISK-/go/src/hyperledger/fabric-samples/MedicalChaincode"
export MedPATH="/Users/duxiaotu/GolandProjects/MedicalChaincode/MedicalChaincode"
export testnetwork="/Users/duxiaotu/go/src/github.com/hyperledger/fabric-release-2.5/scripts/fabric-samples/test-network"
cd $testnetwork

# 启动网络
./network.sh down
./network.sh up createChannel -ca -s couchdb
# ./network.sh up createChannel -ca
# 部署链码
./network.sh deployCC -ccn MedicalChaincode -ccp $MedPATH/main -ccl go

