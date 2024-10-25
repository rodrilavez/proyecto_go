package handlers

import (
	"net/http"

	"sistema_escolar/database"
	"sistema_escolar/models"

	"github.com/gin-gonic/gin"
)

// Función para crear una calificación
func CreateGrade(c *gin.Context) {
	var grade models.Grade

	// Intenta enlazar el JSON entrante a la estructura de datos
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos JSON no válidos: " + err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos."})
		return
	}
	defer db.Close()

	// Asegúrate de que todos los campos necesarios están presentes
	if grade.StudentID == 0 || grade.SubjectID == 0 || grade.Grade == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios (StudentID, SubjectID, Grade)."})
		return
	}

	_, err = db.Exec("INSERT INTO grades (student_id, subject_id, grade) VALUES (?, ?, ?)", grade.StudentID, grade.SubjectID, grade.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calificación creada exitosamente."})
}

// Función para eliminar una calificación por ID
func DeleteGrade(c *gin.Context) {
	gradeID := c.Param("grade_id")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos."})
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM grades WHERE grade_id = ?", gradeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la calificación."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calificación eliminada exitosamente."})
}

// Función para obtener una calificación por ID
func GetGrade(c *gin.Context) {
	gradeID := c.Param("grade_id")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos."})
		return
	}
	defer db.Close()

	var grade models.Grade

	// Consulta SQL para obtener la calificación específica
	row := db.QueryRow("SELECT grade_id, student_id, subject_id, grade FROM grades WHERE grade_id = ?", gradeID)
	err = row.Scan(&grade.GradeID, &grade.StudentID, &grade.SubjectID, &grade.Grade)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada."})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// Función para obtener todas las calificaciones de un estudiante con detalles adicionales
func GetAllGradesByStudent(c *gin.Context) {
	studentID := c.Param("student_id")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos."})
		return
	}
	defer db.Close()

	// Estructura para cada calificación
	type GradeDetail struct {
		GradeID     int     `json:"grade_id"`
		Grade       float64 `json:"grade"`
		StudentName string  `json:"student_name"`
		SubjectName string  `json:"subject_name"`
	}

	// Slice de tipo GradeDetail para almacenar todas las calificaciones
	var grades []GradeDetail

	// Consulta SQL ajustada para devolver detalles adicionales
	rows, err := db.Query(`
        SELECT g.grade_id, g.grade, s.name as student_name, sub.name as subject_name
        FROM grades g
        JOIN students s ON g.student_id = s.student_id
        JOIN subjects sub ON g.subject_id = sub.subject_id
        WHERE g.student_id = ?`, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener las calificaciones."})
		return
	}
	defer rows.Close()

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var grade GradeDetail
		if err := rows.Scan(&grade.GradeID, &grade.Grade, &grade.StudentName, &grade.SubjectName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear las calificaciones."})
			return
		}
		// Añadir la calificación al slice de grades
		grades = append(grades, grade)
	}

	if len(grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay calificaciones para este estudiante."})
		return
	}

	// Responder con el slice de calificaciones
	c.JSON(http.StatusOK, grades)
}

// Función para actualizar una calificación por ID
func UpdateGrade(c *gin.Context) {
	gradeID := c.Param("grade_id")
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos."})
		return
	}
	defer db.Close()

	// Actualizar la calificación en la base de datos
	_, err = db.Exec("UPDATE grades SET grade = ? WHERE grade_id = ?", grade.Grade, gradeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la calificación."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calificación actualizada exitosamente."})
}
