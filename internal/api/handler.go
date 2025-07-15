package api

import (
	"bytes"
	"fmt"
	"io"
	"mailer-service/internal/mailer"
	"mailer-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	mailer *mailer.Mailer	
}

func NewHandler(mailer *mailer.Mailer) *Handler {
	return &Handler{
		mailer: mailer,
	}
}

// SendEmail envia un email amb template i attachments
// @Summary      Envia un email
// @Description  Envia un email utilitzant un template amb possibles attachments
// @Tags         emails
// @Accept       multipart/form-data
// @Produce      json
// @Param        mail_to        formData  string  true   "Email destinatari"
// @Param        subject        formData  string  true   "Assumpte del email"
// @Param        body           formData  string  false  "Cos del email"
// @Param        template_name  formData  string  false  "Nom del template a utilitzar"
// @Param        footer         formData  string  false  "Peu del email"
// @Param        mail_cc        formData  string  false  "Email en còpia"
// @Param        attachments    formData  file    false  "Fitxers adjunts (pot haver-n'hi més d'un)"
// @Success      200            {object}  map[string]string
// @Failure      400            {object}  map[string]string
// @Failure      500            {object}  map[string]string
// @Router       /send [post]
func (h *Handler) SendEmail(c *gin.Context) {
	var req model.Request

	// Camps de text
	req.MailTo = c.PostForm("mail_to")
	req.Body = c.PostForm("body")
	req.TemplateName = c.PostForm("template_name")
	req.Subject = c.PostForm("subject")
	req.Footer = c.PostForm("footer")

	if cc := c.PostForm("mail_cc"); cc != "" {
		req.MailCc = &cc
	}

	// Fitxers adjunts (pot haver-n'hi més d'un)
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	files := form.File["attachments"]
	if len(files) > 0 {
		var attachments []model.Attachment

		for _, fh := range files {
			file, err := fh.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading attachment"})
				return
			}
			defer file.Close()

			var buf bytes.Buffer
			if _, err := io.Copy(&buf, file); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
				return
			}
			mimetype := http.DetectContentType(buf.Bytes())
			attachments = append(attachments, model.Attachment{
				Filename: fh.Filename,
				Content:  buf.Bytes(),
				ContentType: mimetype,
			})
		}

		req.Attachments = &attachments
	}

	// Trucar el mailer
	if err := h.mailer.SendEmail(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send email: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

