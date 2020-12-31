package app

import (
	"github.com/google/uuid"
)

type User struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Age         int64  `json:"age,omitempty"`
	DOB         string `json:"dob,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Height      int64  `json:"height,omitempty"`
	Weight      int64  `json:"weight,omitempty"`
}

type UserModel struct {
	id          uuid.UUID `db:"id"`
	Name        string    `db:"name,omitempty"`
	PhoneNumber string    `db:"phone_number"`
	Address     string    `db:"address"`
	Age         int64     `db:"age"`
	DOB         string    `db:"dob"`
	Sex         string    `db:"sex"`
	Height      int64     `db:"height"`
	Weight      int64     `db:"weight"`
}

type ProfileRequest struct {
	PatentID                  string    `json:"patentId"`
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
	VisualImpairment         int64 `josn:"visualImpairment"`
	VisualAcuity             int64 `json:"visualAcuity"`
	MeasurementMehtod        int64 `json:"measurementMehtod"`
	NonproflDiabRetinopathy  bool  `json:"nonproflDiabRetinopathy"`
	ProflDiabRetinopathy     bool  `json:"proflDiabRetinopathy"`
	UnspecDiabRetinopathy    bool  `json:"unspecDiabRetinopathy"`
	MacularEdema             bool  `json:"macularEdema"`
	Other                    bool  `json:"other"`
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
