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

	// Rutas para Estudiantes
	router.POST("/api/students", handlers.CreateStudent)               //Crear Estudiante
	router.DELETE("/api/students/:student_id", handlers.DeleteStudent) // Eliminar Estudiante
	router.PUT("/api/students/:student_id", handlers.UpdateStudent)    // Actualizar Estudiante
	router.GET("/api/students", handlers.GetAllStudents)               // Obtener Todos Los Estudiantes
	router.GET("/api/students/:student_id", handlers.GetStudent)       //Obtener Estudiante

	// Rutas para Materias
	router.POST("/api/subjects", handlers.CreateSubject)               // Crear una nueva materia
	router.PUT("/api/subjects/:subject_id", handlers.UpdateSubject)    // Actualizar una materia por ID
	router.GET("/api/subjects/:subject_id", handlers.GetSubject)       // Obtener información de una materia por ID
	router.GET("/api/subjects", handlers.GetAllSubjects)               // Obtener la lista de todas las materias
	router.DELETE("/api/subjects/:subject_id", handlers.DeleteSubject) // Eliminar una materia por ID
	// Rutas para Calificaciones

	router.POST("/api/grades", handlers.CreateGrade)                              // Crear Calificación
	router.PUT("/api/grades/:grade_id", handlers.UpdateGrade)                     // Actualizar Calificación por ID
	router.DELETE("/api/grades/:grade_id", handlers.DeleteGrade)                  // Eliminar Calificación por ID
	router.GET("/api/grades/:grade_id", handlers.GetGrade)                        // Obtener Calificación de una materia de un estudiante
	router.GET("/api/grades/student/:student_id", handlers.GetAllGradesByStudent) // Obtener todas las calificaciones de un estudiante

	router.Run(":8080")
}
