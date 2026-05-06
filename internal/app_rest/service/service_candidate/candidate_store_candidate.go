package service_candidate

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/model"
	"helicopter-hr/internal/app_rest/util"
	"html"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	errTransactionDatabase = errors.New("failure database transaction")
)

func (c *candidateService) StoreCandidate(ctx context.Context, payload StoreCandidatePayload) error {
	var (
		guid                   = ctx.Value("request_id").(string)
		payloadEducations      PayloadEducations
		payloadWorkExperiences PayloadWorkExperiences
		payloadContact         PayloadContact
		payloadReferenceLinks  PayloadReferenceLinks
		payloadSkills          PayloadSkills
		respUser               *model.UserAuth
		cfg                    = c.config
	)
	cLogger := zap.L().With(
		zap.String("layer", "service.candidate.store"),
		zap.String("request_id", guid),
	)

	// Get a Tx for making transaction requests.
	tx, err := c.repoCandidate.NewBeginTx(ctx)
	if err != nil {
		cLogger.Error("error create transaction database ", zap.Error(err))
		return errTransactionDatabase
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	if payload.Type == "recomendation" {
		// find url profile
		respUserRecomendation, err := c.repoCandidate.FindOneUserByUrlProfileWithTx(ctx, tx, payload.ProfileLink)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}
		if respUserRecomendation != nil {
			return errors.New("candidate_already_applied")
		}

	} else {
		// find email
		respUser, err = c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}
	}
	if respUser != nil {
		appliedDate, _ := time.Parse("2006-01-02", payload.AppliedDate)
		respUserApplied, err := c.repoCandidate.FindOneUserHasApplied(ctx, model.Candidate{
			StafID:      respUser.StafID,
			Channel:     payload.Channel,
			AppliedFor:  payload.AppliedFor,
			AppliedDate: appliedDate,
		})
		if err != nil && err.Error() != "notFound" {
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}
		if respUserApplied != nil {
			return errors.New("candidate_already_applied")
		}

		// update user data
		dataUserUpdate := c.ValidateDataUser(ctx, respUser.StafID, payload)

		if dataUserUpdate != nil {
			err = c.repoCandidate.UpdateUserDataWithTx(ctx, tx, model.UserData{
				StafID:          respUser.StafID,
				FirstName:       dataUserUpdate.FirstName,
				MiddleName:      dataUserUpdate.MiddleName,
				LastName:        dataUserUpdate.LastName,
				Nickname:        dataUserUpdate.Nickname,
				Gender:          dataUserUpdate.Gender,
				Birthdate:       dataUserUpdate.Birthdate,
				CurrentProvince: dataUserUpdate.CurrentProvince,
				CurrentRegency:  dataUserUpdate.CurrentRegency,
			})
			if err != nil {
				tx.Rollback()
				cLogger.Error("error update table user data", zap.Error(err))
				return err
			}
		}

		if dataUserUpdate == nil {
			// insert table user data
			err = c.doInsertUserDataWithTx(ctx, tx, payload)
			if err != nil {
				tx.Rollback()
				cLogger.Error("error store table user data", zap.Error(err))
				return err
			}
		}

	}

	// insert data candidate
	if respUser == nil {
		if payload.Email == "" {
			payload.Email = GenerateRandomString()
		}
		// store table user auth
		err = c.repoCandidate.StoreUserAuthWithTx(ctx, tx, model.Candidate{
			Email: payload.Email,
		})
		if err != nil {
			tx.Rollback()
			cLogger.Error("error store table user auth", zap.Error(err))
			return err
		}
		// insert table user data
		err = c.doInsertUserDataWithTx(ctx, tx, payload)
		if err != nil {
			tx.Rollback()
			cLogger.Error("error store table user data", zap.Error(err))
			return err
		}
	}

	// store table user_formal_education
	if payload.Educations != "" {
		err := json.Unmarshal([]byte(payload.Educations), &payloadEducations)
		if err != nil {
			tx.Rollback()
			cLogger.Error("error payload education", zap.Error(err))
			return errors.New("invalid_format_education_payload")
		}

		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}

		for _, v := range payloadEducations {
			// find education
			respEducation, err := c.repoCandidate.FindEducation(ctx, v.Education)
			if err != nil && err.Error() != "notFound" {
				tx.Rollback()
				cLogger.Error("error find gender ", zap.Error(err))
				return err
			}

			if respEducation == nil {
				cLogger.Error("error find user formal education ", zap.Error(err))
				return errors.New("invalid_code_education")
			}

			respUserFormalEducation, err := c.repoCandidate.FindUserFormalEducation(ctx, model.UserFormalEducation{
				IDStaff:           respUserData.StafID,
				IDMasterEducation: respEducation.ID,
			})
			if err != nil && err.Error() != "notFound" {
				tx.Rollback()
				cLogger.Error("error find user formal education ", zap.Error(err))
				return err
			}

			if respUserFormalEducation != nil {
				// update data user formal education
				err = c.repoCandidate.UpdateUserFormalEducationWithTx(ctx, tx, model.UserFormalEducation{
					IDStaff:           respUserData.StafID,
					IDMasterEducation: respEducation.ID,
					InstitutionName:   v.Institution,
					Major:             v.Major,
					GraduationYear:    v.PeriodEndYear,
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error update table user formal education", zap.Error(err))
					return err
				}
			} else {
				// store table user_formal_education
				err = c.repoCandidate.StoreUserFormalEducationWithTx(ctx, tx, model.Education{
					StafID:            respUserData.StafID,
					IDMasterEducation: respEducation.ID,
					Institution:       v.Institution,
					Major:             v.Major,
					PeriodEndYear:     v.PeriodEndYear,
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error store table user formal education", zap.Error(err))
					return err
				}

			}
		}
	}

	// store table user_work_experience
	if payload.WorkExperiences != "" {
		err := json.Unmarshal([]byte(payload.WorkExperiences), &payloadWorkExperiences)
		if err != nil {
			tx.Rollback()
			cLogger.Error("error payload work experience", zap.Error(err))
			return errors.New("invalid_format_work_experiences_payload")
		}

		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}

		for _, v := range payloadWorkExperiences {
			periodFrom, err := time.Parse("2006-01-02", v.PeriodFrom)
			if err != nil || v.PeriodFrom == "" || v.PeriodFrom == "-" || v.PeriodFrom == "0" {
				periodFrom, _ = time.Parse("2006-01-02", "1970-01-01")
			}
			periodTo, err := time.Parse("2006-01-02", v.PeriodTo)
			if err != nil || v.PeriodTo == "" || v.PeriodTo == "-" || v.PeriodFrom == "0" {
				periodTo, _ = time.Parse("2006-01-02", "1970-01-01")
			}
			respUserWorkExperience, err := c.repoCandidate.FindUserWorkExperience(ctx, model.UserWorkExperience{
				IDUser:           respUserData.StafID,
				OrganizationName: v.Organization,
				PeriodFrom:       periodFrom,
				PeriodTo:         periodTo,
				Position:         v.Position,
				JobDesc:          html.EscapeString(v.JobDesc),
			})
			if err != nil && err.Error() != "notFound" {
				tx.Rollback()
				cLogger.Error("error find user work experience ", zap.Error(err))
				return err
			}

			if respUserWorkExperience != nil {
				// update data user work experience
				err = c.repoCandidate.UpdateUserWorkExperienceWithTx(ctx, tx, model.UserWorkExperience{
					IDUser:           respUserData.StafID,
					OrganizationName: v.Organization,
					PeriodFrom:       periodFrom,
					PeriodTo:         periodTo,
					Position:         v.Position,
					JobDesc:          html.EscapeString(v.JobDesc),
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error update table user work experience", zap.Error(err))
					return err
				}
			} else {
				// store table user_work_experience
				err = c.repoCandidate.StoreUserWorkExperienceWithTx(ctx, tx, model.UserWorkExperience{
					IDUser:           respUserData.StafID,
					OrganizationName: v.Organization,
					PeriodFrom:       periodFrom,
					PeriodTo:         periodTo,
					Position:         v.Position,
					JobDesc:          html.EscapeString(v.JobDesc),
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error store table user work experience", zap.Error(err))
					return err
				}

			}
		}
	}

	// store table user_contact
	if payload.Contact != "" {
		err := json.Unmarshal([]byte(payload.Contact), &payloadContact)
		if err != nil {
			tx.Rollback()
			cLogger.Error("error payload contacts", zap.Error(err))
			return errors.New("invalid_format_contact_payload")
		}

		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}

		respUserContact, err := c.repoCandidate.FindUserContact(ctx, model.UserContact{
			IDUser:         respUserData.StafID,
			ContactCaption: payloadContact.ContactNumber,
			ContactLabel:   "mobile",
			ContactType:    payloadContact.Type,
		})
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user contact ", zap.Error(err))
			return err
		}

		if respUserContact == nil {
			// store table user_contact
			err = c.repoCandidate.StoreUserContactWithTx(ctx, tx, model.UserContact{
				IDUser:         respUserData.StafID,
				ContactCaption: payloadContact.ContactNumber,
				ContactLabel:   "mobile",
				ContactType:    payloadContact.Type,
			})
			if err != nil {
				tx.Rollback()
				cLogger.Error("error store table user contact", zap.Error(err))
				return err
			}
		}

		// store user contact email
		respUserContact, err = c.repoCandidate.FindUserContact(ctx, model.UserContact{
			IDUser:         respUserData.StafID,
			ContactCaption: payload.Email,
			ContactLabel:   "other",
			ContactType:    "email",
		})
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user contact ", zap.Error(err))
			return err
		}

		if respUserContact == nil {
			// store table user_contact
			err = c.repoCandidate.StoreUserContactWithTx(ctx, tx, model.UserContact{
				IDUser:         respUserData.StafID,
				ContactCaption: payload.Email,
				ContactLabel:   "other",
				ContactType:    "email",
			})
			if err != nil {
				tx.Rollback()
				cLogger.Error("error store table user contact", zap.Error(err))
				return err
			}
		}

		if payload.ReferenceLinks != "" {
			err := json.Unmarshal([]byte(payload.ReferenceLinks), &payloadReferenceLinks)
			if err != nil {
				tx.Rollback()
				cLogger.Error("error payload reference links", zap.Error(err))
				return errors.New("invalid_format_reference_links_payload")
			}
			// store user contact reference link
			// store table user_contact
			for _, v := range payloadReferenceLinks {
				err = c.repoCandidate.StoreUserContactWithTx(ctx, tx, model.UserContact{
					IDUser:         respUserData.StafID,
					ContactCaption: v.Link,
					ContactLabel:   "other",
					ContactType:    "other",
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error store table user contact- reference link", zap.Error(err))
					return err
				}
			}
		}

	}

	// store user job candidate portal
	err = c.doInsertUserJobCandidatePortalWithTx(ctx, tx, payload)
	if err != nil {
		tx.Rollback()
		cLogger.Error("error store table user data", zap.Error(err))
		return err
	}

	// store user job candidate portal skill
	if payload.Skills != "" {
		err := json.Unmarshal([]byte(payload.Skills), &payloadSkills)
		if err != nil {
			tx.Rollback()
			cLogger.Error("error payload skills", zap.Error(err))
			return errors.New("invalid_format_skills_payload")
		}

		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			cLogger.Error("error find user ", zap.Error(err))
			return err
		}

		for _, v := range payloadSkills {
			_, err = c.repoCandidate.FindUserSkillWithTx(ctx, tx, model.UserJobCandidatePortalSkill{
				StafID:    respUserData.StafID,
				JCPSSkill: v,
			})
			if err != nil && err.Error() != "notFound" {
				tx.Rollback()
				cLogger.Error("error find job candidate portal skill", zap.Error(err))
				return err
			}

			// insert if skill is not found
			if err != nil && err.Error() == "notFound" {
				err = c.repoCandidate.StoreUserJobCandidatePortalSkillWithTx(ctx, tx, model.UserJobCandidatePortalSkill{
					StafID:    respUserData.StafID,
					JCPSSkill: v,
				})
				if err != nil {
					tx.Rollback()
					cLogger.Error("error store table job candidate portal skill", zap.Error(err))
					return err
				}
			}

		}

	}

	// store file CV
	if payload.Cv != nil {
		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			return err
		}

		respUserFile, err := c.repoCandidate.FindUserFileWithTx(ctx, tx, model.UserFile{
			IDUser: respUserData.StafID,
		})
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			return err
		}

		if respUserFile == nil {
			ext := filepath.Ext(payload.Cv.Filename)
			fileName := fmt.Sprintf("%d_cv%s", respUserData.StafID, ext)
			fileName = strings.Replace(fileName, " ", "-", -1)

			err = UploadFile(cfg, fileName, payload.Cv)
			if err != nil {
				tx.Rollback()
				return err
			}

			// store user file
			err = c.repoCandidate.StoreUserFileWithTx(ctx, tx, model.UserFile{
				IDUser:         respUserData.StafID,
				IDFileCategory: 24,
				FileAttachment: fileName,
			})
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			ext := filepath.Ext(payload.Cv.Filename)
			fileName := fmt.Sprintf("%d_cv%s", respUserData.StafID, ext)
			fileName = strings.Replace(fileName, " ", "-", -1)
			err = UploadFile(cfg, fileName, payload.Cv)
			if err != nil {
				tx.Rollback()
				return err
			}

			// update user file
			err = c.repoCandidate.UpdateUserFileWithTx(ctx, tx, model.UserFile{
				IDUser:         respUserData.StafID,
				FileAttachment: fileName,
			})
			if err != nil {
				tx.Rollback()
				return err
			}
		}

	}

	// update user photo
	if payload.Photo != nil {
		// find user
		respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
		if err != nil && err.Error() != "notFound" {
			tx.Rollback()
			return err
		}

		fileName := fmt.Sprintf("%d_profile_photo_%s", respUserData.StafID, payload.Photo.Filename)
		fileName = strings.Replace(fileName, " ", "-", -1)

		err = UploadFile(cfg, fileName, payload.Photo)
		if err != nil {
			tx.Rollback()
			return err
		}

		// update user photo
		err = c.repoCandidate.UpdateUserPhotoWithTx(ctx, tx, model.UserData{
			StafID:    respUserData.StafID,
			UserPhoto: fileName,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err == nil {
		err := tx.Commit()
		if err != nil {
			cLogger.Error("error commit transaction", zap.Error(err))
			return err
		}
	}

	cLogger.Info("success store candidate")
	return nil
}

func (c *candidateService) ValidateDataUser(ctx context.Context, stafID int, payload StoreCandidatePayload) *model.UserData {
	var dataUser model.UserData
	// find user data
	respUserData, err := c.repoCandidate.FindOneUserDataByStaffID(ctx, stafID)
	if err != nil {
		return nil
	}

	firstName, middleName, lastName := util.SplitFullname(payload.Fullname)

	if firstName == "" {
		dataUser.FirstName = respUserData.FirstName
	} else {
		dataUser.FirstName = firstName
	}

	if middleName == "" {
		dataUser.MiddleName = respUserData.MiddleName
	} else {
		dataUser.MiddleName = middleName
	}

	if lastName == "" {
		dataUser.LastName = respUserData.LastName
	} else {
		dataUser.LastName = lastName
	}

	if payload.Nickname == "" {
		dataUser.Nickname = respUserData.Nickname
	} else {
		dataUser.Nickname = payload.Nickname
	}

	if payload.Gender == "" {
		dataUser.Gender = respUserData.Gender
	} else {
		// find gender
		respGender, _ := c.repoCandidate.FindGender(ctx, payload.Gender)
		dataUser.Gender = respGender.ID
	}

	if payload.DateOfBirth == "" {
		dataUser.Birthdate = respUserData.Birthdate
	} else {
		dataUser.Birthdate = sql.NullString{
			String: payload.DateOfBirth,
			Valid:  true,
		}
	}

	if payload.Location == "" {
		dataUser.CurrentProvince = respUserData.CurrentProvince
		dataUser.CurrentRegency = respUserData.CurrentRegency
	} else {
		// handle location
		var (
			IDProv    string
			IDRegency string
		)
		loc := strings.Split(payload.Location, ",")
		if len(loc) == 2 {
			prov := loc[1]
			reg := loc[0]

			// convert province and regency
			prov = ConvertProvince(prov)
			reg = ConvertRegency(reg)

			// find province
			masterProvices, err := c.repoCandidate.FindMasterProvince(ctx, prov)
			if err == nil {
				IDProv = masterProvices.ID
			}

			// find regency
			masterRegency, err := c.repoCandidate.FindMasterRegency(ctx, reg)
			if err == nil {
				IDRegency = masterRegency.ID
			}
		}
		dataUser.CurrentProvince = IDProv
		dataUser.CurrentRegency = IDRegency
	}

	return &dataUser
}

func (c *candidateService) doInsertUserDataWithTx(ctx context.Context, tx *sql.Tx, payload StoreCandidatePayload) error {
	var (
		genderID  int
		IDProv    string
		IDRegency string
	)
	// find user
	respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
	if err != nil && err.Error() != "notFound" {
		return err
	}

	if payload.Gender != "" {
		// find gender
		respGender, err := c.repoCandidate.FindGender(ctx, payload.Gender)
		if err != nil && err.Error() != "notFound" {
			return err
		}
		genderID = respGender.ID
	}

	if payload.Gender == "" || payload.Gender == "-" {
		genderID = 0
	}

	// store table user data
	firstName, middleName, lastName := util.SplitFullname(payload.Fullname)

	var dateOfBirth sql.NullTime
	if payload.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", payload.DateOfBirth)
		if err != nil {
			dateOfBirth = sql.NullTime{Time: time.Time{}, Valid: false}
		}
		dateOfBirth = sql.NullTime{Time: dob, Valid: true}
	}
	if payload.DateOfBirth == "" {
		dateOfBirth = sql.NullTime{Time: time.Time{}, Valid: false}
	}

	// handle location
	loc := strings.Split(payload.Location, ",")
	if len(loc) == 2 {
		prov := loc[1]
		reg := loc[0]

		// convert province and regency
		prov = ConvertProvince(prov)
		reg = ConvertRegency(reg)

		// find province
		masterProvices, err := c.repoCandidate.FindMasterProvince(ctx, prov)
		if err == nil {
			IDProv = masterProvices.ID
		}

		// find regency
		masterRegency, err := c.repoCandidate.FindMasterRegency(ctx, reg)
		if err == nil {
			IDRegency = masterRegency.ID
		}
	}

	err = c.repoCandidate.StoreUserDataWithTx(ctx, tx, model.Candidate{
		StafID:          respUserData.StafID,
		FirstName:       firstName,
		MiddleName:      middleName,
		LastName:        lastName,
		Nickname:        payload.Nickname,
		Gender:          genderID,
		DateOfBirth:     dateOfBirth,
		CurrentProvince: IDProv,
		CurrentRegency:  IDRegency,
		Channel:         payload.Channel,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *candidateService) doInsertUserJobCandidatePortalWithTx(ctx context.Context, tx *sql.Tx, payload StoreCandidatePayload) error {
	// find user
	respUserData, err := c.repoCandidate.FindOneUserByEmailWithTx(ctx, tx, payload.Email)
	if err != nil && err.Error() != "notFound" {
		return err
	}

	var salary int
	var salaryUnit string
	switch payload.Channel {
	case "pintarnya":
		x, _ := strconv.Atoi(payload.LatestSalary)
		salary = x
		salaryUnit = "latest"
	case "jooble":
		salary = 0
		salaryUnit = ""
	default:
		x, _ := strconv.Atoi(payload.SalaryExpectation)
		salary = x
		salaryUnit = "expectation"
	}

	dateApply, _ := time.Parse("2006-01-02", payload.AppliedDate)

	err = c.repoCandidate.StoreUserJobCandidatePortalWithTx(ctx, tx, model.UserJobCandidatePortal{
		StafID:        respUserData.StafID,
		UUID:          payload.AppliedForId,
		JCPName:       payload.AppliedFor,
		JCPURLProfile: payload.ProfileLink,
		JCPSummary:    html.EscapeString(payload.Summary),
		JCPSalary:     salary,
		JCPSalaryUnit: salaryUnit,
		JCPFrom:       payload.Channel,
		JCPDateApply:  dateApply,
	})
	if err != nil {
		return err
	}
	return nil
}
