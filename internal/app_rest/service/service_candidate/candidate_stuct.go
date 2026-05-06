package service_candidate

import "mime/multipart"

type StoreCandidatePayload struct {
	Channel           string                `form:"channel"`
	Type              string                `form:"type"`
	AppliedFor        string                `form:"applied_for"`
	AppliedForId      string                `form:"applied_for_id"`
	AppliedDate       string                `form:"applied_date"`
	ProfileLink       string                `form:"profile_link"`
	Email             string                `form:"email"`
	Fullname          string                `form:"fullname"`
	Nickname          string                `form:"nickname"`
	Photo             *multipart.FileHeader `form:"photo"`
	Gender            string                `form:"gender"`
	DateOfBirth       string                `form:"date_of_birth"`
	Age               string                `form:"age"`
	Contact           string                `form:"contact"`
	Summary           string                `form:"summary"`
	LatestSalary      string                `form:"latest_salary"`
	SalaryExpectation string                `form:"salary_expectation"`
	WorkExperiences   string                `form:"work_experiences"`
	Educations        string                `form:"educations"`
	Skills            string                `form:"skills"`
	Location          string                `form:"location"`
	ReferenceLinks    string                `form:"reference_links""`
	Cv                *multipart.FileHeader `form:"cv"`
}

type PayloadEducations []Education

type Education struct {
	Education       string `json:"education"`
	Institution     string `json:"institution"`
	Major           string `json:"major"`
	PeriodStartYear string `json:"period_start_year"`
	PeriodEndYear   string `json:"period_end_year"`
}

type PayloadWorkExperiences []WorkExperience

type WorkExperience struct {
	Position     string `json:"position"`
	Organization string `json:"organization"`
	JobDesc      string `json:"job_desc"`
	PeriodFrom   string `json:"period_from"`
	PeriodTo     string `json:"period_to"`
}

type PayloadContact struct {
	Type          string `json:"type"`
	ContactNumber string `json:"contact_number"`
}

type PayloadSkills []string

type PayloadReferenceLinks []ReferenceLink

type ReferenceLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
