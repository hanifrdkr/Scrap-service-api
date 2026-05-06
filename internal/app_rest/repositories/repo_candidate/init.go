package repo_candidate

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"helicopter-hr/internal/app_rest/model"
)

type CandidateRepositoryInterface interface {
	NewBeginTx(ctx context.Context) (*sql.Tx, error)
	FindOneUserHasApplied(ctx context.Context, p model.Candidate) (*model.Candidate, error)
	FindOneUserByEmailWithTx(ctx context.Context, tx *sql.Tx, email string) (*model.UserAuth, error)
	FindOneUserByUrlProfileWithTx(ctx context.Context, tx *sql.Tx, url_profile string) (*model.UserJobCandidatePortal, error)
	FindGender(ctx context.Context, gender string) (*model.MasterGender, error)
	FindEducation(ctx context.Context, code string) (*model.MasterEducation, error)
	FindUserFormalEducation(ctx context.Context, payload model.UserFormalEducation) (*model.UserFormalEducation, error)
	FindOneUserDataByStaffID(ctx context.Context, staffID int) (*model.UserData, error)
	FindUserWorkExperience(ctx context.Context, payload model.UserWorkExperience) (*model.UserWorkExperience, error)
	FindUserContact(ctx context.Context, payload model.UserContact) (*model.UserContact, error)
	FindOneUserJobCandidatePortalByStaffID(ctx context.Context, staffID int) (*model.UserJobCandidatePortal, error)
	FindUserFileWithTx(ctx context.Context, tx *sql.Tx, payload model.UserFile) (*model.UserFile, error)
	FindUserSkillWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortalSkill) (*model.UserJobCandidatePortalSkill, error)
	FindMasterProvince(ctx context.Context, province string) (*model.MasterProvince, error)
	FindMasterRegency(ctx context.Context, regency string) (*model.MasterRegency, error)
	StoreUserAuthWithTx(ctx context.Context, tx *sql.Tx, payload model.Candidate) error
	StoreUserDataWithTx(ctx context.Context, tx *sql.Tx, payload model.Candidate) error
	StoreUserFormalEducationWithTx(ctx context.Context, tx *sql.Tx, payload model.Education) error
	StoreUserWorkExperienceWithTx(ctx context.Context, tx *sql.Tx, payload model.UserWorkExperience) error
	StoreUserContactWithTx(ctx context.Context, tx *sql.Tx, payload model.UserContact) error
	StoreUserJobCandidatePortalWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortal) error
	StoreUserFileWithTx(ctx context.Context, tx *sql.Tx, payload model.UserFile) error
	StoreUserJobCandidatePortalSkillWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortalSkill) error
	UpdateUserDataWithTx(ctx context.Context, tx *sql.Tx, p model.UserData) error
	UpdateUserFormalEducationWithTx(ctx context.Context, tx *sql.Tx, p model.UserFormalEducation) error
	UpdateUserWorkExperienceWithTx(ctx context.Context, tx *sql.Tx, p model.UserWorkExperience) error
	UpdateUserContact(ctx context.Context, p model.UserContact) error
	UpdateUserJobCandidatePortal(ctx context.Context, p model.UserJobCandidatePortal) error
	UpdateUserFileWithTx(ctx context.Context, tx *sql.Tx, p model.UserFile) error
	UpdateUserPhotoWithTx(ctx context.Context, tx *sql.Tx, p model.UserData) error
}

type candidateRepository struct {
	db *sqlx.DB
}

func NewCandidateRepository(db *sqlx.DB) CandidateRepositoryInterface {
	return &candidateRepository{
		db: db,
	}
}

func (c *candidateRepository) NewBeginTx(ctx context.Context) (*sql.Tx, error) {
	return c.db.BeginTx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelSerializable, ReadOnly: false,
		})
}
