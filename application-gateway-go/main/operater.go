package main

import (
	"fmt"
	"encoding/json"
	// "hyperledgerfabric/MedicalChaincode/MedicalChaincode/main"
)

// 写入公共数据->摘要链&云
func AddMedicalRecord(byteRecord string) {
	/** 数据上云 */
	index, medicalJSON := MarshalCSVtoMR(byteRecord)
	err := UploadLocalFileSystem(index, string(medicalJSON))
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "add", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] : %w", err))
	}
	url, err := UploadCloudSystem(index, string(medicalJSON))
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "add", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] : %w", err))
	}
	/** 摘要上链 */
	_, digestJSON := MarshalCSVtoDR(byteRecord, "public", url)
	contract.SubmitTransaction("AddMedicalRecord", index, string(digestJSON))
	/** 本地记录事务 */
	SubmitTransaction(string(hospitalMSPID), "add", index, "commit")
	fmt.Printf("add successfully, record: %v\n", index)
}

// 写入私有数据->摘要链&私有数据集合
func AddPrivateMedicalRecord(byteRecord string, addr string) {
	/** 摘要上链 */
	// TODO: 如果对应index已有公共数据上链,摘要会被修改为存储private的地址,但云上数据目前不会去删除
	index, digestJSON := MarshalCSVtoDR(byteRecord, "private", addr)
	contract.SubmitTransaction("AddMedicalRecord", index, string(digestJSON))
	/** 数据上私有数据集 */
	_, medicalJSON := MarshalCSVtoMR(byteRecord)
	contract.SubmitTransaction("AddPrivateMedicalRecord", index, string(medicalJSON))
	/** 本地记录数据与事务 */
	UploadLocalFileSystem("(private "+hospitalMSPID+")"+index, string(medicalJSON))
	SubmitTransaction(string(hospitalMSPID), "add(private)", index, "commit")
	fmt.Printf("add(private) successfully, record: %v\n", index)
}

// 读取数据<-摘要链&云 or 摘要链&私有数据集
func GetMedicalRecord(index string) {
	/** 读取摘要链 */
	var digestRecord Digest
	DigestJSON, err := contract.SubmitTransaction("GetMedicalRecord", index)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "get", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w", err))
	}
	err = json.Unmarshal(DigestJSON, &digestRecord)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "get", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] 解析医疗记录json失败 from AddMedicalRecord: %w", err))
	}
	fmt.Printf("get digest: %v\n", digestRecord)
	/** 判断是否私有，并读取数据 */
	if digestRecord.PRIVATE == "private" {
		/** 读取私有数据集 */
		medicalJSON, err := contract.SubmitTransaction("GetPrivateMedicalRecord", index)
		if err != nil {
			SubmitTransaction(string(hospitalMSPID), "get(private)", index, "abort")
			panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w, orgCollection:%v", err, digestRecord.ADDRESS))
		}
		UploadLocalFileSystem("(private "+string(hospitalMSPID)+")"+index, string(medicalJSON))
		SubmitTransaction(string(hospitalMSPID), "get(private)", index, "commit")
		fmt.Printf("private record: %v loaded, orgCollection:%v\n", digestRecord.Index, digestRecord.ADDRESS)
	} else {
		/** 云拉取，更新本地记录 */
		DownloadCloudSystem(digestRecord.Index)
		SubmitTransaction(string(hospitalMSPID), "get", index, "commit")
		fmt.Printf("record: %v loaded, url:%v\n", digestRecord.Index, digestRecord.ADDRESS)
	}
}

// 删除数据 x (摘要链&云 or 摘要链&私有数据集)
func DeleteMedicalRecord(index string) {
	/** 读取摘要链 */
	var digestRecord Digest
	DigestJSON, err := contract.SubmitTransaction("GetMedicalRecord", index)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "del", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w", err))
	}
	err = json.Unmarshal(DigestJSON, &digestRecord)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "del", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] 解析医疗记录json失败: %w", err))
	}
	fmt.Printf("get digest: %v\n", digestRecord)
	/** 删除摘要 */
	_, err = contract.SubmitTransaction("DeleteMedicalRecord", index)
	if err != nil {
		SubmitTransaction(string(hospitalMSPID), "del", index, "abort")
		panic(fmt.Errorf("[BACKEND SERVICE] failed to DeleteMedicalRecord: %w", err))
	}
	/** 删除数据 */
	if digestRecord.PRIVATE == "private" {
		/** 删除私有数据集 */
		_, err := contract.SubmitTransaction("DeletePrivateMedicalRecord", index)
		if err != nil {
			SubmitTransaction(string(hospitalMSPID), "del(private)", index, "abort")
			panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w, orgCollection:%v", err, digestRecord.ADDRESS))
		}
		SubmitTransaction(string(hospitalMSPID), "del(private)", index, "commit")
	} else {
		/** 删除云数据 */
		filename := "MedicalRecord" + "_ID:" + index
		DeleteObjectFromCloudSystem(filename)
		SubmitTransaction(string(hospitalMSPID), "del", index, "commit")
	}
}


// TODO
func UpdateMedicalRecordByField(index string, field string, newValue string) {
	_, err := contract.SubmitTransaction("UpdateMedicalRecordByField", index, field, newValue)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to submit transaction: %w", err))
	}
}
