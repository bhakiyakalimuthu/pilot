package app

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EmailID     string `db:"emailID,omitempty"`
	Address     string `json:"address,omitempty"`
	Age         int64  `json:"age,omitempty"`
	DOB         string `json:"dob,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Height      int64  `json:"height,omitempty"`
	Weight      int64  `json:"weight,omitempty"`
}

type UserModel struct {
	ID          uuid.UUID `db:"id"`
	ShortID     string    `db:"short_id"`
	Name        string    `db:"name"`
	PhoneNumber string    `db:"phone_number"`
	EmailID     string    `db:"email_id"`
	Address     string    `db:"address"`
	Age         int64     `db:"age"`
	DOB         string    `db:"dob"`
	Sex         string    `db:"sex"`
	Height      int64     `db:"height"`
	Weight      int64     `db:"weight"`
}

type MedicalData struct {
	PatientID                 string    `json:"patientId"`
	DemographicFactor         DemoFac   `json:"demographicFactor"`
	DiagnosedWithDiabetes     bool      `json:"diagnosedWithDiabetes"`
	DiagnosisProfile          DiagProf  `json:"diagnosisProfile"`
	TreatmentVariables        TreatVar  `json:"treatmentVariables"`
	ClinicianReportedOutcomes CliRptOut `json:"clinicReportedOutcomes"`
}
type DemoFac struct {
	Year int    `json:"year"`
	Sex  string `json:"sex"`
}
type DiagProf struct {
	DiabTypeIndic        int64  `json:"diabTypeIndic"`
	DiabeticDiagnoseYear int64  `json:"diabeticDiagnoseyear"`
	Comorbidities        Comorb `json:"comorbidities"`
}
type Comorb struct {
	LiverDisease              bool `json:"liverDisease"`
	Maligancy                 bool `json:"maligancy"`
	AIDS                      bool `json:"aids"`
	PulmonaryDisease          bool `json:"pulmonaryDisease"`
	PeripheralVascularDisease bool `json:"peripheralVascularDisease"`
	Dementia                  bool `json:"dementia"`
	Hemiplegia                bool `json:"hemiplegia"`
	Tuberculosis              bool `json:"tuberculosis"`
	HepatitisB                bool `json:"hepatitisB"`
	HepatitisC                bool `json:"hepatitisC"`
	AnxietyDisorder           bool `json:"anxietyDisorder"`
	Depression                bool `json:"Depression"`
	DisorderedEating          bool `json:"disorderedEating"`
	PsychoticMentalIllness    bool `json:"psychoticMentalIllness"`
	Thyroid                   bool `json:"thyroid"`
	ComorbOthers              bool `json:"comorbOthers"`
	None                      bool `json:"none"`
	ThyroidHarmoneLevel       bool `json:"thyroidHarmoneLevel"`
}
type TreatVar struct {
	NonPharmacologicalTherapy bool  `json:"nonPharmacologicalTherapy"`
	OralAntidiabeticDrugs     bool  `json:"oralAntidiabeticDrugs"`
	NonInsulinInjection       bool  `json:"nonInsulinInjection"`
	Insulin                   bool  `json:"insulin"`
	TreatVarOther             bool  `json:"treatVarOther"`
	NoTreatment               bool  `json:"noTreatment"`
	BPLowerMedication         int64 `josn:"bplowerMedication"`
	StatinsLipidTherapy       int64 `json:"statinsLipidTherapy"`
}
type CliRptOut struct {
	GlycemicControl      GlycContrl     `json:"glycemicControl"`
	IntermediateOutcomes IntmedOutcomes `json:"intermediateOutcomes"`
	AcuteEvents          AcutEvnt       `json:"acuteEvents"`
	ChronicComplications ChronicComplic `json:"chronicComplictions"`
}
type GlycContrl struct {
	HbA1cReading     int64 `json:"hba1cReading"`
	Hb1AcReadingUnit int64 `json:"hb1acReadingUnit"`
	PercentaRange    int64 `json:"percentageRange"`
}
type IntmedOutcomes struct {
	SystolicBP     int64 `json:"systolicBP"`
	DiastolicBP    int64 `json:"diastolicBP"`
	Cholesterol    int64 `json:"cholesterol"`
	LDLCholesterol int64 `json:"ldlCholesterol"`
	HDLCholesterol int64 `json:"hdlCholesterol"`
	Triglycerids   int64 `json:"triglicerids"`
	LDLCholUnit    int64 `json:"ldlCholUnit"`
	HDLCholUnit    int64 `json:"hdlCholUnit"`
	TriglycUnit    int64 `json:"triglycUnit"`
	Weight         int64 `json:"weight"`
	Height         int64 `json:"height"`
	Waist          int64 `json:"waist"`
}
type AcutEvnt struct {
	DKAHHS               int64 `json:"dkahhs"`
	Hypoglycemia         int64 `json:"hypoglycemia"`
	HypoglycemiaL2Exp    int64 `json:"hypoglycemiaL2Exp"`
	HypglycemWOSymp      int64 `json:"hypglycemWOSymp"`
	AcuteIschemic        int64 `json:"acuteIschemic"`
	AcuteCerebrovascular int64 `json:"acuteCerebrovascular"`
	LimbAmputation       int64 `json:"limbAmputation"`
	LimbAmputationLevel  int64 `json:"limbAmputationLevel"`
}
type ChronicComplic struct {
	VisualImpairment         int64 `josn:"visualImpairment"`
	VisualAcuity             int64 `json:"visualAcuity"`
	MeasurementMehtod        int64 `json:"measurementMehtod"`
	NonproflDiabRetinopathy  bool  `json:"nonproflDiabRetinopathy"`
	ProflDiabRetinopathy     bool  `json:"proflDiabRetinopathy"`
	UnspecDiabRetinopathy    bool  `json:"unspecDiabRetinopathy"`
	MacularEdema             bool  `json:"macularEdema"`
	ChronicComplicOther      bool  `json:"chronicComplicOther"`
	NosightThreaten          bool  `json:"nosightThreaten"`
	UnknownsightThreaten     bool  `json:"unknownSightThreaten"`
	AutonomicNeuropathy      int64 `json:"autonomicNeuropathy"`
	PeripheralNeuropathy     int64 `json:"peripheralNeuropathy"`
	PeripherNeuropSympt      int64 `json:"peripherNeuropSympt"`
	CharcotFoot              int64 `json:"charcotFoot"`
	LowerLimbUlcer           int64 `json:"lowerLimbUlcer"`
	LLUSeverity              int64 `json:"lluSeverity"`
	SevereStageDiagnosed     int64 `json:"severeStageDiagnosed"`
	PeripheralArteryDisease  int64 `json:"peripheralArteryDisease"`
	PeripheralArterySymptoms int64 `json:"peripheralArterySymptoms"`
	IschemicHeartDisease     int64 `json:"ischemicHeartDisease"`
	HeartFailure             int64 `josn:"heartFailure"`
	HeartFailureStage        int64 `json:"heartFailureStage"`
	EGFR                     int64 `json:"egfr"`
	UrinaryACR               int64 `json:"urinaryACR"`
	Dialysis                 int64 `json:"dialysis"`
	CerebrovascularDisease   int64 `json:"cerebrovascularDisease"`
	PeriodontalHealth        int64 `json:"periodontalHealth"`
	ErectileDysfunction      int64 `json:"erectileDysfunction"`
	Lipodystrophy            int64 `json:"lipodystrophy"`
}

type MedicalDataModel struct {
	ShortID   string    `db:"short_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Year      int       `db:"year"`
	Sex       string    `db:"sex"`
	// DiagnosedWithDiabetes     bool      `db:"diagnosed_with_diabetes"`
	// DiabTypeIndic        int64  `json:"diab_type_indic"`
	// DiabeticDiagnoseYear int64  `json:"diabetic_diagnoseyear"`
	// Comorbidities        Comorb `json:"comorbidities"`
	// DiagnosisProfile          DiagProf `json:"diagnosisProfile"`
	// TreatmentVariables        TreatVar  `json:"treatmentVariables"`
	// ClinicianReportedOutcomes CliRptOut `json:"clinicReportedOutcomes"`
}

