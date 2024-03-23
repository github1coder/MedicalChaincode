package main

/** private.go
	存放涉及私有数据'增删改查'的链码函数 */

// 导入所需的包
import (
	// "encoding/json"
	// "reflect"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "strings"
	// "time"
)

// 添加私有医疗记录
func (mc *MedicalChaincode) AddPrivateMedicalRecord(ctx contractapi.TransactionContextInterface, index string, RecordJSON string) error {
	orgCollection, err := mc.GetCollectionName(ctx)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	err = ctx.GetStub().PutPrivateData(orgCollection, mc.KeyCombine(ctx, index, orgCollection), []byte(RecordJSON)) 
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// 获取私有医疗记录
func (mc *MedicalChaincode) GetPrivateMedicalRecord(ctx contractapi.TransactionContextInterface, index string) (string, error) {
	orgCollection, err := mc.GetCollectionName(ctx)
	if err != nil {
		return "nil", fmt.Errorf("%v", err)
	}
	medicalRecordJSON, err := ctx.GetStub().GetPrivateData(orgCollection, mc.KeyCombine(ctx, index, orgCollection))
	if err != nil {
		return "nil", fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return string(medicalRecordJSON), nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeletePrivateMedicalRecord(ctx contractapi.TransactionContextInterface, index string) error {
	orgCollection, err := mc.GetCollectionName(ctx)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	err = ctx.GetStub().DelPrivateData(orgCollection, mc.KeyCombine(ctx, index, orgCollection)) 
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// // 更新医疗记录，根据传入的index、字段名和新值修改medicalRecord
// func (mc *MedicalChaincode) UpdatePrivateMedicalRecordByField(ctx contractapi.TransactionContextInterface, index string, field string, newValue string) error {
// 	// 获取当前的医疗记录
// 	medicalRecord, err := mc.GetMedicalRecord(ctx, index)
// 	if err != nil {
// 		return fmt.Errorf("获取医疗记录失败: %v", err)
// 	}

// 	// 使用反射来更新指定字段的值
// 	rv := reflect.ValueOf(medicalRecord).Elem()
// 	rv.FieldByName(field).SetString(newValue)

// 	// 转换回 JSON
// 	updatedRecordJSON, err := json.Marshal(medicalRecord)
// 	if err != nil {
// 		return fmt.Errorf("转换json失败: %v", err)
// 	}

// 	// 存回链码状态
// 	err = ctx.GetStub().PutState(index, updatedRecordJSON)
// 	if err != nil {
// 		return fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return nil
// }

