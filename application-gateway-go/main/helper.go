package main

import (
	"fmt"
	"encoding/json"
	"strings"
	// "time"
	"os"
	"io/ioutil"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	// "github.com/hyperledger/fabric-gateway/pkg/client"
	// "github.com/hyperledger/fabric-gateway/pkg/identity"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)

func MarshalCSV(bytes string) (string, []byte) {
	var medicalString []string
	quoteOpen := false
	var field strings.Builder

	for _, char := range bytes {
		if char == '"' {
			quoteOpen = !quoteOpen
			continue
		}

		if char != ',' || (char == ',' && quoteOpen) {
			field.WriteRune(char)
		} else {
			medicalString = append(medicalString, field.String())
			field.Reset()
		}
	}
	medicalString = append(medicalString, field.String())
	if len(medicalString) != 32 {  // 31+1
		panic(fmt.Errorf("[BACKEND SERVICE] failed to Get MarshalCSV, actual len: %v", len(medicalString)))
	}
	medicalRecord := MedicalRecord {
		Index:                medicalString[0],
		SUBJECT_ID:           medicalString[1],
		GENDER:               medicalString[2],
		DOB:                  medicalString[3],
		DOD:                  medicalString[4],
		DOD_HOSP:             medicalString[5],
		DOD_SSN:              medicalString[6],
		EXPIRE_FLAG:          medicalString[7],
		HADM_ID:              medicalString[8],
		ADMITTIME:            medicalString[9],
		DISCHTIME:            medicalString[10],
		DEATHTIME:            medicalString[11],
		ADMISSION_TYPE:       medicalString[12],
		ADMISSION_LOCATION:   medicalString[13],
		DISCHARGE_LOCATION:   medicalString[14],
		INSURANCE:            medicalString[15],
		LANGUAGE:             medicalString[16],
		RELIGION:             medicalString[17],
		MARITAL_STATUS:       medicalString[18],
		ETHNICITY:            medicalString[19],
		EDREGTIME:            medicalString[20],
		EDOUTTIME:            medicalString[21],
		DIAGNOSIS:            medicalString[22],
		HOSPITAL_EXPIRE_FLAG: medicalString[23],
		HAS_CHARTEVENTS_DATA: medicalString[24],
		ICUSTAY_DETAILS:      medicalString[25],
		ICD9_CODE:            medicalString[26],
		PRO_CODE:             medicalString[27],
		Drug_Details:         medicalString[28],
		Input_Details:        medicalString[29],
		Note_Details:         medicalString[30],
	}
	RecordJSON, err := json.Marshal(medicalRecord)
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] 转换json失败: %w", err))
	}
	return medicalRecord.Index, RecordJSON
}


func SubmitTransaction(clientMSPID string, funcType string, paraVal string, commit string) {
	hospitalMSPID, err := contract.SubmitTransaction("GetHospitalMSPID")
	if err != nil {
		panic(fmt.Errorf("[BACKEND SERVICE] 获取MSPID失败: %w", err))
	}
	currentTime := now.Format("2006-01-02 15:04:05")
	log_record := currentTime + " " + string(hospitalMSPID) + " " + funcType + " " + paraVal + " : " + commit
	filename := "LOG"
	fileContent, err := ioutil.ReadFile(filename)  
	if err != nil {  
		if os.IsNotExist(err) {  
			/** 如果文件不存在，则创建文件并写入新条目 */
			ofile, _ := os.OpenFile("./" + filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			defer ofile.Close()
			fmt.Fprintf(ofile, "%v\n", log_record)  
			return;
		}  
		panic(fmt.Errorf("[BACKEND SERVICE] LOG 文件读取失败: %w", err))
	}
	// 将文件内容按行分割  
	existingEntries := strings.Split(string(fileContent), "\n")   
	// 将新内容写回文件  
	ioutil.WriteFile(filename, []byte(""), 0644) 
	ofile, _ := os.OpenFile("./" + filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer ofile.Close()
	fmt.Fprintf(ofile, "%v", log_record)
	for i:=0; i<len(existingEntries); i++ {
		fmt.Fprintf(ofile, "\n%v", existingEntries[i])
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
fmt.Printf("File %s uploaded to OSS successfully.\n", filename) 

return nil
}
