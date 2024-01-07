package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 添加医疗记录
// 以index为主键，medicalRecord结构体为值，存入状态数据库
func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, medicalRecordJSON string) error {
	var medicalRecord MedicalRecord

	err := json.Unmarshal([]byte(medicalRecordJSON), &medicalRecord)
	if err != nil {
		return fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	err = ctx.GetStub().PutState(fmt.Sprintf("%d", medicalRecord.index), []byte(medicalRecordJSON))
	if err != nil {
		return fmt.Errorf("医疗记录存入状态数据库失败: %v", err)
	}

	return nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, index int) (*MedicalRecord, error) {
	medicalRecordJSON, err := ctx.GetStub().GetState(fmt.Sprintf("%d", index))
	if err != nil {
		return nil, fmt.Errorf("读取状态数据库失败: %v", err)
	}
	// TODO: 在不存在该键值记录时GetStub().GetState()会直接返回下面报错，不清楚为什么，不过目前影响不大
	// Error: endorsement failure during query. response: status:500 message:"\344\270\215 ......
	if medicalRecordJSON == nil {
		return nil, fmt.Errorf("不存在该病人的医疗记录： %d", index)
	}

	var medicalRecord MedicalRecord
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	return &medicalRecord, nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeleteMedicalRecord(ctx contractapi.TransactionContextInterface, index int) error {
	err := ctx.GetStub().DelState(fmt.Sprintf("%d", index))
	if err != nil {
		return fmt.Errorf("删除医疗记录失败: %v", err)
	}
	return nil
}

//// 更新医疗记录
//func (mc *MedicalChaincode) UpdateMedicalRecord(ctx contractapi.TransactionContextInterface, patient, newDiagnosis string) error {
//	existingRecord, err := mc.GetMedicalRecord(ctx, patient)
//	if err != nil {
//		return fmt.Errorf("获取医疗记录失败: %v", err)
//	}
//
//	existingRecord.Diagnosis = newDiagnosis
//
//	updatedRecordJSON, err := json.Marshal(existingRecord)
//	if err != nil {
//		return fmt.Errorf("转换json失败: %v", err)
//	}
//	err = ctx.GetStub().PutState(patient, updatedRecordJSON)
//	if err != nil {
//		return fmt.Errorf("存入状态数据库失败: %v", err)
//	}
//
//	return nil
//}
