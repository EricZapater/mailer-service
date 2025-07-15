package server

import (
	"mailer-service/config"
	"mailer-service/internal/api"
	"mailer-service/internal/mailer"
	"mailer-service/internal/middleware"
	"mailer-service/internal/service"

	_ "mailer-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Mailer Service API
// @version         1.0
// @description     API per enviar emails amb templates i attachments
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

type Server struct {
	router *gin.Engine	
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	router := gin.Default()
	return &Server{
		router: router,
		config: cfg,
	}
}

func (s *Server) Start() error {
	s.router.Use(middleware.SetupCORS())
	templateService := service.NewTemplateService()
	mailerService := mailer.NewMailer(templateService, s.config)	
	handler := api.NewHandler(mailerService)	
	
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public := s.router.Group("/api")
	public.POST("/send", handler.SendEmail)

	return s.router.Run(":" + s.config.APIPort)
}