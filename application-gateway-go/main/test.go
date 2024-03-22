package main

import (
	"fmt"
)

var hospitalMSPID string

func test() {
	GetHospitalMSPID()
	AddMedicalRecord(input_string_0)
	GetMedicalRecord("0")
	// GetMedicalRecord("1")
	//DeleteMedicalRecord("0")
}


func GetHospitalMSPID() {
	hospital, err := contract.SubmitTransaction("GetHospitalMSPID")
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] 获取MSPID失败: %w", err))
	}
	hospitalMSPID = string(hospital)
}