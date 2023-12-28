package main

// 导入所需的包
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	//"Medical/fabric/fabric-contract-api-go/contractapi"
)

// 医疗记录
type MedicalRecord struct {
	Patient   string `json:"patient"`
	Doctor    string `json:"doctor"`
	Diagnosis string `json:"diagnosis"`
}

// 链码入口点
type MedicalChaincode struct {
	contractapi.Contract
}

// 添加医疗记录
func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, patient, doctor, diagnosis string) error {
	medicalRecord := MedicalRecord{
		Patient:   patient,
		Doctor:    doctor,
		Diagnosis: diagnosis,
	}

	medicalRecordJSON, err := json.Marshal(medicalRecord)
	if err != nil {
		return fmt.Errorf("转换json失败: %v", err)
	}

	err = ctx.GetStub().PutState(patient, medicalRecordJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, patient string) (*MedicalRecord, error) {
	medicalRecordJSON, err := ctx.GetStub().GetState(patient)
	if err != nil {
		return nil, fmt.Errorf("读取状态数据库失败: %v", err)
	}
	if medicalRecordJSON == nil {
		return nil, fmt.Errorf("不存在该病人的医疗记录： %s", patient)
	}

	var medicalRecord MedicalRecord
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	return &medicalRecord, nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeleteMedicalRecord(ctx contractapi.TransactionContextInterface, patient string) error {
	err := ctx.GetStub().DelState(patient)
	if err != nil {
		return fmt.Errorf("删除医疗记录失败: %v", err)
	}
	return nil
}

// 更新医疗记录
func (mc *MedicalChaincode) UpdateMedicalRecord(ctx contractapi.TransactionContextInterface, patient, newDiagnosis string) error {
	existingRecord, err := mc.GetMedicalRecord(ctx, patient)
	if err != nil {
		return fmt.Errorf("获取医疗记录失败: %v", err)
	}

	existingRecord.Diagnosis = newDiagnosis

	updatedRecordJSON, err := json.Marshal(existingRecord)
	if err != nil {
		return fmt.Errorf("转换json失败: %v", err)
	}
	err = ctx.GetStub().PutState(patient, updatedRecordJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
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
