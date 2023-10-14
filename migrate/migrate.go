package main

import (
	"github.com/ricoaditya-u/hris_dev/db"
	"github.com/ricoaditya-u/hris_dev/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(
		// &models.Grade{},
		// &models.Level{},
		// &models.JobDescription{},
		// &models.Division{},
		// &models.Department{},
		// &models.Supervision{},
		// &models.Ptkp{},
		// &models.Employee{},
		// &models.Candidate{},
		// &models.Role{},
		// &models.User{},
		// &models.Family{},
		// &models.Education{},
		// &models.Experience{},
		// &models.HealthDisease{},
		// &models.CriminalNote{},
		// &models.Course{},
		// &models.Reference{},
		// &models.SelfPerformance{},
		// &models.EvaluationForm{},
		// &models.EvaluationPoint{},
		// &models.Evaluation{},
		// &models.EvaluationResult{},
		// &models.SalarySlip{},
		// &models.SalarySlipDetail{},
		// &models.Pph{},
		// &models.Mpp{},
		&models.Casbin_rule{},
	)
}
