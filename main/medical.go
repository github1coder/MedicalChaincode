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
	index                int    `json:"index"`
	SUBJECT_ID           int    `json:"SUBJECT_ID"`
	GENDER               string `json:"GENDER"`
	DOB                  string `json:"DOB"`
	DOD                  string `json:"DOD"`
	DOD_HOSP             string `json:"DOD_HOSP"`
	DOD_SSN              string `json:"DOD_SSN"`
	EXPIRE_FLAG          int    `json:"EXPIRE_FLAG"`
	HADM_ID              int    `json:"HADM_ID"`
	ADMITTIME            string `json:"ADMITTIME"`
	DISCHTIME            string `json:"DISCHTIME"`
	DEATHTIME            string `json:"DEATHTIME"`
	ADMISSION_TYPE       string `json:"ADMISSION_TYPE"`
	ADMISSION_LOCATION   string `json:"ADMISSION_LOCATION"`
	DISCHARGE_LOCATION   string `json:"DISCHARGE_LOCATION"`
	INSURANCE            string `json:"INSURANCE"`
	LANGUAGE             string `json:"LANGUAGE"`
	RELIGION             string `json:"RELIGION"`
	MARITAL_STATUS       string `json:"MARITAL_STATUS"`
	ETHNICITY            string `json:"ETHNICITY"`
	EDREGTIME            string `json:"EDREGTIME"`
	EDOUTTIME            string `json:"EDOUTTIME"`
	DIAGNOSIS            string `json:"DIAGNOSIS"`
	HOSPITAL_EXPIRE_FLAG int    `json:"HOSPITAL_EXPIRE_FLAG"`
	HAS_CHARTEVENTS_DATA int    `json:"HAS_CHARTEVENTS_DATA"`
	ICUSTAY_DETAILS      string `json:"ICUSTAY_DETAILS"`
	ICD9_CODE            string `json:"ICD9_CODE"`
	PRO_CODE             string `json:"PRO_CODE"`
	Drug_Details         string `json:"Drug_Details"`
	Input_Details        string `json:"Input_Details"`
	Note_Details         string `json:"Note_Details"`
}

// 链码入口点
type MedicalChaincode struct {
	contractapi.Contract
}

// 添加医疗记录
func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, medicalRecordJSON string) error {
	var medicalRecord MedicalRecord

	err := json.Unmarshal([]byte(medicalRecordJSON), &medicalRecord)
	if err != nil {
		return fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	err = ctx.GetStub().PutState(fmt.Sprintf("%d", medicalRecord.index), []byte(medicalRecordJSON))
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, index int) (*MedicalRecord, error) {
	medicalRecordJSON, err := ctx.GetStub().GetState(fmt.Sprintf("%d", index))
	if err != nil {
		return nil, fmt.Errorf("读取状态数据库失败: %v", err)
	}
	if medicalRecordJSON == nil {
		return nil, fmt.Errorf("不存在该医疗记录： %d", index)
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

// 主函数
func main() {
	chaincode, err := contractapi.NewChaincode(&MedicalChaincode{})
	if err != nil {
		fmt.Printf("创建医疗链码失败: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("启动医疗链码失败: %v", err)
	}
}