type DiagnosisProfileModel struct {
	ShortID                   string    `db:"short_id"`
	CreatedAt                 time.Time `db:"created_at"`
	UpdatedAt                 time.Time `db:"updated_at"`
	Year                      int       `db:"year"`
	Sex                       string    `db:"sex"`
	DiagnosedWithDiabetes     bool      `db:"diagnosed_with_diabetes"`
	DiabTypeIndic             int64     `db:"diab_type_indic"`
	DiabeticDiagnoseYear      int64     `db:"diabetic_diagnoseyear"`
	LiverDisease              bool      `db:"comorb_liver_disease"`
	Maligancy                 bool      `db:"comorb_maligancy"`
	AIDS                      bool      `db:"comorb_aids"`
	PulmonaryDisease          bool      `db:"comorb_pulmonary_disease"`
	PeripheralVascularDisease bool      `db:"comorb_peripheral_vascular_disease"`
	Dementia                  bool      `db:"comorb_dementia"`
	Hemiplegia                bool      `db:"comorb_hemiplegia"`
	Tuberculosis              bool      `db:"comorb_tuberculosis"`
	HepatitisB                bool      `db:"comorb_hepatitisB"`
	HepatitisC                bool      `db:"comorb_hepatitisC"`
	AnxietyDisorder           bool      `db:"comorb_anxietyDisorder"`
	Depression                bool      `db:"comorb_depression"`
	DisorderedEating          bool      `db:"comorb_disordered_eating"`
	PsychoticMentalIllness    bool      `db:"comorb_psychotic_mental_illness"`
	Thyroid                   bool      `db:"comorb_thyroid"`
	ComorbOthers              bool      `db:"comorb_others"`
	None                      bool      `db:"comorb_none"`
	ThyroidHarmoneLevel       bool      `db:"comorb_thyroid_harmone_level"`
	NonPharmacologicalTherapy bool      `db:"treatvar_non_pharmacological_therapy"`
	OralAntidiabeticDrugs     bool      `db:"treatvar_oral_antidiabetic_drugs"`
	NonInsulinInjection       bool      `db:"treatvar_non_insulin_injection"`
	Insulin                   bool      `db:"treatvar_insulin"`
	TreatVarOther             bool      `db:"treatvar_other"`
	NoTreatment               bool      `db:"treatvar_no_treatment"`
	BPLowerMedication         int64     `db:"treatvar_bp_lower_medication"`
	StatinsLipidTherapy       int64     `db:"treatvar_statins_lipid_therapy"`
	HbA1cReading              int64     `db:"clirptdoutc_glyc_hba1c_reading"`
	Hb1AcReadingUnit          int64     `db:"clirptdoutc_glyc_hb1ac_reading_unit"`
	PercentaRange             int64     `db:"clirptdoutc_glyc_percentage_range"`
	SystolicBP                int64     `db:"clirptdoutc_intmoutc_systolic_bp"`
	DiastolicBP               int64     `db:"clirptdoutc_intmoutc_diastolic_bp"`
	Cholesterol               int64     `db:"clirptdoutc_intmoutc_cholesterol"`
	LDLCholesterol            int64     `db:"clirptdoutc_intmoutc_ldlCholesterol"`
	LDLCholUnit               int64     `db:"clirptdoutc_intmoutc_ldlCholUnit"`
	HDLCholesterol            int64     `db:"clirptdoutc_intmoutc_hdlCholesterol"`
	HDLCholUnit               int64     `db:"clirptdoutc_intmoutc_hdlCholUnit"`
	Triglycerids              int64     `db:"clirptdoutc_intmoutc_triglicerids"`
	TriglycUnit               int64     `db:"clirptdoutc_intmoutc_triglycUnit"`
	Weight                    int64     `db:"clirptdoutc_intmoutc_weight"`
	Height                    int64     `db:"clirptdoutc_intmoutc_height"`
	Waist                     int64     `db:"clirptdoutc_intmoutc_waist"`
	DKAHHS                    int64     `db:"clirptdoutc_acutevnt_dkahhs"`
	Hypoglycemia              int64     `db:"clirptdoutc_acutevnt_hypoglycemia"`
	HypoglycemiaL2Exp         int64     `db:"clirptdoutc_acutevnt_hypoglycemia_l2exp"`
	HypglycemWOSymp           int64     `db:"clirptdoutc_acutevnt_hypglycem_wosymp"`
	AcuteIschemic             int64     `db:"clirptdoutc_acutevnt_acute_ischemic"`
	AcuteCerebrovascular      int64     `db:"clirptdoutc_acutevnt_acute_cerebro_vascular"`
	LimbAmputation            int64     `db:"clirptdoutc_acutevnt_limb_amputation"`
	LimbAmputationLevel       int64     `db:"clirptdoutc_acutevnt_limb_amputation_level"`
	VisualImpairment          int64     `db:"clirptdoutc_chroniccomplic_visual_impairment"`
	VisualAcuity              int64     `db:"clirptdoutc_chroniccomplic_visual_acuity"`
	MeasurementMehtod         int64     `db:"clirptdoutc_chroniccomplic_measurement_mehtod"`
	NonproflDiabRetinopathy   bool      `db:"clirptdoutc_chroniccomplic_nonprofl_diab_retinopathy"`
	ProflDiabRetinopathy      bool      `db:"clirptdoutc_chroniccomplic_profl_diab_retinopathy"`
	UnspecDiabRetinopathy     bool      `db:"clirptdoutc_chroniccomplic_unspec_diab_retinopathy"`
	MacularEdema              bool      `db:"clirptdoutc_chroniccomplic_macular_edema"`
	ChronicComplicOther       bool      `db:"clirptdoutc_chroniccomplic_other"`
	NosightThreaten           bool      `db:"clirptdoutc_chroniccomplic_nosight_threaten"`
	UnknownsightThreaten      bool      `db:"clirptdoutc_chroniccomplic_unknown_sight_threaten"`
	AutonomicNeuropathy       int64     `db:"clirptdoutc_chroniccomplic_autonomic_neuropathy"`
	PeripheralNeuropathy      int64     `db:"clirptdoutc_chroniccomplic_peripheral_neuropathy"`
	PeripherNeuropSympt       int64     `db:"clirptdoutc_chroniccomplic_peripher_neurop_sympt"`
	CharcotFoot               int64     `db:"clirptdoutc_chroniccomplic_charcot_foot"`
	LowerLimbUlcer            int64     `db:"clirptdoutc_chroniccomplic_lower_limb_ulcer"`
	LLUSeverity               int64     `db:"clirptdoutc_chroniccomplic_llu_severity"`
	SevereStageDiagnosed      int64     `db:"clirptdoutc_chroniccomplic_severe_stage_diagnosed"`
	PeripheralArteryDisease   int64     `db:"clirptdoutc_chroniccomplic_peripheral_artery_disease"`
	PeripheralArterySymptoms  int64     `db:"clirptdoutc_chroniccomplic_peripheral_artery_symptoms"`
	IschemicHeartDisease      int64     `db:"clirptdoutc_chroniccomplic_ischemic_heart_disease"`
	HeartFailure              int64     `db:"clirptdoutc_chroniccomplic_heart_failure"`
	HeartFailureStage         int64     `db:"clirptdoutc_chroniccomplic_heart_failure_stage"`
	EGFR                      int64     `db:"clirptdoutc_chroniccomplic_egfr"`
	UrinaryACR                int64     `db:"clirptdoutc_chroniccomplic_urinary_ACR"`
	Dialysis                  int64     `db:"clirptdoutc_chroniccomplic_dialysis"`
	CerebrovascularDisease    int64     `db:"clirptdoutc_chroniccomplic_cerebrovascular_disease"`
	PeriodontalHealth         int64     `db:"clirptdoutc_chroniccomplic_periodontal_health"`
	ErectileDysfunction       int64     `db:"clirptdoutc_chroniccomplic_erectile_dysfunction"`
	Lipodystrophy             int64     `db:"clirptdoutc_chroniccomplic_lipodystrophy"`
}
