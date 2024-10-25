package handlers

import (
	"net/http"

	"sistema_escolar/database"
	"sistema_escolar/models"

	"github.com/gin-gonic/gin"
)

// Función para crear una materia
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO subjects (name) VALUES (?)", subject.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Materia creada exitosamente."})
}

// Función para eliminar una materia por ID
func DeleteSubject(c *gin.Context) {
	id := c.Param("subject_id")

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM subjects WHERE subject_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la materia."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada exitosamente."})
}

// Función para actualizar una materia por ID
func UpdateSubject(c *gin.Context) {
	id := c.Param("subject_id")
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := database.ConnectDB()
	defer db.Close()

	_, err := db.Exec("UPDATE subjects SET name = ? WHERE subject_id = ?", subject.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la materia."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Materia actualizada exitosamente."})
}

// Función para obtener una materia por ID
func GetSubject(c *gin.Context) {
	id := c.Param("subject_id")
	var subject models.Subject

	db, _ := database.ConnectDB()
	defer db.Close()

	row := db.QueryRow("SELECT subject_id, name FROM subjects WHERE subject_id = ?", id)
	err := row.Scan(&subject.SubjectID, &subject.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada."})
		return
	}
	c.JSON(http.StatusOK, subject)
}

func GetAllSubjects(c *gin.Context) {
	db, _ := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT subject_id, name FROM subjects")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las materias."})
		return
	}
	defer rows.Close()

	var subjects []models.Subject
	for rows.Next() {
		var subject models.Subject
		if err := rows.Scan(&subject.SubjectID, &subject.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear las materias."})
			return
		}
		subjects = append(subjects, subject)
	}

	if len(subjects) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay ninguna materia registrada."})
		return
	}

	c.JSON(http.StatusOK, subjects)
}
