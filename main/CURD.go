package main

// 导入所需的包
import (
	"encoding/json"
	"reflect"

	// "encoding/csv"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strings"
)

// 添加医疗记录
// 数据集里好像没有doctor，我先改成index了
// func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, patient, index, diagnosis string) error {
// 	medicalRecord := MedicalRecord{
// 		ID:            index,
// 		SUBJECT_ID:    patient,
// 		Note_Details:  diagnosis,
// 		ICD9_CODE:     []string{},
// 		PRO_CODE:      []int{},
// 		Drug_Details:  make(map[string]string),
// 		Input_Details: make(map[string]string),
// 	}

// 	medicalRecordJSON, err := json.Marshal(medicalRecord)
// 	if err != nil {
// 		return fmt.Errorf("转换json失败: %v", err)
// 	}

// 	err = ctx.GetStub().PutState(patient, medicalRecordJSON)
// 	if err != nil {
// 		return fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return nil
// }

// // 添加医疗记录
// func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, medicalRecordJSON string) error {
// 	var medicalRecord MedicalRecord

// 	err := json.Unmarshal([]byte(medicalRecordJSON), &medicalRecord)
// 	if err != nil {
// 		return fmt.Errorf("解析医疗记录json失败: %v", err)
// 	}

// 	err = ctx.GetStub().PutState(medicalRecord.ID, []byte(medicalRecordJSON))
// 	if err != nil {
// 		return fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return nil
// }

// 对ICUSTAY_Details之后的信息目前还是单独作为参数，目前没找到比较好的办法，peer invoke的json格式不允许""里面套复杂的字符
// func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, byteRecord string,
// 	ICUSTAY_Details string, ICD9_CODE []string, PRO_CODE []int, Drug_Details string, Input_Details string, Note_Details string) ([]string,error) {

// 	medicalStream := strings.Split(byteRecord, ",")
// 	medicalRecord := MedicalRecord{
// 		ID:
// 		SUBJECT_ID:         string       `json:"患者ID"`  // 249
// 		GENDER:             string       `json:"性别"`    // F
// 		DOB:                time.Time    `json:"患者出生日期"`           // 2075-03-13 00:00:00
// 		DOD:                string       `json:"患者死亡日期"`           // NaN
// 		DOD_HOSP:           string       `json:"患者在医院中死亡的日期"`           // NaN
// 		DOD_SSN:            string                  // NaN
// 		EXPIRE_FLAG:        bool         `json:"是否过期"`  // 0
// 		HADM_ID:            string       `json:"住院ID"`   // 患者住院期间的唯一标识； 若为门诊患者则字段为空 149546
// 		ADMITTIME:          time.Time    `json:"患者入院的时间"`                // 2155-02-03 20:16:00
// 		DISCHTIME:          time.Time    `json:"患者出院的时间"`                // 2155-02-14 11:15:00
// 		DEATHTIME:          time.Time    `json:"患者死亡的具体时间`                 // NaN
// 		ADMISSION_TYPE:     string                        // EMERGENCY
// 		ADMISSION_LOCATION: string                        // EMERGENCY ROOM ADMIT
// 		DISCHARGE_LOCATION: string                        // REHAB/DISTINCT PART HOSP
// 		INSURANCE:          string                        // Medicare
// 		LANGUAGE:           string                        // ENGL
// 		RELIGION:           string                        // CATHOLIC
// 		MARITAL_STATUS:     string                        // DIVORCED
// 		ETHNICITY:          string                        // WHITE
// 		EDREGTIME:          time.Time     `json:"在急诊科登记的时间"`                // 2155-02-03 17:43:00
// 		EDOUTTIME:          time.Time     `json:"从急诊科出院的时间"`                // 2155-02-03 21:26:00
// 		DIAGNOSIS:          string                        // GI BLEED/ CHEST PAIN
// 		HOSPITAL_EXPIRE_FLAG:   bool                      // 0
// 		HAS_CHARTEVENTS_DATA   bool                      // 1

// 		ICUSTAY_Details    string   // ICUSTAY_ID: 269035, FIRST_CAREUNIT: MICU, LAST...
// 		ICD9_CODE          []string // 包含了这个患者这次住院所有的症状代码（和其他对应信息） //ICD-9是一套编码系统，用于标识各种疾病、疾病症状、损伤和其他健康相关情况 // ['56985', '41071', '43411', '5849', '2851', '3...
// 		PRO_CODE           []int    // [3974, 40, 8841, 9910]
// 		// Drug_Details       map[string]string   //  ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
// 		// Input_Details      map[string]string   //  ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
// 		Drug_Details       string   //  ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
// 		Input_Details      string   //  ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
// 		Note_Details       string
// 	}

// 	return medicalStream ,nil
// }

