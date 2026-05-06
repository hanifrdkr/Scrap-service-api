package middleware

import (
	"bytes"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func CustomGinzap(logger *zap.Logger) gin.HandlerFunc {
	return ginzap.GinzapWithConfig(logger.With(zap.String("component", "gin")), &ginzap.Config{
		//TimeFormat: logger.LogTimeFmt,
		UTC: false,
		Context: func(c *gin.Context) []zapcore.Field {
			var (
				fields []zapcore.Field
				guid   = c.Value("request_id").(string)
			)

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)

			fields = append(fields, zap.String("body", string(body)))
			fields = append(fields, zap.String("request_id", guid))

			return fields
		}})
}
