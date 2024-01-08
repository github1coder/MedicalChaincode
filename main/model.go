package main

// 导入所需的包
import (
	// "time"
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
	index                string               `json:"index"`                // 0
	SUBJECT_ID           string               `json:"SUBJECT_ID"`           //249
	GENDER               string            		`json:"GENDER"`               // F
	DOB                  string         			`json:"DOB"`                  // 2075-03-13 00:00:00
	DOD                  string         			`json:"DOD"`                  // NaN
	DOD_HOSP             string         			`json:"DOD_HOSP"`             // NaN
	DOD_SSN              string         			`json:"DOD_SSN"`              // NaN
	EXPIRE_FLAG          string               `json:"EXPIRE_FLAG"`          // 0
	HADM_ID              string               `json:"HADM_ID"`              // 患者住院期间的唯一标识； 若为门诊患者则字段为空 149546
	ADMITTIME            string         			`json:"ADMITTIME"`            // 2155-02-03 20:16:00
	DISCHTIME            string         			`json:"DISCHTIME"`            // 2155-02-03 20:16:00
	DEATHTIME            string         			`json:"DEATHTIME"`            // NaN
	ADMISSION_TYPE       string            		`json:"ADMISSION_TYPE"`       // EMERGENCY
	ADMISSION_LOCATION   string            		`json:"ADMISSION_LOCATION"`   // EMERGENCY ROOM ADMIT
	DISCHARGE_LOCATION   string            		`json:"DISCHARGE_LOCATION"`   // REHAB/DISTINCT PART HOSP
	INSURANCE            string            		`json:"INSURANCE"`            // Medicare
	LANGUAGE             string            		`json:"LANGUAGE"`             // ENGL
	RELIGION             string            		`json:"RELIGION"`             // CATHOLIC
	MARITAL_STATUS       string            		`json:"MARITAL_STATUS"`       // DIVORCED
	ETHNICITY            string            		`json:"ETHNICITY"`            // WHITE
	EDREGTIME            string         			`json:"EDREGTIME"`            // 2155-02-03 17:43:00
	EDOUTTIME            string         			`json:"EDOUTTIME"`            // 2155-02-03 17:43:00
	DIAGNOSIS            string            		`json:"DIAGNOSIS"`            // GI BLEED/ CHEST PAIN
	HOSPITAL_EXPIRE_FLAG string               `json:"HOSPITAL_EXPIRE_FLAG"` // 0
	HAS_CHARTEVENTS_DATA string               `json:"HAS_CHARTEVENTS_DATA"` // 1
	ICUSTAY_DETAILS      string 							`json:"ICUSTAY_DETAILS"`      // ICUSTAY_ID: 269035, FIRST_CAREUNIT: MICU, LAST...
	ICD9_CODE            string          			`json:"ICD9_CODE"`            // 包含了这个患者这次住院所有的症状代码（和其他对应信息） //ICD-9是一套编码系统，用于标识各种疾病、疾病症状、损伤和其他健康相关情况 // ['56985', '41071', '43411', '5849', '2851', '3...
	PRO_CODE             string           	  `json:"PRO_CODE"`             // [3974, 40, 8841, 9910]
	Drug_Details         string 							`json:"Drug_Details"`         //  ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
	Input_Details        string 							`json:"Input_Details"`        //  ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
	Note_Details         string            		`json:"Note_Details"`         //  1495461495461495461495461495461495461495461495...
}