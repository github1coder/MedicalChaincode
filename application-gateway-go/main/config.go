package main

const (
	mspID = "Org1MSP"
	// 根据本地路径，配置test-network路径
	testNetworkPath = "/Users/duxiaotu/go/src/github.com/hyperledger/fabric-release-2.5/scripts/fabric-samples/test-network"
	cryptoPath      = testNetworkPath + "/organizations/peerOrganizations/org1.example.com"
	certPath        = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	keyPath         = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath     = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	// 连接本地或远程端口
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)
