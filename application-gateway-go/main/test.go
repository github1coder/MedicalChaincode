package main

import (
	"fmt"
)

var hospitalMSPID string
// var accessKeyId = "" 
// var accessKeySecret = ""
var accessKeyId = "LTAI5t5trSAxMSRUDwP3t9uy"  
var accessKeySecret = "8GaeoaqtGPqY9Z0Epjr28br8QjVgQr"  
var endpoint = "oss-cn-beijing.aliyuncs.com"  
var bucketName = "sspku"

func test() {
	GetHospitalMSPID()
	
	AddMedicalRecord(input_string_0)
	GetMedicalRecord("0")
	AddPrivateMedicalRecord(input_string_0, hospitalMSPID+"PrivateMedicalCollection")
	GetMedicalRecord("0")
	AddMedicalRecord(input_string_1)
	GetMedicalRecord("1")
	DeleteMedicalRecord("0")

}


func GetHospitalMSPID() {
	hospital, err := contract.SubmitTransaction("GetHospitalMSPID")
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] 获取MSPID失败: %w", err))
	}
	hospitalMSPID = string(hospital)
}