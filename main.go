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

	router.Run(":8080")
}
