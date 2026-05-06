package handler_candidate

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/service/service_candidate"
	"helicopter-hr/pkg/ginx"
	"helicopter-hr/pkg/validatorx"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const (
	errorOnlyFilePDFImageAllowed = "Only PDF, Doc, or Images files are allowed"
	errorOnlyFileImageAllowed    = "Only Images files are allowed"
)

var (
	extAllowedCV = []string{".jpg", ".jpeg", ".jpe", ".jif", ".jfif", ".jfi",
		".png", ".gif", ".webp", ".tiff", ".tif", ".bmp", ".dib", ".psd", ".heif", ".heic",
		".ind", ".indd", ".indt", ".jp2", ".j2k", ".jpf", ".jpx", ".jpm", ".jpx", ".mj2",
		".svg", ".svgz", ".eps", ".ai", ".pdf", ".doc", ".docx"}

	extAllowedPhoto = []string{".jpg", ".jpeg", ".jpe", ".jif", ".jfif", ".jfi",
		".png", ".gif", ".webp", ".tiff", ".tif", ".bmp", ".dib", ".psd", ".heif", ".heic",
		".ind", ".indd", ".indt", ".jp2", ".j2k", ".jpf", ".jpx", ".jpm", ".jpx", ".mj2",
		".svg", ".svgz", ".eps", ".ai"}
)

func (h *Handler) HandlerStoreCandidate(ctx *gin.Context) {
	var (
		guid    = ctx.Value("request_id").(string)
		payload service_candidate.StoreCandidatePayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.candidate.store"),
		zap.String("request_id", guid),
	)

	if err := ctx.ShouldBind(&payload); err != nil {
		cLogger.Error("error decode payload store candidate", zap.Error(err))
		ginx.RespondWithError(
			ctx,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err,
		)
		return
	}

	// validation rules
	violations := h.validationStoreCandidate(payload)
	if len(violations) > 0 {
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), violations)
		return
	}

	// validate file cv
	violations = h.validationFileCV(payload.Cv)
	if len(violations) > 0 {
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), violations)
		return
	}

	// validate file photo
	violations = h.validationFilePhoto(payload.Photo)
	if len(violations) > 0 {
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), violations)
		return
	}

	err := h.candidateService.StoreCandidate(ctx, payload)
	if err != nil {
		switch err.Error() {
		case "notFound":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
				nil,
			)
			return
		case "invalid_format_education_payload":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "invalid_format_work_experiences_payload":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "invalid_format_contact_payload":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "invalid_format_skills_payload":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "invalid_format_reference_links_payload":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "candidate_already_applied":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		case "invalid_code_education":
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnprocessableEntity,
				err.Error(),
				nil,
			)
			return
		default:
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				err.Error(),
				nil,
			)
			return
		}

	}

	cLogger.Info("success")
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", nil)
	return
}

func (h *Handler) validationStoreCandidate(payload service_candidate.StoreCandidatePayload) []ginx.ErrorField {
	validation := validatorx.New()

	if payload.Type == "applicant" {
		return validation.Set("email", validatorx.Required().Validate(payload.Email)).
			Set("email", validatorx.ValidEmail().Validate(payload.Email)).Apply()
	}

	return validation.Set("channel", validatorx.Required().Validate(payload.Channel)).
		Set("channel", validatorx.ValidIn([]string{"kita_lulus", "pintarnya", "seek", "glints", "jooble"}).Validate(payload.Channel)).
		Set("type", validatorx.Required().Validate(payload.Type)).
		Set("type", validatorx.ValidIn([]string{"applicant", "recomendation"}).Validate(payload.Type)).
		Set("applied_for", validatorx.Required().Validate(payload.AppliedFor)).
		Set("fullname", validatorx.Required().Validate(payload.Fullname)).
		Set("fullname", validatorx.ValidLength(3, 255).Validate(payload.Fullname)).
		Set("gender", validatorx.ValidIn([]string{"MALE", "FEMALE"}).Validate(payload.Gender)).
		Set("contact", validatorx.Required().Validate(payload.Contact)).
		Apply()

}

func (h *Handler) validationFileCV(header *multipart.FileHeader) []ginx.ErrorField {
	var respondErrorFields []ginx.ErrorField

	if header == nil {
		return respondErrorFields
	}

	// Validate file extension
	ext := filepath.Ext(header.Filename)
	isAllowedExt := false
	for _, v := range extAllowedCV {
		if v == ext {
			isAllowedExt = true
		}
	}
	if !isAllowedExt {
		respondErrorFields = append(respondErrorFields, ginx.ErrorField{
			Field: "cv",
			Error: errorOnlyFilePDFImageAllowed,
		})
	}

	if respondErrorFields != nil {
		return respondErrorFields
	}

	return nil
}

func (h *Handler) validationFilePhoto(header *multipart.FileHeader) []ginx.ErrorField {
	var respondErrorFields []ginx.ErrorField

	if header == nil {
		return respondErrorFields
	}

	// Validate file extension
	ext := filepath.Ext(header.Filename)
	isAllowedExt := false
	for _, v := range extAllowedPhoto {
		if v == ext {
			isAllowedExt = true
		}
	}
	if !isAllowedExt {
		respondErrorFields = append(respondErrorFields, ginx.ErrorField{
			Field: "photo",
			Error: errorOnlyFileImageAllowed,
		})
	}

	if respondErrorFields != nil {
		return respondErrorFields
	}

	return nil
}
