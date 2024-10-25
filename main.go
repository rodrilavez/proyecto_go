package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"sistema_escolar/database"
	"sistema_escolar/handlers"
)

var db *sql.DB

func main() {
	var err error
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.Default())

	// Rutas para Calificaciones

	router.POST("/api/grades", handlers.CreateGrade)                              // Crear Calificación
	router.PUT("/api/grades/:grade_id", handlers.UpdateGrade)                     // Actualizar Calificación por ID
	router.DELETE("/api/grades/:grade_id", handlers.DeleteGrade)                  // Eliminar Calificación por ID
	router.GET("/api/grades/:grade_id", handlers.GetGrade)                        // Obtener Calificación de una materia de un estudiante
	router.GET("/api/grades/student/:student_id", handlers.GetAllGradesByStudent) // Obtener todas las calificaciones de un estudiante

	router.Run(":8080")
}
