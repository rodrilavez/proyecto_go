package handlers

import (
	"net/http"

	"sistema_escolar/database"
	"sistema_escolar/models"

	"github.com/gin-gonic/gin"
)

// Función para crear un estudiante
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO students (name, `group`, email) VALUES (?, ?, ?)", student.Name, student.Group, student.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante creado exitosamente."})
}

// Función para eliminar un estudiante por ID
func DeleteStudent(c *gin.Context) {
	id := c.Param("student_id")

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM students WHERE student_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el estudiante."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado exitosamente."})
}

// Función para actualizar un estudiante por ID
func UpdateStudent(c *gin.Context) {
	id := c.Param("student_id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("UPDATE students SET name = ?, `group` = ?, email = ? WHERE student_id = ?", student.Name, student.Group, student.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el estudiante."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante actualizado exitosamente."})
}

// Función para obtener todos los estudiantes
func GetAllStudents(c *gin.Context) {
	db, _ := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT student_id, name, `group`, email FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los estudiantes."})
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.StudentID, &student.Name, &student.Group, &student.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los estudiantes."})
			return
		}
		students = append(students, student)
	}

	if len(students) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay ningún estudiante registrado."})
		return
	}

	c.JSON(http.StatusOK, students)
}

// Función para obtener un estudiante por ID
func GetStudent(c *gin.Context) {
	id := c.Param("student_id")
	var student models.Student

	db, _ := database.ConnectDB()
	defer db.Close()

	row := db.QueryRow("SELECT student_id, name, `group`, email FROM students WHERE student_id = ?", id)
	err := row.Scan(&student.StudentID, &student.Name, &student.Group, &student.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado."})
		return
	}
	c.JSON(http.StatusOK, student)
}
