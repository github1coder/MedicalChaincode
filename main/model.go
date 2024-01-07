package main

import "time"

// 医疗记录
type MedicalRecord struct {
	index                int               `json:"index"`                // 0
	SUBJECT_ID           int               `json:"SUBJECT_ID"`           //249
	GENDER               string            `json:"GENDER"`               // F
	DOB                  time.Time         `json:"DOB"`                  // 2075-03-13 00:00:00
	DOD                  time.Time         `json:"DOD"`                  // NaN
	DOD_HOSP             time.Time         `json:"DOD_HOSP"`             // NaN
	DOD_SSN              time.Time         `json:"DOD_SSN"`              // NaN
	EXPIRE_FLAG          int               `json:"EXPIRE_FLAG"`          // 0
	HADM_ID              int               `json:"HADM_ID"`              // 患者住院期间的唯一标识； 若为门诊患者则字段为空 149546
	ADMITTIME            time.Time         `json:"ADMITTIME"`            // 2155-02-03 20:16:00
	DISCHTIME            time.Time         `json:"DISCHTIME"`            // 2155-02-03 20:16:00
	DEATHTIME            time.Time         `json:"DEATHTIME"`            // NaN
	ADMISSION_TYPE       string            `json:"ADMISSION_TYPE"`       // EMERGENCY
	ADMISSION_LOCATION   string            `json:"ADMISSION_LOCATION"`   // EMERGENCY ROOM ADMIT
	DISCHARGE_LOCATION   string            `json:"DISCHARGE_LOCATION"`   // REHAB/DISTINCT PART HOSP
	INSURANCE            string            `json:"INSURANCE"`            // Medicare
	LANGUAGE             string            `json:"LANGUAGE"`             // ENGL
	RELIGION             string            `json:"RELIGION"`             // CATHOLIC
	MARITAL_STATUS       string            `json:"MARITAL_STATUS"`       // DIVORCED
	ETHNICITY            string            `json:"ETHNICITY"`            // WHITE
	EDREGTIME            time.Time         `json:"EDREGTIME"`            // 2155-02-03 17:43:00
	EDOUTTIME            time.Time         `json:"EDOUTTIME"`            // 2155-02-03 17:43:00
	DIAGNOSIS            string            `json:"DIAGNOSIS"`            // GI BLEED/ CHEST PAIN
	HOSPITAL_EXPIRE_FLAG int               `json:"HOSPITAL_EXPIRE_FLAG"` // 0
	HAS_CHARTEVENTS_DATA int               `json:"HAS_CHARTEVENTS_DATA"` // 1
	ICUSTAY_DETAILS      map[string]string `json:"ICUSTAY_DETAILS"`      // ICUSTAY_ID: 269035, FIRST_CAREUNIT: MICU, LAST...
	ICD9_CODE            []string          `json:"ICD9_CODE"`            // 包含了这个患者这次住院所有的症状代码（和其他对应信息） //ICD-9是一套编码系统，用于标识各种疾病、疾病症状、损伤和其他健康相关情况 // ['56985', '41071', '43411', '5849', '2851', '3...
	PRO_CODE             []int             `json:"PRO_CODE"`             // [3974, 40, 8841, 9910]
	Drug_Details         map[string]string `json:"Drug_Details"`         //  ICUSTAY_ID: 208413, DRUG: 0.9% Sodium Chloride...
	Input_Details        map[string]string `json:"Input_Details"`        //  ICUSTAY_ID: 269035.0, STARTTIME: 2155-02-05 17...
	Note_Details         string            `json:"Note_Details"`         //  1495461495461495461495461495461495461495461495...
}
