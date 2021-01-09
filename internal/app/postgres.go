package app

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var _ Store = (*PostgresStore)(nil)

type PostgresStore struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewPostgresStore(logger *zap.Logger, db *sqlx.DB) *PostgresStore {
	return &PostgresStore{
		logger: logger,
		db:     db,
	}
}

func (p *PostgresStore) CreateUser(ctx context.Context, model *UserModel) error {
	q := `INSERT INTO users (
	id,
	short_id,
	name,
	email_id,
	phone_number,
	address,
	age,
	dob,
	sex,
	height,
	weight
	) VALUES (
	:id,
	:short_id,
	:name,
	:email_id,
	:phone_number,
	:address,
	:age,
	:dob,
	:sex,
	:height,
	:weight)`
	_, err := p.db.NamedExecContext(ctx, q, model)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresStore) GetUser(ctx context.Context, id string) (*UserModel, error) {
	u := UserModel{}
	q := `SELECT * FROM users where id = $1`
	if err := p.db.GetContext(ctx, &u, q, id); err != nil {
		p.logger.Error("failed to get usermodel from db", zap.Error(err))
		return nil, err
	}
	return &u, nil
}

func (p *PostgresStore) CreateMedicalData(ctx context.Context, model *MedicalDataModel) error {
	q := `INSERT INTO medical_profile (
	short_id,
	created_at,
	updated_at,
	year,
	sex
	) VALUES (
	:short_id,
	:created_at,
	:updated_at,
	:year,
	:sex)`
	_, err := p.db.NamedExecContext(ctx, q, model)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostgresStore) CreateDiagnosisProfile(ctx context.Context, model *DiagnosisProfileModel) error {
	q := `INSERT INTO Diagnosis_profile (
	short_id,
	created_at,
	updated_at,
	year,
	sex,
	diagnosed_with_diabetes,
	diab_type_indic,
	diabetic_diagnoseyear,
	comorb_liver_disease,
	comorb_maligancy
	comorb_aids,
	comorb_pulmonary_disease,
	comorb_peripheral_vascular_disease,
	comorb_dementia,
	comorb_hemiplegia,
	comorb_tuberculosis,
	comorb_hepatitisB,
	comorb_hepatitisC,
	comorb_anxietyDisorder,
	comorb_depression,
	comorb_disordered_eating,
	comorb_psychotic_mental_illness,
	comorb_thyroid,
	comorb_others,
	comorb_none,
	comorb_thyroid_harmone_level,
	treatvar_non_pharmacological_therapy,
	treatvar_oral_antidiabetic_drugs,
	treatvar_non_insulin_injection,
	treatvar_insulin,
	treatvar_other,
	treatvar_no_treatment,
	treatvar_bp_lower_medication,
treatvar_statins_lipid_therapy,
	clirptdoutc_glyc_hba1c_reading,
	clirptdoutc_glyc_hb1ac_reading_unit,
	clirptdoutc_glyc_percentage_range,
	clirptdoutc_intmoutc_systolic_bp,
	clirptdoutc_intmoutc_diastolic_bp,
	clirptdoutc_intmoutc_cholesterol,
	clirptdoutc_intmoutc_ldlCholesterol,
	clirptdoutc_intmoutc_ldlCholUnit,
	clirptdoutc_intmoutc_hdlCholesterol,
	clirptdoutc_intmoutc_hdlCholUnit,
	clirptdoutc_intmoutc_triglicerids,
	clirptdoutc_intmoutc_triglycUnit,
	clirptdoutc_intmoutc_weight,
	clirptdoutc_intmoutc_height,
	clirptdoutc_intmoutc_waist,
	clirptdoutc_acutevnt_dkahhs,
	clirptdoutc_acutevnt_hypoglycemia,
	clirptdoutc_acutevnt_hypoglycemia_l2exp,
	clirptdoutc_acutevnt_hypglycem_wosymp,
	clirptdoutc_acutevnt_acute_ischemic,
	clirptdoutc_acutevnt_acute_cerebro_vascular,
	clirptdoutc_acutevnt_limb_amputation,
	clirptdoutc_acutevnt_limb_amputation_level,
	clirptdoutc_chroniccomplic_visual_impairment,
	clirptdoutc_chroniccomplic_visual_acuity,
	clirptdoutc_chroniccomplic_measurement_mehtod,
	clirptdoutc_chroniccomplic_nonprofl_diab_retinopathy,
	clirptdoutc_chroniccomplic_profl_diab_retinopathy,
	clirptdoutc_chroniccomplic_unspec_diab_retinopathy,
	clirptdoutc_chroniccomplic_macular_edema,
	clirptdoutc_chroniccomplic_other,
	clirptdoutc_chroniccomplic_nosight_threaten,
	clirptdoutc_chroniccomplic_unknown_sight_threaten,
	clirptdoutc_chroniccomplic_autonomic_neuropathy,
	clirptdoutc_chroniccomplic_peripheral_neuropathy,
	clirptdoutc_chroniccomplic_peripher_neurop_sympt,
	clirptdoutc_chroniccomplic_charcot_foot,
	clirptdoutc_chroniccomplic_lower_limb_ulcer,
	clirptdoutc_chroniccomplic_llu_severity,
	clirptdoutc_chroniccomplic_severe_stage_diagnosed,
	clirptdoutc_chroniccomplic_peripheral_artery_disease,
	clirptdoutc_chroniccomplic_peripheral_artery_symptoms,
	clirptdoutc_chroniccomplic_ischemic_heart_disease,
	clirptdoutc_chroniccomplic_heart_failure,
	clirptdoutc_chroniccomplic_heart_failure_stage,
	clirptdoutc_chroniccomplic_egfr,
	clirptdoutc_chroniccomplic_urinary_ACR,
	clirptdoutc_chroniccomplic_dialysis,
	clirptdoutc_chroniccomplic_cerebrovascular_disease,
	clirptdoutc_chroniccomplic_periodontal_health,
	clirptdoutc_chroniccomplic_erectile_dysfunction,
	clirptdoutc_chroniccomplic_lipodystrophy
	) VALUES (
	:short_id,
	:created_at,
	:updated_at,
	:year,
	:sex,
	:diagnosed_with_diabetes,
	:diab_type_indic,
	:diabetic_diagnoseyear,
	:comorb_liver_disease,
	:comorb_maligancy
	:comorb_aids,
	:comorb_pulmonary_disease,
	:comorb_peripheral_vascular_disease,
	:comorb_dementia,
	:comorb_hemiplegia,
	:comorb_tuberculosis,
	:comorb_hepatitisB,
	:comorb_hepatitisC,
	:comorb_anxietyDisorder,
	:comorb_depression,
	:comorb_disordered_eating,
	:comorb_psychotic_mental_illness,
	:comorb_thyroid,
	:comorb_others,
	:comorb_none,
	:comorb_thyroid_harmone_level,
	:treatvar_non_pharmacological_therapy,
	:treatvar_oral_antidiabetic_drugs,
	:treatvar_non_insulin_injection,
	:treatvar_insulin,
	:treatvar_other,
	:treatvar_no_treatment,
	:treatvar_bp_lower_medication,
:treatvar_statins_lipid_therapy,
	:clirptdoutc_glyc_hba1c_reading,
	:clirptdoutc_glyc_hb1ac_reading_unit,
	:clirptdoutc_glyc_percentage_range,
	:clirptdoutc_intmoutc_systolic_bp,
	:clirptdoutc_intmoutc_diastolic_bp,
	:clirptdoutc_intmoutc_cholesterol,
	:clirptdoutc_intmoutc_ldlCholesterol,
	:clirptdoutc_intmoutc_ldlCholUnit,
	:clirptdoutc_intmoutc_hdlCholesterol,
	:clirptdoutc_intmoutc_hdlCholUnit,
	:clirptdoutc_intmoutc_triglicerids,
	:clirptdoutc_intmoutc_triglycUnit,
	:clirptdoutc_intmoutc_weight,
	:clirptdoutc_intmoutc_height,
	:clirptdoutc_intmoutc_waist,
	:clirptdoutc_acutevnt_dkahhs,
	:clirptdoutc_acutevnt_hypoglycemia,
	:clirptdoutc_acutevnt_hypoglycemia_l2exp,
	:clirptdoutc_acutevnt_hypglycem_wosymp,
	:clirptdoutc_acutevnt_acute_ischemic,
	:clirptdoutc_acutevnt_acute_cerebro_vascular,
	:clirptdoutc_acutevnt_limb_amputation,
	:clirptdoutc_acutevnt_limb_amputation_level,
	:clirptdoutc_chroniccomplic_visual_impairment,
	:clirptdoutc_chroniccomplic_visual_acuity,
	:clirptdoutc_chroniccomplic_measurement_mehtod,
	:clirptdoutc_chroniccomplic_nonprofl_diab_retinopathy,
	:clirptdoutc_chroniccomplic_profl_diab_retinopathy,
	:clirptdoutc_chroniccomplic_unspec_diab_retinopathy,
	:clirptdoutc_chroniccomplic_macular_edema,
	:clirptdoutc_chroniccomplic_other,
	:clirptdoutc_chroniccomplic_nosight_threaten,
	:clirptdoutc_chroniccomplic_unknown_sight_threaten,
	:clirptdoutc_chroniccomplic_autonomic_neuropathy,
	:clirptdoutc_chroniccomplic_peripheral_neuropathy,
	:clirptdoutc_chroniccomplic_peripher_neurop_sympt,
	:clirptdoutc_chroniccomplic_charcot_foot,
	:clirptdoutc_chroniccomplic_lower_limb_ulcer,
	:clirptdoutc_chroniccomplic_llu_severity,
	:clirptdoutc_chroniccomplic_severe_stage_diagnosed,
	:clirptdoutc_chroniccomplic_peripheral_artery_disease,
	:clirptdoutc_chroniccomplic_peripheral_artery_symptoms,
	:clirptdoutc_chroniccomplic_ischemic_heart_disease,
	:clirptdoutc_chroniccomplic_heart_failure,
	:clirptdoutc_chroniccomplic_heart_failure_stage,
	:clirptdoutc_chroniccomplic_egfr,
	:clirptdoutc_chroniccomplic_urinary_ACR,
	:clirptdoutc_chroniccomplic_dialysis,
	:clirptdoutc_chroniccomplic_cerebrovascular_disease,
	:clirptdoutc_chroniccomplic_periodontal_health,
	:clirptdoutc_chroniccomplic_erectile_dysfunction,
	:clirptdoutc_chroniccomplic_lipodystrophy)`
	_, err := p.db.NamedExecContext(ctx, q, model)
	if err != nil {
		return err
	}
	return nil
}
