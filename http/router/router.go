package router

import (
	"github.com/gin-gonic/gin"

	// dh "epiket-api/internal/http/handler/dashboard"
	gh "epiket-api/http/handler/global"
	kh "epiket-api/http/handler/kegiatan"
	rh "epiket-api/http/handler/role"
	seh "epiket-api/http/handler/seksi"
	suh "epiket-api/http/handler/subdirektorat"

	// ph "epiket-api/internal/http/handler/pegawai"
	// pkh "epiket-api/internal/http/handler/poinkuesioner"
	// pkih "epiket-api/internal/http/handler/poinkuesioneritem"
	// rpt "epiket-api/internal/http/handler/report"
	// rh "epiket-api/internal/http/handler/respondent"
	// ssh "epiket-api/internal/http/handler/struktursurvey"
	// sh "epiket-api/internal/http/handler/survey"
	// sah "epiket-api/internal/http/handler/surveyassignment"
	// saih "epiket-api/internal/http/handler/surveyassignmentinstansi"
	// sapih "epiket-api/internal/http/handler/surveyassignmentpegawaiinstansi"
	// sih "epiket-api/internal/http/handler/surveyitem"
	// tph "epiket-api/internal/http/handler/tipepoin"
	// uh "epiket-api/internal/http/handler/user"
	"epiket-api/http/middleware/auth"
	"epiket-api/http/middleware/cors"
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

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	globalHandler := gh.NewHandler()
	seksiHandler := seh.NewHandler()
	subdirektoratHandler := suh.NewHandler()
	kegiatanHandler := kh.NewHandler()
	roleHandler := rh.NewHandler()
	// poinKuesionerHandler := pkh.NewHandler()
	// poinKuesionerItemHandler := pkih.NewHandler()
	// kuesionerHandler := kh.NewHandler()
	// instansiHandler := ih.NewHandler()
	// userHandler := uh.NewHandler()
	// surveyHandler := sh.NewHandler()
	// surveyItemHandler := sih.NewHandler()
	// surveyAssignmentHandler := sah.NewHandler()
	// surveyAssignmentPegawaiInstansiHandler := sapih.NewHandler()
	// surveyAssignmentInstansiHandler := saih.NewHandler()
	// dashboardHandler := dh.NewHandler()
	// respondentHandler := rh.NewHandler()
	// strukturSurveyHandler := ssh.NewHandler()
	//repot
	// reportHandler := rpt.NewHandler()

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

			resources.POST("/seksi", seksiHandler.Create)
			resources.GET("/seksi/:id", seksiHandler.GetOneByID)
			resources.PUT("/seksi/:id", seksiHandler.UpdateOneByID)
			resources.DELETE("/seksi/:id", seksiHandler.DeleteOneByID)
			resources.GET("/seksi", seksiHandler.GetAll)

			// --------------------------------------------------

			resources.POST("/subdirektorat", subdirektoratHandler.Create)
			resources.GET("/subdirektorat/:id", subdirektoratHandler.GetOneByID)
			resources.PUT("/subdirektorat/:id", subdirektoratHandler.UpdateOneByID)
			resources.DELETE("/subdirektorat/:id", subdirektoratHandler.DeleteOneByID)
			resources.GET("/subdirektorat", subdirektoratHandler.GetAll)

			resources.POST("/kegiatan", kegiatanHandler.Create)
			resources.GET("/kegiatan/:id", kegiatanHandler.GetOneByID)
			resources.PUT("/kegiatan/:id", kegiatanHandler.UpdateOneByID)
			resources.DELETE("/kegiatan/:id", kegiatanHandler.DeleteOneByID)
			resources.GET("/kegiatan", kegiatanHandler.GetAll)

			resources.POST("/role", roleHandler.Create)
			resources.GET("/role/:id", roleHandler.GetOneByID)
			resources.PUT("/role/:id", roleHandler.UpdateOneByID)
			resources.DELETE("/role/:id", roleHandler.DeleteOneByID)
			resources.GET("/role", roleHandler.GetAll)

			// 	resources.POST("/jabatan", jabatanHandler.Create)
			// 	resources.GET("/jabatan/:id", jabatanHandler.GetOneByID)
			// 	resources.PUT("/jabatan/:id", jabatanHandler.UpdateOneByID)
			// 	resources.DELETE("/jabatan/:id", jabatanHandler.DeleteOneByID)
			// 	resources.GET("/jabatan", jabatanHandler.GetAll)

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
