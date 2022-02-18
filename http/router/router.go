package router

import (
	"github.com/gin-gonic/gin"

	// dh "epiket-api/internal/http/handler/dashboard"
	gh "sertifikasi_listrik/http/handler/global"
	level_h "sertifikasi_listrik/http/handler/level"
	pelanggan_h "sertifikasi_listrik/http/handler/pelanggan"
	pembayaran_h "sertifikasi_listrik/http/handler/pembayaran"
	penggunaan_h "sertifikasi_listrik/http/handler/penggunaan"
	tagihan_h "sertifikasi_listrik/http/handler/tagihan"
	tarif_h "sertifikasi_listrik/http/handler/tarif"
	user_h "sertifikasi_listrik/http/handler/user"
	"sertifikasi_listrik/http/middleware/auth"
	"sertifikasi_listrik/http/middleware/cors"
)

// Routes ...
func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Middleware())

	// for health check purpose only
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.MaxMultipartMemory = 8 << 20 // 8 MiB

	globalHandler := gh.NewHandler()
	levelHendler := level_h.NewHandler()
	pelangganHendler := pelanggan_h.NewHandler()
	pembayaranHandler := pembayaran_h.NewHandler()
	penggunaanHandler := penggunaan_h.NewHandler()
	tagihanHandler := tagihan_h.NewHandler()
	tarifHandler := tarif_h.NewHandler()
	userHandler := user_h.NewHandler()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", globalHandler.Login)
		v1.POST("/register", globalHandler.Register)
		// v1.GET("/test-file", func(c *gin.Context) {
		// 	log.Println("oke")
		// 	c.FileAttachment(fmt.Sprintf("%s/report.pdf", os.Getenv("EXP_PDF_PATH")), "report.pdf")
		// })

		resources := v1.Group("/resources").Use(auth.Middleware())
		{
			// ----------------------------------PASSED CHECK!!!

			resources.POST("/level", levelHendler.Create)
			resources.GET("/level/:id", levelHendler.GetOneByID)
			resources.PUT("/level/:id", levelHendler.UpdateOneByID)
			resources.DELETE("/level/:id", levelHendler.DeleteOneByID)
			resources.GET("/level", levelHendler.GetAll)

			// --------------------------------------------------

			resources.POST("/pelanggan", pelangganHendler.Create)
			resources.GET("/pelanggan/:id", pelangganHendler.GetOneByID)
			resources.PUT("/pelanggan/:id", pelangganHendler.UpdateOneByID)
			resources.DELETE("/pelanggan/:id", pelangganHendler.DeleteOneByID)
			resources.GET("/pelanggan", pelangganHendler.GetAll)

			resources.POST("/pembayaran/:id/:idt", pembayaranHandler.Create)
			resources.GET("/pembayaran/:id", pembayaranHandler.GetOneByID)
			resources.PUT("/pembayaran/:id", pembayaranHandler.UpdateOneByID)
			resources.DELETE("/pembayaran/:id", pembayaranHandler.DeleteOneByID)
			resources.GET("/pembayaran", pembayaranHandler.GetAll)

			resources.POST("/penggunaan", penggunaanHandler.Create)
			resources.GET("/penggunaan/:id", penggunaanHandler.GetOneByID)
			resources.PUT("/penggunaan/:id", penggunaanHandler.UpdateOneByID)
			resources.PUT("/penggunaan_status/:id", penggunaanHandler.UpdateStatus)
			resources.DELETE("/penggunaan/:id", penggunaanHandler.DeleteOneByID)
			resources.GET("/penggunaan", penggunaanHandler.GetAll)

			resources.POST("/tagihan", tagihanHandler.Create)
			resources.GET("/tagihan/:id", tagihanHandler.GetOneByID)
			resources.PUT("/tagihan/:id", tagihanHandler.UpdateOneByID)
			resources.PUT("/tagihan_status/:id", tagihanHandler.UpdateStatus)
			resources.DELETE("/tagihan/:id", tagihanHandler.DeleteOneByID)
			resources.GET("/tagihan", tagihanHandler.GetAll)

			resources.POST("/tarif", tarifHandler.Create)
			resources.GET("/tarif/:id", tarifHandler.GetOneByID)
			resources.PUT("/tarif/:id", tarifHandler.UpdateOneByID)
			resources.DELETE("/tarif/:id", tarifHandler.DeleteOneByID)
			resources.GET("/tarif", tarifHandler.GetAll)

			resources.POST("/user", userHandler.Create)
			resources.GET("/user/:id", userHandler.GetOneByID)
			resources.PUT("/user/:id", userHandler.UpdateOneByID)
			resources.DELETE("/user/:id", userHandler.DeleteOneByID)
			resources.GET("/user", userHandler.GetAll)

			// 	resources.POST("/user", userHandler.Create)
			// 	resources.POST("/user/:id", userHandler.ResetPasswordByID)

			// 	resources.GET("/instansi-parent/:id", instansiHandler.GetOneByParentID)

			// 	// resources.GET("/user/:id", userHandler.GetOneByID)
			// 	// resources.PUT("/user/:id", userHandler.UpdateOneByID)
			// 	// resources.DELETE("/user/:id", userHandler.DeleteOneByID)
			// 	// resources.GET("/user", userHandler.GetAll)

			// 	resources.POST("/kuesioner", kuesionerHandler.Create)
			// 	resources.GET("/kuesioner/:id", kuesionerHandler.GetOneByID)
			// 	resources.PUT("/kuesioner/:id", kuesionerHandler.UpdateOneByID)
			// 	resources.DELETE("/kuesioner/:id", kuesionerHandler.DeleteOneByID)
			// 	resources.GET("/kuesioner", kuesionerHandler.GetAll)
			// 	resources.GET("/kuesioner/:id/poin-kuesioner/:id_poin_kuesioner", kuesionerHandler.GetOneByIDWithDetails)

			// 	resources.POST("/poin-kuesioner", poinKuesionerHandler.Create)
			// 	resources.GET("/poin-kuesioner/:id", poinKuesionerHandler.GetOneByID)
			// 	resources.PUT("/poin-kuesioner/:id", poinKuesionerHandler.UpdateOneByID)
			// 	resources.DELETE("/poin-kuesioner/:id", poinKuesionerHandler.DeleteOneByID)
			// 	// resources.GET("/poin-kuesioner", poinKuesionerHandler.GetAll)
			// 	resources.PUT("/poin-kuesioner-with-item/:id", poinKuesionerHandler.UpdateOneWithItemByID)

			// 	resources.POST("/poin-kuesioner-item", poinKuesionerItemHandler.Create)
			// 	resources.GET("/poin-kuesioner-item/:id", poinKuesionerItemHandler.GetOneByID)
			// 	resources.PUT("/poin-kuesioner-item/:id", poinKuesionerItemHandler.UpdateOneByID)
			// 	resources.DELETE("/poin-kuesioner-item/:id", poinKuesionerItemHandler.DeleteOneByID)
			// 	resources.GET("/poin-kuesioner-item", poinKuesionerItemHandler.GetAll)

			// 	resources.POST("/survey", surveyHandler.Create)
			// 	resources.GET("/survey/:id", surveyHandler.GetOneByID)
			// 	resources.PUT("/survey/:id", surveyHandler.UpdateOneByID)
			// 	resources.DELETE("/survey/:id", surveyHandler.DeleteOneByID)
			// 	resources.GET("/survey", surveyHandler.GetAll)
			// 	resources.GET("/survey-history", surveyHandler.GetHistoryByID)

			// 	resources.POST("/survey-item", surveyItemHandler.Create)
			// 	resources.GET("/survey-item/:id", surveyItemHandler.GetOneByID)
			// 	resources.PUT("/survey-item/:id", surveyItemHandler.UpdateOneByID)
			// 	resources.DELETE("/survey-item/:id", surveyItemHandler.DeleteOneByID)
			// 	resources.GET("/survey-item", surveyItemHandler.GetAll)

			// 	resources.POST("/survey-assignment", surveyAssignmentHandler.Create)
			// 	// resources.PUT("/survey-assignment/:id", surveyAssignmentHandler.UpdateOneByID)
			// 	resources.PUT("/survey-assignment-header/:id", surveyAssignmentHandler.UpdateOneHeaderByID)
			// 	resources.GET("/survey-assignment/:id", surveyAssignmentHandler.GetOneByID)
			// 	resources.GET("/survey-assignment", surveyAssignmentHandler.GetAll)
			// 	// resources.DELETE("/survey-assignment/:id", surveyAssignmentHandler.DeleteOneByID)

			// 	resources.POST("/survey-assignment-item", surveyAssignmentPegawaiInstansiHandler.Create)
			// 	resources.GET("/survey-assignment-item/:id", surveyAssignmentPegawaiInstansiHandler.GetOneByID)
			// 	resources.GET("/survey-assignment-item", surveyAssignmentPegawaiInstansiHandler.GetAll)

			// 	resources.POST("/survey-assignment-instansi", surveyAssignmentInstansiHandler.Create)
			// 	resources.GET("/survey-assignment-instansi/:id", surveyAssignmentInstansiHandler.GetOneByID)
			// 	resources.GET("/survey-assignment-instansi", surveyAssignmentInstansiHandler.GetAll)

			// 	resources.GET("/dashboard-chartbar", dashboardHandler.GetAllDataChartBarDashboard)
			// 	resources.GET("/dashboard-chartbar-table", dashboardHandler.GetAllDataChartTableDashboard)
			// 	resources.GET("/dashboard-chartbar-active-directorate", dashboardHandler.GetAllDataChartBarActiveDirectorate)
			// 	resources.GET("/dashboard-chartbar-pivot-bymonth", dashboardHandler.GetAllDataChartPivotmonth)
			// 	resources.GET("/dashboard-chartbar-eachassignment-perdirectorate", dashboardHandler.GetAllDataChartBarEachAssignmentPerDirektorat)
			// 	resources.GET("/dashboard-table-average-value-eachquestion-allassignment-perdirectorate", dashboardHandler.GetAllDataAverageValueEachQuestionFromAllAssignmentPerDirektorat)
			// 	resources.GET("/dashboard-table-average-value-allassignment-perdirectorate", dashboardHandler.GetAllDataChartIndexAverageAllAssignmentPerDirektorat)

			// 	// dibawah ini adalah router khusus untuk aplikasi responden
			// 	// dashboard
			// 	resources.GET("/respondent/survey", respondentHandler.GetAllSurveyAssigned)
			// 	resources.GET("/respondent/survey/history", surveyHandler.GetHistoryByID)

			// 	// // start survey
			// 	// resources.GET("/respondent/survey/assigned/:id", respondentHandler.GetDetailSurveyAssignedByIDSurvey)
			// 	resources.GET("/temporary-uri/:id_survey/details/:id_survey_assignment", respondentHandler.GetDetailSurveyAssignedByIDSurvey)
			// 	resources.POST("/temporary-start-survey", respondentHandler.StartNewSurvey)

			// 	// // on going survey
			// 	resources.GET("/temporary-uri-survey/:id_survey_assignment", respondentHandler.GetOnePageSurveyAssignedByIDSurvey)
			// 	resources.POST("/temporary-submit", respondentHandler.Submit)
			// 	resources.POST("/temporary-upload", respondentHandler.UploadFile)

			// 	// end survey
		}
	}

	return r
}
