package model

import (
	"database/sql"
	"mime/multipart"
	"time"
)

type UserAuth struct {
	StafID    int    `json:"staf_id"`
	AuthEmail string `json:"auth_email"`
}

type UserData struct {
	StafID          int            `json:"staf_id"`
	FirstName       string         `json:"first_name"`
	MiddleName      string         `json:"middle_name"`
	LastName        string         `json:"last_name"`
	Nickname        string         `json:"nickname"`
	UserPhoto       string         `json:"user_photo"`
	Gender          int            `json:"gender"`
	Birthdate       sql.NullString `json:"birthdate"`
	CurrentProvince string         `json:"current_province"`
	CurrentRegency  string         `json:"current_regency"`
}

type Candidate struct {
	StafID            int                  `json:"staf_id"`
	Channel           string               `json:"channel"`
	Type              string               `json:"type"`
	AppliedFor        string               `json:"applied_for"`
	AppliedForId      string               `json:"applied_for_id"`
	AppliedDate       time.Time            `json:"applied_date"`
	ProfileLink       string               `json:"profile_link"`
	Email             string               `json:"email"`
	FirstName         string               `json:"first_name"`
	MiddleName        string               `json:"middle_name"`
	LastName          string               `json:"last_name"`
	Nickname          string               `json:"nickname"`
	Photo             multipart.FileHeader `json:"photo"`
	Gender            int                  `json:"gender"`
	DateOfBirth       sql.NullTime         `json:"date_of_birth"`
	Age               string               `json:"age"`
	Contact           string               `json:"contact"`
	Summary           string               `json:"summary"`
	LatestSalary      string               `json:"latest_salary"`
	SalaryExpectation string               `json:"salary_expectation"`
	WorkExperiences   string               `json:"work_experiences"`
	Educations        []Education          `json:"educations"`
	Skills            []string             `json:"skills"`
	CurrentProvince   string               `json:"current_province"`
	CurrentRegency    string               `json:"current_regency"`
	ReferenceLinks    string               `json:"reference_links""`
	Cv                multipart.FileHeader `json:"cv"`
}

type UserJobCandidatePortal struct {
	ID            int       `json:"id"`
	StafID        int       `json:"staf_id"`
	UUID          string    `json:"uuid"`
	JCPName       string    `json:"jcp_name"`
	JCPURLProfile string    `json:"jcp_url_profile"`
	JCPSummary    string    `json:"jcp_summary"`
	JCPSalary     int       `json:"jcp_salary"`
	JCPSalaryUnit string    `json:"jcp_salary_unit"`
	JCPFrom       string    `json:"jcp_from"`
	JCPDateApply  time.Time `json:"jcp_date_apply"`
}

type UserJobCandidatePortalSkill struct {
	JCPSID    int    `json:"jcps_id"`
	StafID    int    `json:"staf_id"`
	JCPSSkill string `json:"jcps_skill"`
}

type (
	Contact struct {
		Type          string `json:"type"`
		ContactNumber string `json:"contact_number"`
	}

	UserContact struct {
		ID             int    `json:"id"`
		IDUser         int    `json:"id_user"`
		ContactCaption string `json:"contact_caption"`
		ContactLabel   string `json:"contact_label"`
		ContactType    string `json:"contact_type"`
	}
)

type (
	WorkExperience struct {
		Position     string `json:"position"`
		Organization string `json:"organization"`
		JobDesc      string `json:"job_desc"`
		PeriodFrom   string `json:"period_from"`
		PeriodTo     string `json:"period_to"`
	}

	UserWorkExperience struct {
		ID               int       `json:"id"`
		IDUser           int       `json:"id_user"`
		OrganizationName string    `json:"organization_name"`
		PeriodFrom       time.Time `json:"period_from"`
		PeriodTo         time.Time `json:"period_to"`
		Position         string    `json:"position"`
		JobDesc          string    `json:"job_desc"`
	}
)

type (
	Education struct {
		StafID            int    `json:"staf_id"`
		IDMasterEducation int    `json:"id_master_education"`
		Education         string `json:"education"`
		Institution       string `json:"institution"`
		Major             string `json:"major"`
		PeriodStartYear   string `json:"period_start_year"`
		PeriodEndYear     string `json:"period_end_year"`
	}

	MasterEducation struct {
		ID        int    `json:"id"`
		Education string `json:"education"`
		Level     string `json:"level"`
		Code      string `json:"code"`
	}

	UserFormalEducation struct {
		ID                int    `json:"id"`
		IDStaff           int    `json:"id_staff"`
		IDMasterEducation int    `json:"id_master_education"`
		InstitutionName   string `json:"institution_name"`
		Major             string `json:"major"`
		GraduationYear    string `json:"graduation_year"`
	}
)

type MasterGender struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
}

type ReferenceLinks struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type UserFile struct {
	ID             int    `json:"id"`
	IDUser         int    `json:"id_user"`
	IDFileCategory int    `json:"id_file_category"`
	FileAttachment string `json:"file_attachment"`
}

type MasterProvince struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MasterRegency struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
