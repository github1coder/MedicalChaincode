package main

// 导入所需的包
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


// 添加医疗记录
// 数据集里好像没有doctor，我先改成index了
func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, patient, index, diagnosis string) error {
	medicalRecord := MedicalRecord{
		ID:            index,
		SUBJECT_ID:    patient,
		Note_Details:  diagnosis,
		ICD9_CODE:     []string{},
		PRO_CODE:      []int{},
		Drug_Details:  make(map[string]string),
		Input_Details: make(map[string]string),
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
	// TODO: 在不存在该键值记录时GetStub().GetState()会直接返回下面报错，不清楚为什么，不过目前影响不大
	// Error: endorsement failure during query. response: status:500 message:"\344\270\215 ......
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

	existingRecord.Note_Details = newDiagnosis

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