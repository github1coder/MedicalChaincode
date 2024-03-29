package main

// 导入所需的包
import (
	// "encoding/json"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	// "time"
)

func (mc *MedicalChaincode) GetHospitalMSPID(ctx contractapi.TransactionContextInterface) (string, error) {
	hospitalMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "nil", fmt.Errorf("获取MSPID失败: %v", err)
	}
	return hospitalMSPID, nil
}

func (mc *MedicalChaincode) GetCollectionName(ctx contractapi.TransactionContextInterface) (string, error) {
	/** Get the MSP ID of submitting client identity */
	clientMSPID, err := mc.GetHospitalMSPID(ctx)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	/** Verify */
	err = mc.VerifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	/** Create the collection name */
	orgCollection := clientMSPID + "PrivateMedicalCollection"
	return orgCollection, nil
}

// 只有当前节点的客户端才能通过验证，因为节点存储着私有数据
func (mc *MedicalChaincode) VerifyClientOrgMatchesPeerOrg(ctx contractapi.TransactionContextInterface) error {
	clientMSPID, err := mc.GetHospitalMSPID(ctx)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	peerMSPID, err := shim.GetMSPID()
	if err != nil {
		return fmt.Errorf("failed getting the peer's MSPID: %v", err)
	}

	if clientMSPID != peerMSPID {
		return fmt.Errorf("client from org %v is not authorized to read or write private data from an org %v peer", clientMSPID, peerMSPID)
	}
	return nil
}

func (mc *MedicalChaincode) KeyCombine(ctx contractapi.TransactionContextInterface, index string, collection string) string {
	return collection + " ID:" + index
}

// // 提交事务
// func (mc *MedicalChaincode) SubmitTransaction(ctx contractapi.TransactionContextInterface,
// 	index string, clientMSPID string, funcType string, paraVal string, commit string) (*TxnRecord, error) {
// 	// 提交区块前数据变更不会体现在区块中，因此同一个区块内的重复事务变更不会冲突
// 	txnRecordJSON, _ := ctx.GetStub().GetState("txn" + index)
// 	if txnRecordJSON != nil {
// 		return nil, fmt.Errorf("重复提交事务: %s", index)
// 	}

// 	currentTime := time.Now()
// 	txnRecord := TxnRecord{
// 		Index:       index,
// 		FUNCTION:    funcType,
// 		PARAMETER:   paraVal,
// 		ClientMSPID: clientMSPID,
// 		TIME:        currentTime.Format("2006-01-02 15:04:05"),
// 		SUCCESS:     "submiting or aborted",
// 	}
// 	// 目前尚不能完成事务状态的转变，因为如果失败会直接挂掉、之前写入的submiting事务也不会被写入区块
// 	if commit == "commit" {
// 		txnRecord.SUCCESS = "commited"
// 	}

// 	RecordJSON, err := json.Marshal(txnRecord)
// 	if err != nil {
// 		return nil, fmt.Errorf("转换json失败: %v", err)
// 	}
// 	err = ctx.GetStub().PutState("txn"+txnRecord.Index, RecordJSON)
// 	if err != nil {
// 		return nil, fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return &txnRecord, nil
// }

// // 读取事务
// func (mc *MedicalChaincode) GetTxnRecord(ctx contractapi.TransactionContextInterface, index string) (*TxnRecord, error) {
// 	txnRecordJSON, err := ctx.GetStub().GetState("txn" + index)
// 	if err != nil {
// 		return nil, fmt.Errorf("读取事务数据库失败: %v", err)
// 	}

// 	if txnRecordJSON == nil {
// 		return nil, fmt.Errorf("不存在该事务记录： %s", index)
// 	}

// 	var txnRecord TxnRecord
// 	err = json.Unmarshal(txnRecordJSON, &txnRecord)
// 	if err != nil {
// 		return nil, fmt.Errorf("解析事务记录json失败: %v", err)
// 	}

// 	return &txnRecord, nil
// }