func (mc *MedicalChaincode) AddMedicalRecord(ctx contractapi.TransactionContextInterface, byteRecord string) (*MedicalRecord, error) {
	var medStrings []string
	quoteOpen := false
	var str strings.Builder

	for _, char := range byteRecord {
		if char == '[' || char == ']' {
			quoteOpen = !quoteOpen
			continue
		}

		if char == ',' && !quoteOpen {
			medStrings = append(medStrings, str.String())
			str.Reset()
		} else {
			str.WriteRune(char)
		}
	}
	medStrings = append(medStrings, str.String())
	// medicalRecord := MedicalRecord {
	// 		// index: 								medStrings[0],
	// 		// SUBJECT_ID: 					medStrings[1],
	// 	}

	medicalRecord := MedicalRecord{
		index:                medStrings[0],
		SUBJECT_ID:           medStrings[1],
		GENDER:               medStrings[2],
		DOB:                  medStrings[3],
		DOD:                  medStrings[4],
		DOD_HOSP:             medStrings[5],
		DOD_SSN:              medStrings[6],
		EXPIRE_FLAG:          medStrings[7],
		HADM_ID:              medStrings[8],
		ADMITTIME:            medStrings[9],
		DISCHTIME:            medStrings[10],
		DEATHTIME:            medStrings[11],
		ADMISSION_TYPE:       medStrings[12],
		ADMISSION_LOCATION:   medStrings[13],
		DISCHARGE_LOCATION:   medStrings[14],
		INSURANCE:            medStrings[15],
		LANGUAGE:             medStrings[16],
		RELIGION:             medStrings[17],
		MARITAL_STATUS:       medStrings[18],
		ETHNICITY:            medStrings[19],
		EDREGTIME:            medStrings[20],
		EDOUTTIME:            medStrings[21],
		DIAGNOSIS:            medStrings[22],
		HOSPITAL_EXPIRE_FLAG: medStrings[23],
		HAS_CHARTEVENTS_DATA: medStrings[24],
		ICUSTAY_DETAILS:      medStrings[25],
		ICD9_CODE:            medStrings[26],
		PRO_CODE:             medStrings[27],
		Drug_Details:         medStrings[28],
		Input_Details:        medStrings[29],
		Note_Details:         medStrings[30],
	}

	RecordJSON, err := json.Marshal(medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("转换json失败: %v", err)
	}
	err = ctx.GetStub().PutState(medicalRecord.index, RecordJSON)
	if err != nil {
		return nil, fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return &medicalRecord, nil
}

// 获取医疗记录
func (mc *MedicalChaincode) GetMedicalRecord(ctx contractapi.TransactionContextInterface, index string) (*MedicalRecord, error) {
	medicalRecordJSON, err := ctx.GetStub().GetState(index)
	if err != nil {
		return nil, fmt.Errorf("读取状态数据库失败: %v", err)
	}
	// TODO: 在不存在该键值记录时GetStub().GetState()会直接返回下面报错，不清楚为什么，不过目前影响不大
	// Error: endorsement failure during query. response: status:500 message:"\344\270\215 ......
	if medicalRecordJSON == nil {
		return nil, fmt.Errorf("不存在该病人的医疗记录： %s", index)
	}

	var medicalRecord MedicalRecord
	err = json.Unmarshal(medicalRecordJSON, &medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("解析医疗记录json失败: %v", err)
	}

	return &medicalRecord, nil
}

// 删除医疗记录
func (mc *MedicalChaincode) DeleteMedicalRecord(ctx contractapi.TransactionContextInterface, index string) error {
	err := ctx.GetStub().DelState(index)
	if err != nil {
		return fmt.Errorf("删除医疗记录失败: %v", err)
	}
	return nil
}

// 更新医疗记录，根据传入的index、字段名和新值修改medicalRecord
func (mc *MedicalChaincode) UpdateMedicalRecordByField(ctx contractapi.TransactionContextInterface, index string, field string, newValue string) error {
	// 获取当前的医疗记录
	medicalRecord, err := mc.GetMedicalRecord(ctx, index)
	if err != nil {
		return fmt.Errorf("获取医疗记录失败: %v", err)
	}

	// 使用反射来更新指定字段的值
	rv := reflect.ValueOf(&medicalRecord).Elem()
	rv.FieldByName(field).SetString(newValue)

	// 转换回 JSON
	updatedRecordJSON, err := json.Marshal(medicalRecord)
	if err != nil {
		return fmt.Errorf("转换json失败: %v", err)
	}

	// 存回链码状态
	err = ctx.GetStub().PutState(index, updatedRecordJSON)
	if err != nil {
		return fmt.Errorf("存入状态数据库失败: %v", err)
	}

	return nil
}

// // 更新医疗记录
// func (mc *MedicalChaincode) UpdateMedicalRecord(ctx contractapi.TransactionContextInterface, patient, newDiagnosis string) error {
// 	existingRecord, err := mc.GetMedicalRecord(ctx, patient)
// 	if err != nil {
// 		return fmt.Errorf("获取医疗记录失败: %v", err)
// 	}

// 	existingRecord.Note_Details = newDiagnosis

// 	updatedRecordJSON, err := json.Marshal(existingRecord)
// 	if err != nil {
// 		return fmt.Errorf("转换json失败: %v", err)
// 	}
// 	err = ctx.GetStub().PutState(patient, updatedRecordJSON)
// 	if err != nil {
// 		return fmt.Errorf("存入状态数据库失败: %v", err)
// 	}

// 	return nil
// }
