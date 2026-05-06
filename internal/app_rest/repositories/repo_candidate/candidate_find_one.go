package repo_candidate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"helicopter-hr/internal/app_rest/model"
)

func (c *candidateRepository) FindOneUserHasApplied(ctx context.Context, p model.Candidate) (*model.Candidate, error) {
	const q = `SELECT staf_id
	FROM job_candidate_portal WHERE 
    job_candidate_portal.staf_id = ? AND
	job_candidate_portal.jcp_name = ? AND 
	job_candidate_portal.jcp_from = ? AND 
	job_candidate_portal.jcp_date_apply = ?;`

	var result model.Candidate
	err := c.db.QueryRowContext(ctx, q, p.StafID, p.AppliedFor, p.Channel, p.AppliedDate).Scan(&result.StafID)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindOneUserByEmailWithTx(ctx context.Context, tx *sql.Tx, email string) (*model.UserAuth, error) {
	const q = `SELECT staf_id, auth_email FROM user_auth WHERE user_auth.auth_email = ?;`

	var result model.UserAuth
	err := tx.QueryRowContext(ctx, q, email).Scan(&result.StafID, &result.AuthEmail)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindOneUserByUrlProfileWithTx(ctx context.Context, tx *sql.Tx, url_profile string) (*model.UserJobCandidatePortal, error) {
	const q = `SELECT jcp_id,staf_id FROM job_candidate_portal WHERE job_candidate_portal.jcp_url_profile = ?;`

	var result model.UserJobCandidatePortal
	err := tx.QueryRowContext(ctx, q, url_profile).Scan(&result.ID, &result.StafID)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindOneUserDataByStaffID(ctx context.Context, staffID int) (*model.UserData, error) {
	const q = `SELECT staf_id,first_name,middle_name,last_name,nickname,gender,birthdate,current_province,current_regency
			FROM user_data WHERE user_data.staf_id = ?;`

	var result model.UserData
	err := c.db.QueryRowContext(ctx, q, staffID).
		Scan(&result.StafID,
			&result.FirstName,
			&result.MiddleName,
			&result.LastName,
			&result.Nickname,
			&result.Gender,
			&result.Birthdate,
			&result.CurrentProvince,
			&result.CurrentRegency)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindOneUserJobCandidatePortalByStaffID(ctx context.Context, staffID int) (*model.UserJobCandidatePortal, error) {
	const q = `SELECT staf_id,uuid,jcp_name,jcp_summary,jcp_salary,jcp_salary_unit,jcp_from,jcp_date_apply
			FROM job_candidate_portal WHERE job_candidate_portal.staf_id = ?;`

	var result model.UserJobCandidatePortal
	err := c.db.QueryRowContext(ctx, q, staffID).
		Scan(&result.StafID,
			&result.UUID,
			&result.JCPName,
			&result.JCPSummary,
			&result.JCPSalary,
			&result.JCPSalaryUnit,
			&result.JCPFrom,
			&result.JCPDateApply)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindGender(ctx context.Context, gender string) (*model.MasterGender, error) {
	const q = `SELECT id, gender FROM master_gender where master_gender.gender = ?;`

	var result model.MasterGender
	err := c.db.QueryRowContext(ctx, q, gender).Scan(&result.ID, &result.Gender)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindEducation(ctx context.Context, code string) (*model.MasterEducation, error) {
	const q = `SELECT id, education, level, code FROM master_education where master_education.code = ?;`

	var result model.MasterEducation
	err := c.db.QueryRowContext(ctx, q, code).Scan(&result.ID, &result.Education, &result.Level, &result.Code)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindUserFormalEducation(ctx context.Context, payload model.UserFormalEducation) (*model.UserFormalEducation, error) {
	const q = `SELECT id, id_staff, id_master_education, institution_name, graduation_year 
    FROM user_formal_education where user_formal_education.id_staff = ? AND user_formal_education.id_master_education = ?;`

	var result model.UserFormalEducation
	err := c.db.QueryRowContext(ctx, q, payload.IDStaff, payload.IDMasterEducation).Scan(&result.ID, &result.IDStaff, &result.IDMasterEducation, &result.InstitutionName, &result.GraduationYear)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindUserWorkExperience(ctx context.Context, payload model.UserWorkExperience) (*model.UserWorkExperience, error) {
	q := fmt.Sprintf(`SELECT id, id_user,organization_name,position,job_desc FROM user_work_experience `+
		`WHERE user_work_experience.id_user = %d AND user_work_experience.organization_name = '%s' AND user_work_experience.position = '%s' AND user_work_experience.job_desc = '%s';`,
		payload.IDUser, payload.OrganizationName, payload.Position, payload.JobDesc)

	var result model.UserWorkExperience
	err := c.db.QueryRowContext(ctx, q).
		Scan(&result.ID,
			&result.IDUser,
			&result.OrganizationName,
			&result.Position,
			&result.JobDesc)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindUserContact(ctx context.Context, payload model.UserContact) (*model.UserContact, error) {
	const q = `SELECT id, id_user,contact_caption,contact_label,contact_type
    FROM user_contact where user_contact.id_user = ? AND user_contact.contact_caption = ? AND user_contact.contact_label = ? AND user_contact.contact_type = ?;`

	var result model.UserContact
	err := c.db.QueryRowContext(ctx, q, payload.IDUser, payload.ContactCaption, payload.ContactLabel, payload.ContactType).
		Scan(&result.ID,
			&result.IDUser,
			&result.ContactCaption,
			&result.ContactLabel,
			&result.ContactType)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindUserFileWithTx(ctx context.Context, tx *sql.Tx, payload model.UserFile) (*model.UserFile, error) {
	const q = `SELECT id, id_user,id_file_category,file_attachment FROM user_file where user_file.id_user = ?;`

	var result model.UserFile
	err := tx.QueryRowContext(ctx, q, payload.IDUser).
		Scan(&result.ID,
			&result.IDUser,
			&result.IDFileCategory,
			&result.FileAttachment)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindUserSkillWithTx(ctx context.Context, tx *sql.Tx, payload model.UserJobCandidatePortalSkill) (*model.UserJobCandidatePortalSkill, error) {
	const q = `SELECT jcps_id, staf_id, jcps_skill FROM job_candidate_portal_skill WHERE job_candidate_portal_skill.staf_id = ? AND job_candidate_portal_skill.jcps_skill = ?;`

	var result model.UserJobCandidatePortalSkill
	err := tx.QueryRowContext(ctx, q, payload.StafID, payload.JCPSSkill).
		Scan(&result.JCPSID,
			&result.StafID,
			&result.JCPSSkill)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindMasterProvince(ctx context.Context, province string) (*model.MasterProvince, error) {
	q := fmt.Sprintf("SELECT id,name FROM locat_provinces WHERE locat_provinces.name LIKE %s;", "'%"+province+"%'")

	var result model.MasterProvince
	err := c.db.QueryRowContext(ctx, q).
		Scan(&result.ID,
			&result.Name)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}

func (c *candidateRepository) FindMasterRegency(ctx context.Context, regency string) (*model.MasterRegency, error) {
	q := fmt.Sprintf("SELECT id,name FROM locat_regencies WHERE locat_regencies.name LIKE %s;", "'%"+regency+"%'")

	var result model.MasterRegency
	err := c.db.QueryRowContext(ctx, q).
		Scan(&result.ID,
			&result.Name)
	if err != nil {
		if err.Error() == model.NotFound {
			return nil, errors.New("notFound")
		}
		return nil, err
	}

	return &result, nil
}
