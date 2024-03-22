package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// "hyperledgerfabric/MedicalChaincode/MedicalChaincode/main"
)



func AddMedicalRecord(byteRecord string) {
	index, medicalJSON := MarshalCSV(byteRecord)
	contract.SubmitTransaction("AddMedicalRecord", index, string(medicalJSON))
	UploadLocalFileSystem(index, string(medicalJSON))
	UploadCloudSystem(index, string(medicalJSON))
	fmt.Printf("add successfully, record: %v\n", index)
}

func GetMedicalRecord(index string) {
	var medicalRecord MedicalRecord
	medicalRecordJSON, err := contract.SubmitTransaction("GetMedicalRecord", index)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] failed to GetMedicalRecord: %w", err))
	}
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] 解析医疗记录json失败 from AddMedicalRecord: %w", err))
	}
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

// 上传本地文件存储
func UploadLocalFileSystem(index string, byteRecord string) error {
	/** 确保cloud文件夹存在 */
	dirName := "cloud"
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			panic(fmt.Errorf("创建文件夹失败: %w", err))
		}
	} else if err != nil {
		panic(fmt.Errorf("检查文件夹失败: %w", err))
	}
	/** 查看是否已有文件，有则删除之后重写一个 */
	filename := "MedicalRecord" + "_ID:" + index
	pwd, _ := os.Getwd()
	filepath := pwd + "/cloud"
	files, _ := ioutil.ReadDir(filepath)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), filename) {
			err := os.Remove(filepath + "/" + filename)
			if err != nil {
				return err
			}
			break
		}
	}
	/** 重写一个新文件 */
	ofile, _ := os.OpenFile(filepath + "/" + filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer ofile.Close()
	fmt.Fprintf(ofile, "%v\n", byteRecord)

	return nil
}


// 上传阿里云
func UploadCloudSystem(index string, byteRecord string) error {
/** 确保cloud文件夹存在 */
dirName := "cloud"
_, err := os.Stat(dirName)
if os.IsNotExist(err) {
	panic(fmt.Errorf("文件夹不存在: %w", err))
} else if err != nil {
	panic(fmt.Errorf("检查文件夹失败: %w", err))
}
filename := "MedicalRecord" + "_ID:" + index
pwd, _ := os.Getwd()
filepath := pwd + "/cloud"
/** 设置阿里云OSS访问密钥 */
accessKeyId := "LTAI5t5trSAxMSRUDwP3t9uy"  
accessKeySecret := "8GaeoaqtGPqY9Z0Epjr28br8QjVgQr"  
endpoint := "oss-cn-beijing.aliyuncs.com"  
bucketName := "sspku"  
objectName := filename // "hao" // OSS中的对象名称，通常是文件的路径和名称  
localFilename := filepath + "/" +filename 
/** 初始化OSS客户端 */ 
client, err := oss.New(endpoint, accessKeyId, accessKeySecret)  
if err != nil {  
	fmt.Println("Error:", err)  
	os.Exit(-1)  
} 
/** 获取存储空间（Bucket） */ 
bucket, err := client.Bucket(bucketName)  
if err != nil {  
	fmt.Println("Error:", err)  
	os.Exit(-1)  
}  
/** 读取本地文件 */  
file, err := os.Open(localFilename)  
if err != nil {  
	fmt.Println("Error:", err)  
	os.Exit(-1)  
}  
defer file.Close()  
/** 获取文件信息 */
// fileInfo, err := file.Stat()
_, err = file.Stat()  
if err != nil {  
	fmt.Println("Error:", err)  
	os.Exit(-1)  
}  
/** 上传文件 */
options := []oss.Option { 
	oss.ContentType("application/octet-stream"),  
}
err = bucket.PutObjectFromFile(objectName, localFilename, options...)  
if err != nil {  
	fmt.Println("Error:", err)  
	os.Exit(-1)  
}
fmt.Printf("File %s uploaded to OSS successfully.\n", localFilename) 

return nil
}