package repo_candidate

import (
	"context"
	"database/sql"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (c *candidateRepository) UpdateUserDataWithTx(ctx context.Context, tx *sql.Tx, p model.UserData) error {
	const q = `UPDATE user_data
			SET first_name = ?, middle_name = ?, last_name = ?, nickname = ?,gender = ?,birthdate = ? , current_province = ? ,current_regency = ?, last_edited_on = ?
			WHERE user_data.staf_id = ?;`

	_, err := tx.ExecContext(ctx, q, p.FirstName, p.MiddleName, p.LastName, p.Nickname, p.Gender, p.Birthdate, p.CurrentProvince, p.CurrentRegency, time.Now(), p.StafID)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserJobCandidatePortal(ctx context.Context, p model.UserJobCandidatePortal) error {
	const q = `UPDATE job_candidate_portal
			SET uuid = ?, jcp_name = ?, jcp_summary = ?,  jcp_salary = ?,  jcp_salary_unit = ?,  jcp_from = ?, jcp_date_apply = ?
			WHERE job_candidate_portal.staf_id = ?;`

	_, err := c.db.ExecContext(ctx, q, p.UUID, p.JCPName, p.JCPSummary, p.JCPSalary, p.JCPSalaryUnit, p.JCPFrom, p.JCPDateApply)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserFormalEducationWithTx(ctx context.Context, tx *sql.Tx, p model.UserFormalEducation) error {
	const q = `UPDATE user_formal_education
			SET institution_name = ?, major = ?, graduation_year = ?
			WHERE user_formal_education.id_staff = ? AND user_formal_education.id_master_education = ?;`

	_, err := tx.ExecContext(ctx, q, p.InstitutionName, p.Major, p.GraduationYear, p.IDStaff, p.IDMasterEducation)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserWorkExperienceWithTx(ctx context.Context, tx *sql.Tx, p model.UserWorkExperience) error {
	const q = `UPDATE user_work_experience
			SET organization_name = ?, period_from = ?, period_to = ?, position = ?, job_desc = ?
			WHERE user_work_experience.id_user = ? AND organization_name = ? AND position = ?;`

	_, err := tx.ExecContext(ctx, q, p.OrganizationName, p.PeriodFrom, p.PeriodTo, p.Position, p.JobDesc, p.IDUser, p.OrganizationName, p.Position)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserContact(ctx context.Context, p model.UserContact) error {
	const q = `UPDATE user_contact
			SET contact_caption = ?, contact_label = ?, contact_type = ?
			WHERE user_contact.id_user = ?;`

	_, err := c.db.ExecContext(ctx, q, p.ContactCaption, p.ContactLabel, p.ContactType)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserFileWithTx(ctx context.Context, tx *sql.Tx, p model.UserFile) error {
	const q = `UPDATE user_file
			SET file_attachment = ?
			WHERE user_file.id_user = ?;`

	_, err := tx.ExecContext(ctx, q, p.FileAttachment, p.IDUser)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateUserPhotoWithTx(ctx context.Context, tx *sql.Tx, p model.UserData) error {
	const q = `UPDATE user_data
			SET user_photo = ?
			WHERE user_data.staf_id = ?;`

	_, err := tx.ExecContext(ctx, q, p.UserPhoto, p.StafID)
	if err != nil {
		return err
	}

	return nil
}
