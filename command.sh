
# 针对链码函数的测试

# 以本地的测试网络路径为准，不建议相对路径
export testnetwork="../test-network"  
# 以本地的链码目录路径为准，不建议相对路径
export MedPATH="../MedicalChaincode"

# 配置基本环境
# source config.sh

# 基本CURD测试，已弃用
# peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"\
#  -c '{"function":"AddMedicalRecord","Args":["","","","","","","","","","","","","","","","","","","","","","","","","","","","","","",""]}'
# peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'
# peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"\
#  -c '{"function":"UpdateMedicalRecord","Args":["330881","just cold"]}'
# peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'
# peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"\
#  -c '{"function":"DeleteMedicalRecord","Args":["330881"]}'
# peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","330881"]}'

################ MIMICⅢ测试集读入测试，写入联盟链 #################
cp log.csv medical.csv
# 删去第一列
sed -i 's/^[^,]*,//' medical.csv
# 删去所有的单引号，因为如Tony's这种内容会导致shell认为是单引号未匹配而不能正确执行
sed -i "s/'//g" medical.csv
# 列间分隔采用逗号以便链码区分，但因为字符串""中也有逗号因而需要特殊判断""中的逗号、但代表字符串的""留不到函数里所以改为[]、从而需要先将原先就有的[]都删去
sed -i 's/\[//g' medical.csv
sed -i 's/]//g' medical.csv
sed -i 's/"/[/g' medical.csv
# 删去第一行
sed -i '1d' medical.csv
# 删去空行
sed -i '/^\s*$/d' medical.csv
# 工作目录切换到测试网络目录，同时将处理好的数据也移过去
touch ../test-network/info.csv
cp medical.csv  $testnetwork/info.csv
cd $testnetwork
# 清空invoke.sh
echo > invoke.sh
# 将所有invoke命令都写到invoke.sh脚本里
while read line
do
    line=${line//$'\r'}
    echo "peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{\"function\":\"AddMedicalRecord\",\"Args\":[\"$line\"]}'" >> invoke.sh
    echo "sleep 1" >> invoke.sh
done < info.csv
# 执行invoke.sh脚本，添加MIMICⅢ数据集到联盟链中
source invoke.sh
# 回到最开始的目录
cd $MedPATH


# 第10条记录太长了，超出脚本对参数列表的限制。

# peer chaincode query -C mychannel -n Hao -c '{"Args":["GetMedicalRecord","1"]}'
# peer chaincode invoke -o localhost:7050 --waitForEvent --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n Hao --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"DeleteMedicalRecord","Args":["1"]}'

