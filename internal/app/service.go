package app

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	store  Store
}

func NewService(logger *zap.Logger, store Store) *Service {
	return &Service{
		logger: logger,
		store:  store,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *User) error {
	// s.logger.Info("user input",zap.Any("user",user))
	// fmt.Println(user)

	id := uuid.New()
	fmt.Println(id)
	shortID := shortuuid.New()
	fmt.Println(shortID)
	s.logger.Info("user id's", zap.String("id", id.String()), zap.String("shortID", shortID))
	model := &UserModel{
		ID:          id,
		ShortID:     shortID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		EmailID:     user.EmailID,
		Address:     user.Address,
		Age:         user.Age,
		DOB:         user.DOB,
		Sex:         user.Sex,
		Height:      user.Height,
		Weight:      user.Weight,
	}
	if err := s.store.CreateUser(ctx, model); err != nil {
		s.logger.Error("failed to create profile", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, userID string) (*UserModel, error) {
	return s.store.GetUser(ctx, userID)
}

func (s *Service) CreateMedicalData(ctx context.Context, data *MedicalData) error {
	m := MedicalDataModel{
		ShortID:   data.PatientID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Year:      data.DemographicFactor.Year,
		Sex:       data.DemographicFactor.Sex,
	}
	if err := s.store.CreateMedicalData(ctx, &m); err != nil {
		s.logger.Error("failed to create medical data", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) CreateDiagnosisProfile(ctx context.Context, data *MedicalData) error {
	m := DiagnosisProfileModel{
		ShortID:                   data.PatientID,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
		Year:                      data.DemographicFactor.Year,
		Sex:                       data.DemographicFactor.Sex,
		DiabTypeIndic:             data.DiagnosisProfile.DiabTypeIndic,
		DiabeticDiagnoseYear:      data.DiagnosisProfile.DiabeticDiagnoseYear,
		LiverDisease:              data.DiagnosisProfile.Comorbidities.LiverDisease,
		Maligancy:                 data.DiagnosisProfile.Comorbidities.Maligancy,
		AIDS:                      data.DiagnosisProfile.Comorbidities.AIDS,
		PulmonaryDisease:          data.DiagnosisProfile.Comorbidities.PulmonaryDisease,
		PeripheralVascularDisease: data.DiagnosisProfile.Comorbidities.PeripheralVascularDisease,
		Dementia:                  data.DiagnosisProfile.Comorbidities.Dementia,
		Hemiplegia:                data.DiagnosisProfile.Comorbidities.Hemiplegia,
		Tuberculosis:              data.DiagnosisProfile.Comorbidities.Tuberculosis,
		HepatitisB:                data.DiagnosisProfile.Comorbidities.HepatitisB,
		HepatitisC:                data.DiagnosisProfile.Comorbidities.HepatitisC,
		AnxietyDisorder:           data.DiagnosisProfile.Comorbidities.AnxietyDisorder,
		Depression:                data.DiagnosisProfile.Comorbidities.Depression,
		DisorderedEating:          data.DiagnosisProfile.Comorbidities.DisorderedEating,
		PsychoticMentalIllness:    data.DiagnosisProfile.Comorbidities.PsychoticMentalIllness,
		Thyroid:                   data.DiagnosisProfile.Comorbidities.Thyroid,
		ComorbOthers:              data.DiagnosisProfile.Comorbidities.ComorbOthers,
		None:                      data.DiagnosisProfile.Comorbidities.None,
		ThyroidHarmoneLevel:       data.DiagnosisProfile.Comorbidities.ThyroidHarmoneLevel,
		NonPharmacologicalTherapy: data.TreatmentVariables.NonPharmacologicalTherapy,
		OralAntidiabeticDrugs:     data.TreatmentVariables.OralAntidiabeticDrugs,
		NonInsulinInjection:       data.TreatmentVariables.NonInsulinInjection,
		Insulin:                   data.TreatmentVariables.Insulin,
		TreatVarOther:             data.TreatmentVariables.TreatVarOther,
		NoTreatment:               data.TreatmentVariables.NoTreatment,
		BPLowerMedication:         data.TreatmentVariables.BPLowerMedication,
		StatinsLipidTherapy:       data.TreatmentVariables.StatinsLipidTherapy,
		HbA1cReading:              data.ClinicianReportedOutcomes.GlycemicControl.HbA1cReading,
		Hb1AcReadingUnit:          data.ClinicianReportedOutcomes.GlycemicControl.Hb1AcReadingUnit,
		PercentaRange:             data.ClinicianReportedOutcomes.GlycemicControl.PercentaRange,
		SystolicBP:                data.ClinicianReportedOutcomes.IntermediateOutcomes.SystolicBP,
		DiastolicBP:               data.ClinicianReportedOutcomes.IntermediateOutcomes.DiastolicBP,
		Cholesterol:               data.ClinicianReportedOutcomes.IntermediateOutcomes.Cholesterol,
		LDLCholesterol:            data.ClinicianReportedOutcomes.IntermediateOutcomes.LDLCholesterol,
		HDLCholesterol:            data.ClinicianReportedOutcomes.IntermediateOutcomes.HDLCholesterol,
		Triglycerids:              data.ClinicianReportedOutcomes.IntermediateOutcomes.Triglycerids,
		LDLCholUnit:               data.ClinicianReportedOutcomes.IntermediateOutcomes.LDLCholUnit,
		HDLCholUnit:               data.ClinicianReportedOutcomes.IntermediateOutcomes.HDLCholUnit,
		TriglycUnit:               data.ClinicianReportedOutcomes.IntermediateOutcomes.TriglycUnit,
		Weight:                    data.ClinicianReportedOutcomes.IntermediateOutcomes.Weight,
		Height:                    data.ClinicianReportedOutcomes.IntermediateOutcomes.Height,
		Waist:                     data.ClinicianReportedOutcomes.IntermediateOutcomes.Waist,
		DKAHHS:                    data.ClinicianReportedOutcomes.AcuteEvents.DKAHHS,
		Hypoglycemia:              data.ClinicianReportedOutcomes.AcuteEvents.Hypoglycemia,
		HypoglycemiaL2Exp:         data.ClinicianReportedOutcomes.AcuteEvents.HypoglycemiaL2Exp,
		HypglycemWOSymp:           data.ClinicianReportedOutcomes.AcuteEvents.HypglycemWOSymp,
		AcuteIschemic:             data.ClinicianReportedOutcomes.AcuteEvents.AcuteIschemic,
		AcuteCerebrovascular:      data.ClinicianReportedOutcomes.AcuteEvents.AcuteCerebrovascular,
		LimbAmputation:            data.ClinicianReportedOutcomes.AcuteEvents.LimbAmputation,
		LimbAmputationLevel:       data.ClinicianReportedOutcomes.AcuteEvents.LimbAmputationLevel,
		VisualImpairment:          data.ClinicianReportedOutcomes.ChronicComplications.VisualImpairment,
		VisualAcuity:              data.ClinicianReportedOutcomes.ChronicComplications.VisualAcuity,
		MeasurementMehtod:         data.ClinicianReportedOutcomes.ChronicComplications.MeasurementMehtod,
		NonproflDiabRetinopathy:   data.ClinicianReportedOutcomes.ChronicComplications.NonproflDiabRetinopathy,
		ProflDiabRetinopathy:      data.ClinicianReportedOutcomes.ChronicComplications.ProflDiabRetinopathy,
		UnspecDiabRetinopathy:     data.ClinicianReportedOutcomes.ChronicComplications.UnspecDiabRetinopathy,
		MacularEdema:              data.ClinicianReportedOutcomes.ChronicComplications.MacularEdema,
		ChronicComplicOther:       data.ClinicianReportedOutcomes.ChronicComplications.ChronicComplicOther,
		NosightThreaten:           data.ClinicianReportedOutcomes.ChronicComplications.NosightThreaten,
		UnknownsightThreaten:      data.ClinicianReportedOutcomes.ChronicComplications.UnknownsightThreaten,
		AutonomicNeuropathy:       data.ClinicianReportedOutcomes.ChronicComplications.AutonomicNeuropathy,
		PeripheralNeuropathy:      data.ClinicianReportedOutcomes.ChronicComplications.PeripheralNeuropathy,
		PeripherNeuropSympt:       data.ClinicianReportedOutcomes.ChronicComplications.PeripheralNeuropathy,
		CharcotFoot:               data.ClinicianReportedOutcomes.ChronicComplications.CharcotFoot,
		LowerLimbUlcer:            data.ClinicianReportedOutcomes.ChronicComplications.LowerLimbUlcer,
		LLUSeverity:               data.ClinicianReportedOutcomes.ChronicComplications.LLUSeverity,
		SevereStageDiagnosed:      data.ClinicianReportedOutcomes.ChronicComplications.SevereStageDiagnosed,
		PeripheralArteryDisease:   data.ClinicianReportedOutcomes.ChronicComplications.PeripheralArteryDisease,
		PeripheralArterySymptoms:  data.ClinicianReportedOutcomes.ChronicComplications.PeripheralArterySymptoms,
		IschemicHeartDisease:      data.ClinicianReportedOutcomes.ChronicComplications.IschemicHeartDisease,
		HeartFailure:              data.ClinicianReportedOutcomes.ChronicComplications.HeartFailure,
		HeartFailureStage:         data.ClinicianReportedOutcomes.ChronicComplications.HeartFailureStage,
		EGFR:                      data.ClinicianReportedOutcomes.ChronicComplications.EGFR,
		UrinaryACR:                data.ClinicianReportedOutcomes.ChronicComplications.UrinaryACR,
		Dialysis:                  data.ClinicianReportedOutcomes.ChronicComplications.Dialysis,
		CerebrovascularDisease:    data.ClinicianReportedOutcomes.ChronicComplications.CerebrovascularDisease,
		PeriodontalHealth:         data.ClinicianReportedOutcomes.ChronicComplications.PeriodontalHealth,
		ErectileDysfunction:       data.ClinicianReportedOutcomes.ChronicComplications.ErectileDysfunction,
		Lipodystrophy:             data.ClinicianReportedOutcomes.ChronicComplications.Lipodystrophy,
	}
	if err := s.store.CreateDiagnosisProfile(ctx, &m); err != nil {
		s.logger.Error("failed to create Diagnosis Profile", zap.Error(err))
		return err
	}
	return nil
}
