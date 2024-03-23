package main

// 导入所需的包
import (
	"encoding/json"
	"reflect"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "strings"
	// "time"
)

func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, index string, RecordJSON string) error {
	DigestJSON := []byte(RecordJSON)
	index = mc.KeyCombine(ctx, index, Public)
	// err := ctx.GetStub().PutState(index, DigestJSON)
	err := ctx.GetStub().PutPrivateData(Public, index, DigestJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, index string) (string, error) {
	// DigestJSON, err := ctx.GetStub().GetState(index)
	index = mc.KeyCombine(ctx, index, Public)
	DigestJSON, err := ctx.GetStub().GetPrivateData(Public, index)
	if err != nil {
		return "nil", fmt.Errorf("读取状态数据库失败: %v", err)
	}
	if DigestJSON == nil {
		return "nil", fmt.Errorf("不存在该病人的医疗记录： %s", index)
	}

	return string(DigestJSON), nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeleteMedicalRecord(ctx contractapi.TransactionContextInterface, index string) error {
	// err := ctx.GetStub().DelState(index)
	index = mc.KeyCombine(ctx, index, Public)
	err := ctx.GetStub().DelPrivateData(Public, index)
	if err != nil {
		return fmt.Errorf("删除医疗记录失败: %v", err)
	}

	return nil
}

// 更新医疗记录，根据传入的index、字段名和新值修改medicalRecord
func (mc *MedicalChaincode) UpdateMedicalRecordByField(ctx contractapi.TransactionContextInterface, index string, field string, newValue string) error {
	// 获取当前的医疗记录
	// digestRecord, err := mc.GetMedicalRecord(ctx, index)
	index = mc.KeyCombine(ctx, index, Public)
	digestRecord, err := ctx.GetStub().GetPrivateData(Public, index)
	if err != nil {
		return fmt.Errorf("获取医疗记录失败: %v", err)
	}

	/** hao's note: update函数只能修改摘要中字段 */

	// 使用反射来更新指定字段的值
	rv := reflect.ValueOf(digestRecord).Elem()
	rv.FieldByName(field).SetString(newValue)

	// 转换回 JSON
	updatedRecordJSON, err := json.Marshal(digestRecord)
	if err != nil {
		return fmt.Errorf("转换json失败: %v", err)
	}

	// 存回链码状态
	// err = ctx.GetStub().PutState(index, updatedRecordJSON)
	err = ctx.GetStub().PutPrivateData(Public, index, updatedRecordJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

