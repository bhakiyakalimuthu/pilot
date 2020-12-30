package app

type Request struct {
	PatentID                  string    `json:"patentId"`
	DemographicFactor         DemoFac   `json:"demographicFactor"`
	DiagnosedWithDiabetes     bool      `json:"diagnosedWithDiabetes"`
	DiagnosisProfile          DiagProf  `json:"diagonisisProfile"`
	TreatmentVaraibles        TreatVar  `json:"treatmentVariables"`
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
	Others                    bool `json:"others"`
	None                      bool `json:"none"`
	ThyroidHarmoneLevel       bool `json:"thyroidHarmoneLevel"`
}
type TreatVar struct {
	NonPharmacologicalTherapy bool  `json:"nonPharmacologicalTherapy"`
	OralAntidiabeticDrugs     bool  `json:"oralAntidiabeticDrugs"`
	NonInsulinInjection       bool  `json:"nonInsulinInjection"`
	Insulin                   bool  `json:"insulin"`
	Other                     bool  `json:"other"`
	NoTreatment               bool  `json:"noTreatment"`
	BPLowerMedication         int64 `josn:"bplowerMedication"`
	StatinsLipidTherapy       int64 `json:"statinsLipidTherapy`
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
	PercentaRange    int64 `json:"percentageRange`
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
	AcuteCerebrovascular int64
	LimbAmputation       int64 `json:"limbAmputation"`
	LimbAmutationLevel   int64 `json:"limbAmutationLevel"`
}
type ChronicComplic struct {
}
