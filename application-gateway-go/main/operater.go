package main

import (
	"fmt"
)

func AddMedicalRecord(byteRecord string) {
	_, err := contract.SubmitTransaction("AddMedicalRecord", byteRecord)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to AddMedicalRecord: %w", err))
	}
}

func GetMedicalRecord(index string) {
	_, err := contract.SubmitTransaction("GetMedicalRecord", index)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w", err))
	}
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
