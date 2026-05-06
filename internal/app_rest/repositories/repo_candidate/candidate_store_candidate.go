package repo_candidate

import (
	"context"
	"database/sql"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (c *candidateRepository) StoreUserAuthWithTx(ctx context.Context, tx *sql.Tx, payload model.Candidate) error {
	const q = `INSERT INTO user_auth (staf_nama,auth_email,auth_password,active_date,auth_status,auth_group,auth_grading,bypass_koordinat,web_absen)
			VALUES (?,?,?,?,?,?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		"-",
		payload.Email,
		"",
		time.Now(),
		10,
		0,
		0,
		0,
		0,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserDataWithTx(ctx context.Context, tx *sql.Tx, payload model.Candidate) error {
	const q = `INSERT INTO user_data (staf_id,first_name,middle_name,last_name,nickname,user_photo,gender,birthdate,current_province,current_regency,came_from,created_on,
                       user_status,religion,marital_status,id_card_number,id_card_address,id_card_postal_code,current_address,current_postal_code,id_tax_number,id_tax_address)
			VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.StafID,
		payload.FirstName,
		payload.MiddleName,
		payload.LastName,
		payload.Nickname,
		"nofoto.png",
		payload.Gender,
		payload.DateOfBirth,
		payload.CurrentProvince,
		payload.CurrentRegency,
		payload.Channel,
		time.Now(),
		0,
		0,
		0,
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserFormalEducationWithTx(ctx context.Context, tx *sql.Tx, payload model.Education) error {
	const q = `INSERT INTO user_formal_education (id_staff,id_master_education,institution_name,major,graduation_year)
			VALUES (?,?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.StafID,
		payload.IDMasterEducation,
		payload.Institution,
		payload.Major,
		payload.PeriodEndYear,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserWorkExperienceWithTx(ctx context.Context, tx *sql.Tx, payload model.UserWorkExperience) error {
	const q = `INSERT INTO user_work_experience (id_user,organization_name,period_from,period_to,position,job_desc)
			VALUES (?,?,?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.IDUser,
		payload.OrganizationName,
		payload.PeriodFrom,
		payload.PeriodTo,
		payload.Position,
		payload.JobDesc,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserContactWithTx(ctx context.Context, tx *sql.Tx, payload model.UserContact) error {
	const q = `INSERT INTO user_contact(id_user,contact_caption,contact_label,contact_type)
			VALUES (?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.IDUser,
		payload.ContactCaption,
		payload.ContactLabel,
		payload.ContactType,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserJobCandidatePortalWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortal) error {
	const q = `INSERT INTO job_candidate_portal(staf_id,uuid,jcp_name,jcp_url_profile,jcp_summary,jcp_salary,jcp_salary_unit,jcp_from,jcp_date_apply)
			VALUES (?,?,?,?,?,?,?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.StafID,
		payload.UUID,
		payload.JCPName,
		payload.JCPURLProfile,
		payload.JCPSummary,
		payload.JCPSalary,
		payload.JCPSalaryUnit,
		payload.JCPFrom,
		payload.JCPDateApply,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserJobCandidatePortalSkillWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortalSkill) error {
	const q = `INSERT INTO job_candidate_portal_skill(staf_id,jcps_skill) VALUES (?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.StafID,
		payload.JCPSSkill,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) StoreUserFileWithTx(ctx context.Context, tx *sql.Tx, payload model.UserFile) error {
	const q = `INSERT INTO user_file(id_user,id_file_category,file_attachment)
			VALUES (?,?,?);`

	_, err := tx.ExecContext(
		ctx,
		q,
		payload.IDUser,
		payload.IDFileCategory,
		payload.FileAttachment,
	)

	if err != nil {
		return err
	}

	return nil
}
