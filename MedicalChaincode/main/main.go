package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MedicalChaincode struct {  // 自定义合约结构体
	/* contractapi.Contract是一个结构体，其内包含：
	 * BeforeTransaction,  // 在交易函数执行前被调用
	 * AfterTransaction,   // 在交易函数执行后被调用
	 * UnknownTransaction, // 当交易函数不存在时会默认调用UnknownTransaction
	 * TransactionContextHandler, // 合约的当前交易内容集
	 * etc.
	 */
	contractapi.Contract
}


// 主函数
func main() {
	chaincode, err := contractapi.NewChaincode(&MedicalChaincode{})
	if err != nil {
		fmt.Printf("创建医疗链码失败: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("启动医疗链码失败: %v", err)
	}
}
