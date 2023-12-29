package main

// 导入所需的包
import (
	"time"
)

// 医疗记录
// type MedicalRecord struct {
// 	Patient   string `json:"patient"`
// 	Doctor    string `json:"doctor"`
// 	Diagnosis string `json:"diagnosis"`
// }

// 结构体各个成员的涵义、类型都有待商榷，没有仔细看过数据
// 第一列我没加进来，列名为空不知道什么涵义
type MedicalRecord struct {
	ID                 string       `json:"index 可能是记录ID"`  // 0
	SUBJECT_ID         string       `json:"患者ID"` //249
	GENDER             string       `json:"性别"`  // F
	DOB                time.Time    `json:"患者出生日期"`           // 2075-03-13 00:00:00
	DOD                string       `json:"患者死亡日期"`           // NaN
	DOD_HOSP           string       `json:"患者在医院中死亡的日期"`           // NaN
	DOD_SSN            string                  // NaN
	EXPIRE_FLAG        bool         `json:"是否过期"`  // 0
	HADM_ID            string       `json:"住院ID"`   // 患者住院期间的唯一标识； 若为门诊患者则字段为空 149546
	ADMITTIME          time.Time    `json:"患者入院的时间"`                // 2155-02-03 20:16:00
	DISCHTIME          time.Time    `json:"患者出院的时间"`                // 2155-02-14 11:15:00
	DEATHTIME          time.Time    `json:"患者死亡的具体时间`                 // NaN
	ADMISSION_TYPE     string                        // EMERGENCY
	ADMISSION_LOCATION string                        // EMERGENCY ROOM ADMIT
	DISCHARGE_LOCATION string                        // REHAB/DISTINCT PART HOSP
	INSURANCE          string                        // Medicare
	LANGUAGE           string                        // ENGL
	RELIGION           string                        // CATHOLIC
	MARITAL_STATUS     string                        // DIVORCED
	ETHNICITY          string                        // WHITE
	EDREGTIME          time.Time     `json:"在急诊科登记的时间"`                // 2155-02-03 17:43:00
	EDOUTTIME          time.Time     `json:"从急诊科出院的时间"`                // 2155-02-03 21:26:00
	DIAGNOSIS          string                        // GI BLEED/ CHEST PAIN
	HOSPITAL_EXPIRE_FLAG   bool                      // 0
	HAS_CHARTEVENTS_DATA   bool                      // 1
	ICUSTAY_Details    string   // ICUSTAY_ID: 269035, FIRST_CAREUNIT: MICU, LAST...
	ICD9_CODE          []string // 包含了这个患者这次住院所有的症状代码（和其他对应信息） //ICD-9是一套编码系统，用于标识各种疾病、疾病症状、损伤和其他健康相关情况 // ['56985', '41071', '43411', '5849', '2851', '3...
	PRO_CODE           []int    // [3974, 40, 8841, 9910]
	Drug_Details       map[string]string   //  ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
	Input_Details      map[string]string   //  ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
	Note_Details       string   //  1495461495461495461495461495461495461495461495...
}