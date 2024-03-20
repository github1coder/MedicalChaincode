package main

// 导入所需的包
import (
	"encoding/json"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

// 提交事务
func (mc *MedicalChaincode) SubmitTransaction(ctx contractapi.TransactionContextInterface,
	index string, clientMSPID string, funcType string, paraVal string, commit string) (*TxnRecord, error) {
	// 提交区块前数据变更不会体现在区块中，因此同一个区块内的重复事务变更不会冲突
	txnRecordJSON, _ := ctx.GetStub().GetState("txn" + index)
	if txnRecordJSON != nil {
		return nil, fmt.Errorf("重复提交事务: %s", index)
	}

	currentTime := time.Now()
	txnRecord := TxnRecord{
		index:       index,
		FUNCTION:    funcType,
		PARAMETER:   paraVal,
		ClientMSPID: clientMSPID,
		TIME:        currentTime.Format("2006-01-02 15:04:05"),
		SUCCESS:     "submiting or aborted",
	}
	// 目前尚不能完成事务状态的转变，因为如果失败会直接挂掉、之前写入的submiting事务也不会被写入区块
	if commit == "commit" {
		txnRecord.SUCCESS = "commited"
	}

	RecordJSON, err := json.Marshal(txnRecord)
	if err != nil {
		return nil, fmt.Errorf("转换json失败: %v", err)
	}
	err = ctx.GetStub().PutState("txn"+txnRecord.index, RecordJSON)
	if err != nil {
		return nil, fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return &txnRecord, nil
}

// 读取事务
func (mc *MedicalChaincode) GetTxnRecord(ctx contractapi.TransactionContextInterface, index string) (*TxnRecord, error) {
	txnRecordJSON, err := ctx.GetStub().GetState("txn" + index)
	if err != nil {
		return nil, fmt.Errorf("读取事务数据库失败: %v", err)
	}

	if txnRecordJSON == nil {
		return nil, fmt.Errorf("不存在该事务记录： %s", index)
	}

	var txnRecord TxnRecord
	err = json.Unmarshal(txnRecordJSON, &txnRecord)
	if err != nil {
		return nil, fmt.Errorf("解析事务记录json失败: %v", err)
	}

	return &txnRecord, nil
}
