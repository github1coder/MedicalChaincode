package main

// 导入所需的包
import (
	"encoding/json"
	"reflect"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "strings"
	"time"
)

func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, Index string, RecordJSON string) (*MedicalRecord, error) {
	hospitalMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("获取MSPID失败: %v", err)
	}
	currentTime := time.Now()

	medicalRecordJSON := []byte(RecordJSON)
	err = ctx.GetStub().PutState(Index, medicalRecordJSON)
	if err != nil {
		return nil, fmt.Errorf("存入状态数据库失败: %v", err)
	}
	var medicalRecord MedicalRecord
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	_, err = mc.SubmitTransaction(ctx, hospitalMSPID+" "+currentTime.Format("2006-01-02 15:04:05")+" add"+Index, hospitalMSPID, "AddMedicalRecord", Index, "commit")
	if err != nil {
		return nil, err
	}
	return &medicalRecord, nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, index string) (*MedicalRecord, error) {
	hospitalMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("获取MSPID失败: %v", err)
	}
	currentTime := time.Now()

	medicalRecordJSON, err := ctx.GetStub().GetState(index)
	if err != nil {
		return nil, fmt.Errorf("读取状态数据库失败: %v", err)
	}
	// TODO: 在不存在该键值记录时GetStub().GetState()会直接返回下面报错，不清楚为什么，不过目前影响不大
	// Error: endorsement failure during query. response: status:500 message:"\344\270\215 ......
	if medicalRecordJSON == nil {
		return nil, fmt.Errorf("不存在该病人的医疗记录： %s", index)
	}

	var medicalRecord MedicalRecord
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	_, err = mc.SubmitTransaction(ctx, hospitalMSPID+" "+currentTime.Format("2006-01-02 15:04:05")+" get"+index, hospitalMSPID, "GetMedicalRecord", index, "commit")
	if err != nil {
		return nil, err
	}
	return &medicalRecord, nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeleteMedicalRecord(ctx contractapi.TransactionContextInterface, index string) error {
	hospitalMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("获取MSPID失败: %v", err)
	}
	currentTime := time.Now()

	err = ctx.GetStub().DelState(index)
	if err != nil {
		return fmt.Errorf("删除医疗记录失败: %v", err)
	}

	_, err = mc.SubmitTransaction(ctx, hospitalMSPID+" "+currentTime.Format("2006-01-02 15:04:05")+" delete"+index, hospitalMSPID, "GetMedicalRecord", index, "commit")
	if err != nil {
		return err
	}
	return nil
}

// 更新医疗记录，根据传入的index、字段名和新值修改medicalRecord
func (mc *MedicalChaincode) UpdateMedicalRecordByField(ctx contractapi.TransactionContextInterface, index string, field string, newValue string) error {
	hospitalMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("获取MSPID失败: %v", err)
	}
	currentTime := time.Now()

	// 获取当前的医疗记录
	medicalRecord, err := mc.GetMedicalRecord(ctx, index)
	if err != nil {
		return fmt.Errorf("获取医疗记录失败: %v", err)
	}

	// 使用反射来更新指定字段的值
	rv := reflect.ValueOf(medicalRecord).Elem()
	rv.FieldByName(field).SetString(newValue)

	// 转换回 JSON
	updatedRecordJSON, err := json.Marshal(medicalRecord)
	if err != nil {
		return fmt.Errorf("转换json失败: %v", err)
	}

	// 存回链码状态
	err = ctx.GetStub().PutState(index, updatedRecordJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	_, err = mc.SubmitTransaction(ctx, hospitalMSPID+" "+currentTime.Format("2006-01-02 15:04:05")+" update"+index+field, hospitalMSPID, "UpdateMedicalRecordByField", index+" "+field+" "+newValue, "commit")
	if err != nil {
		return err
	}
	return nil
}

// // 更新医疗记录
// func (mc *MedicalChaincode) UpdateMedicalRecord(ctx contractapi.TransactionContextInterface, patient, newDiagnosis string) error {
// 	existingRecord, err := mc.GetMedicalRecord(ctx, patient)
// 	if err != nil {
// 		return fmt.Errorf("获取医疗记录失败: %v", err)
// 	}

// 	existingRecord.Note_Details = newDiagnosis

// 	updatedRecordJSON, err := json.Marshal(existingRecord)
// 	if err != nil {
// 		return fmt.Errorf("转换json失败: %v", err)
// 	}
// 	err = ctx.GetStub().PutState(patient, updatedRecordJSON)
// 	if err != nil {
// 		return fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return nil
// }
