package main

import (
	"fmt"
	"encoding/json"
	// "hyperledgerfabric/MedicalChaincode/MedicalChaincode/main"
)



func AddMedicalRecord(byteRecord string) {
	index, medicalJSON := MarshalCSV(byteRecord)
	contract.SubmitTransaction("AddMedicalRecord", index, string(medicalJSON))
	UploadLocalFileSystem(index, string(medicalJSON))
	// UploadCloudSystem(index, string(medicalJSON))
	SubmitTransaction(string(hospitalMSPID), "add", index, "commit")
	fmt.Printf("add successfully, record: %v\n", index)
}

func GetMedicalRecord(index string) {
	var medicalRecord MedicalRecord
	medicalRecordJSON, err := contract.SubmitTransaction("GetMedicalRecord", index)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "get", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w", err))
	}
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "get", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] 解析医疗记录json失败 from AddMedicalRecord: %w", err))
	}
	SubmitTransaction(string(hospitalMSPID), "get", index, "commit")
	fmt.Printf("record: %v loaded to file MedicalRecord_ID:%v in dir ./cloud\n", medicalRecord.Index, medicalRecord.Index)
	UploadLocalFileSystem(medicalRecord.Index, string(medicalRecordJSON))
}

func DeleteMedicalRecord(index string) {
	_, err := contract.SubmitTransaction("DeleteMedicalRecord", index)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to DeleteMedicalRecord: %w", err))
	}
}

func UpdateMedicalRecordByField(index string, field string, newValue string) {
	_, err := contract.SubmitTransaction("UpdateMedicalRecordByField", index, field, newValue)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to submit transaction: %w", err))
	}
}
