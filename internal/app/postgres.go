package app

import (
	"context"
	"go.uber.org/zap"
	"github.com/jmoiron/sqlx"
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
	u:= UserModel{}
	q:= `SELECT * FROM users where id = $1`
	if err:= p.db.GetContext(ctx,&u,q,id);err!=nil {
		p.logger.Error("failed to get usermodel from db",zap.Error(err))
		return nil,err
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