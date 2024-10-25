document.addEventListener('DOMContentLoaded', function () {

    // URLs de la API
    const apiUrl = 'http://localhost:8080/api';
    const studentsUrl = `${apiUrl}/students`;
    const subjectsUrl = `${apiUrl}/subjects`;
    const gradesUrl = `${apiUrl}/grades`;

    // Elementos del DOM
    const studentForm = document.getElementById('studentForm');
    const subjectForm = document.getElementById('subjectForm');
    const gradeForm = document.getElementById('gradeForm');

    const studentList = document.getElementById('studentList');
    const subjectList = document.getElementById('subjectList');
    const gradeList = document.getElementById('gradeList');

    const searchGradesBtn = document.getElementById('searchGradesBtn');
    const searchStudentID = document.getElementById('searchStudentID');

    const updateStudentForm = document.getElementById('updateStudentForm');
    const updateStudentID = document.getElementById('updateStudentID');
    const updateStudentName = document.getElementById('updateStudentName');
    const updateStudentGroup = document.getElementById('updateStudentGroup');
    const updateStudentEmail = document.getElementById('updateStudentEmail');
    const cancelUpdateStudent = document.getElementById('cancelUpdateStudent');

    const updateSubjectForm = document.getElementById('updateSubjectForm');
    const updateSubjectID = document.getElementById('updateSubjectID');
    const updateSubjectName = document.getElementById('updateSubjectName');
    const cancelUpdateSubject = document.getElementById('cancelUpdateSubject');
    
    const updateGradeForm = document.getElementById('updateGradeForm');
    const updateGradeID = document.getElementById('updateGradeID');
    const updateGradeValue = document.getElementById('updateGradeValue');
    const cancelUpdateGrade = document.getElementById('cancelUpdateGrade');

    // --- Funciones para Estudiantes ---
    
    // Obtener todos los estudiantes
    function fetchStudents() {
        fetch(studentsUrl)
            .then(response => response.json())
            .then(data => {
                studentList.innerHTML = '';
                if (data.message) {
                    studentList.innerHTML = `<li>${data.message}</li>`;
                } else {
                    data.forEach(student => {
                        const li = document.createElement('li');
                        // Mostrar ID y detalles del estudiante con botón de edición
                        li.innerHTML = `
                            <strong>ID:</strong> ${student.student_id} - <strong>Nombre:</strong> ${student.name} (${student.group}) - <strong>Email:</strong> ${student.email}
                            <button class="edit-btn" data-id="${student.student_id}">Editar</button>
                            <button class="delete-btn" data-id="${student.student_id}">Eliminar</button>
                        `;
                        studentList.appendChild(li);
                    });
                }
            })
        .catch(error => console.error('Error al obtener estudiantes:', error));
    }

    // Manejo del botón de edición
    studentList.addEventListener('click', function (e) {
        if (e.target.classList.contains('edit-btn')) {
            const studentId = e.target.getAttribute('data-id');
            fetch(`${studentsUrl}/${studentId}`)
                .then(response => response.json())
                .then(student => {
                    // Rellenar el formulario de actualización con los datos del estudiante
                    updateStudentID.value = student.student_id;
                    updateStudentName.value = student.name;
                    updateStudentGroup.value = student.group;
                    updateStudentEmail.value = student.email;
                    updateStudentForm.style.display = 'block';
                })
                .catch(error => console.error('Error al obtener datos del estudiante:', error));
        }
    });

    // Enviar actualización del estudiante
    updateStudentForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const studentId = updateStudentID.value;

        fetch(`${studentsUrl}/${studentId}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                name: updateStudentName.value,
                group: updateStudentGroup.value,
                email: updateStudentEmail.value
            })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            updateStudentForm.style.display = 'none';
            fetchStudents();
        })
        .catch(error => console.error('Error al actualizar estudiante:', error));
    });

    // Cancelar actualización
    cancelUpdateStudent.addEventListener('click', function () {
        updateStudentForm.style.display = 'none';
    });

    // Agregar un nuevo estudiante
    studentForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const name = document.getElementById('studentName').value;
        const group = document.getElementById('studentGroup').value;
        const email = document.getElementById('studentEmail').value;

        fetch(studentsUrl, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, group, email })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            studentForm.reset();
            fetchStudents();
        })
        .catch(error => console.error('Error al crear estudiante:', error));
    });

    // Eliminar un estudiante
    studentList.addEventListener('click', function (e) {
        if (e.target.classList.contains('delete-btn')) {
            const studentId = e.target.getAttribute('data-id');
            fetch(`${studentsUrl}/${studentId}`, {
                method: 'DELETE'
            })
            .then(response => response.json())
            .then(data => {
                alert(data.message);
                fetchStudents();
            })
            .catch(error => console.error('Error al eliminar estudiante:', error));
        }
    });

    // --- Funciones para Materias ---
    
    // Obtener todas las materias
    function fetchSubjects() {
        fetch(subjectsUrl)
            .then(response => response.json())
            .then(data => {
                subjectList.innerHTML = '';
                if (data.message) {
                    subjectList.innerHTML = `<li>${data.message}</li>`;
                } else {
                    data.forEach(subject => {
                        const li = document.createElement('li');
                        // Mostrar el ID y el nombre de la materia junto con botones de edición y eliminación
                        li.innerHTML = `
                            <strong>ID:</strong> ${subject.subject_id} - <strong>Materia:</strong> ${subject.name}
                            <button class="edit-btn" data-id="${subject.subject_id}">Editar</button>
                            <button class="delete-btn" data-id="${subject.subject_id}">Eliminar</button>
                        `;
                        subjectList.appendChild(li);
                    });
                }
            })
        .catch(error => console.error('Error al obtener materias:', error));
    }

    // Manejo del botón de edición para materias
    subjectList.addEventListener('click', function (e) {
        if (e.target.classList.contains('edit-btn')) {
            const subjectId = e.target.getAttribute('data-id');
            fetch(`${subjectsUrl}/${subjectId}`)
                .then(response => response.json())
                .then(subject => {
                    // Rellenar el formulario de actualización con los datos de la materia
                    updateSubjectID.value = subject.subject_id;
                    updateSubjectName.value = subject.name;
                    updateSubjectForm.style.display = 'block';
                })
                .catch(error => console.error('Error al obtener datos de la materia:', error));
        }
    });

    // Enviar actualización de la materia
    updateSubjectForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const subjectId = updateSubjectID.value;

        fetch(`${subjectsUrl}/${subjectId}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: updateSubjectName.value })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            updateSubjectForm.style.display = 'none';
            fetchSubjects();
        })
        .catch(error => console.error('Error al actualizar materia:', error));
    });

    // Cancelar actualización de materia
    cancelUpdateSubject.addEventListener('click', function () {
        updateSubjectForm.style.display = 'none';
    });

    // Agregar una nueva materia
    subjectForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const name = document.getElementById('subjectName').value;

        fetch(subjectsUrl, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            subjectForm.reset();
            fetchSubjects();
        })
        .catch(error => console.error('Error al crear materia:', error));
    });

    // Eliminar una materia
    subjectList.addEventListener('click', function (e) {
        if (e.target.classList.contains('delete-btn')) {
            const subjectId = e.target.getAttribute('data-id');
            fetch(`${subjectsUrl}/${subjectId}`, {
                method: 'DELETE'
            })
            .then(response => response.json())
            .then(data => {
                alert(data.message);
                fetchSubjects();
            })
            .catch(error => console.error('Error al eliminar materia:', error));
        }
    });

    // --- Funciones para Calificaciones ---
    
    // Obtener calificaciones por estudiante, incluyendo nombres
    function fetchGradesByStudent(studentId) {
        fetch(`${gradesUrl}/student/${studentId}`)
            .then(response => response.json())
            .then(data => {
                gradeList.innerHTML = '';
                if (data.message || data.length === 0) {
                    gradeList.innerHTML = `<li>No hay calificaciones para el estudiante con ID: ${studentId}</li>`;
                } else {
                    data.forEach(grade => {
                        const li = document.createElement('li');
                        // Mostrar ID y detalles de la calificación con botón de edición y eliminación
                        li.innerHTML = `
                            <strong>Calificación ID:</strong> ${grade.grade_id} - 
                            <strong>Estudiante:</strong> ${grade.student_name} - 
                            <strong>Materia:</strong> ${grade.subject_name} - 
                            <strong>Calificación:</strong> ${grade.grade}
                            <button class="edit-btn" data-id="${grade.grade_id}">Editar</button>
                            <button class="delete-btn" data-id="${grade.grade_id}">Eliminar</button>
                        `;
                        gradeList.appendChild(li);
                    });
                }
            })
        .catch(error => console.error('Error al obtener calificaciones:', error));
    }

    // Manejo del botón de edición para calificaciones
    gradeList.addEventListener('click', function (e) {
        if (e.target.classList.contains('edit-btn')) {
            const gradeId = e.target.getAttribute('data-id');
            fetch(`${gradesUrl}/${gradeId}`)
                .then(response => response.json())
                .then(grade => {
                    // Rellenar el formulario de actualización con los datos de la calificación
                    updateGradeID.value = grade.grade_id;
                    updateGradeValue.value = grade.grade;
                    updateGradeForm.style.display = 'block';
                })
                .catch(error => console.error('Error al obtener datos de la calificación:', error));
        }
    });

    // Enviar actualización de la calificación
    updateGradeForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const gradeId = updateGradeID.value;

        fetch(`${gradesUrl}/${gradeId}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ grade: parseFloat(updateGradeValue.value) })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
            updateGradeForm.style.display = 'none';
            const studentId = searchStudentID.value;
            if (studentId) {
                fetchGradesByStudent(studentId); // Actualiza la lista después de editar
            }
        })
        .catch(error => console.error('Error al actualizar calificación:', error));
    });

    // Cancelar actualización de calificación
    cancelUpdateGrade.addEventListener('click', function () {
        updateGradeForm.style.display = 'none';
    });

    // Agregar una nueva calificación
    gradeForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const studentID = document.getElementById('gradeStudentID').value;
        const subjectID = document.getElementById('gradeSubjectID').value;
        const gradeValue = document.getElementById('gradeValue').value;

        fetch(gradesUrl, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ student_id: parseInt(studentID), subject_id: parseInt(subjectID), grade: parseFloat(gradeValue) })
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert('Error: ' + data.error);
            } else {
                alert(data.message);
                gradeForm.reset();
            }
        })
        .catch(error => console.error('Error al crear calificación:', error));
    });

    // Manejo del botón de eliminación para calificaciones
    gradeList.addEventListener('click', function (e) {
        if (e.target.classList.contains('delete-btn')) {
            const gradeId = e.target.getAttribute('data-id');
            if (confirm('¿Estás seguro de que deseas eliminar esta calificación?')) {
                fetch(`${gradesUrl}/${gradeId}`, {
                    method: 'DELETE'
                })
                .then(response => response.json())
                .then(data => {
                    alert(data.message);
                    const studentId = searchStudentID.value;
                    if (studentId) {
                        fetchGradesByStudent(studentId);
                    }
                })
                .catch(error => console.error('Error al eliminar calificación:', error));
            }
        }
    });

    // Buscar calificaciones de un estudiante
    searchGradesBtn.addEventListener('click', function () {
        const studentID = searchStudentID.value;
        if (studentID) {
            fetchGradesByStudent(studentID);
        }
    });

    // Inicializar la página obteniendo estudiantes y materias
    fetchStudents();
    fetchSubjects();
});
