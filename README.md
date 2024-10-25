# Sistema Escolar - Proyecto en Go

## Índice

1. [Descripción](#descripción)
2. [Requisitos Previos](#requisitos-previos)
3. [Instalación](#instalación)
4. [Configuración del Proyecto](#configuración-del-proyecto)
5. [Ejecutar el Proyecto](#ejecutar-el-proyecto)
6. [Frontend](#frontend)
7. [API Endpoints](#api-endpoints)
8. [Base de Datos](#base-de-datos)
9. [Estructura del Proyecto](#estructura-del-proyecto)
10. [Notas Adicionales](#notas-adicionales)
11. [Contribuciones](#contribuciones)
12. [Licencia](#licencia)

## Descripción

Este es un sistema escolar desarrollado en Go que permite gestionar estudiantes, materias y calificaciones a través de una API RESTful. Además, incluye un frontend en JavaScript Vanilla que permite interactuar de manera intuitiva con el sistema.

## Requisitos Previos

- **Go >= 1.18**: Descargar desde [golang.org/dl](https://golang.org/dl/)
- **Git**: Descargar desde [git-scm.com](https://git-scm.com/)
- **MySQL**: Instalado localmente o usando Docker
- **Docker** (Opcional, para ejecutar MySQL en Docker): Descargar desde [docker.com/get-started](https://www.docker.com/get-started)
- **Navegador Web**: Para utilizar el frontend

## Instalación

### 1. Clonar el Repositorio

Clona este repositorio en tu máquina local:

git clone https://github.com/rodrilavez/proyecto_go.git

2. Navegar al Directorio del Proyecto

cd proyecto_go


Configuración del Proyecto

1. Configuración de la Base de Datos
   
Puedes usar MySQL localmente o lanzar una instancia de MySQL usando Docker.

Usando Docker
Ejecuta el siguiente comando para crear una instancia de MySQL en Docker:

docker run --name sistema_escolar_db -e MYSQL_ROOT_PASSWORD=proyectogo -e MYSQL_DATABASE=sistema_escolar -p 3306:3306 -d mysql:latest

Esto creará una base de datos llamada sistema_escolar con el usuario root y contraseña proyectogo.

Crear Tablas
Ejecuta los siguientes scripts SQL para crear las tablas necesarias:

CREATE DATABASE sistema_escolar;
USE sistema_escolar;

-- Tabla de estudiantes
CREATE TABLE students (
    student_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    `group` VARCHAR(50),
    email VARCHAR(100)
);

-- Tabla de materias
CREATE TABLE subjects (
    subject_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Tabla de calificaciones
CREATE TABLE grades (
    grade_id INT AUTO_INCREMENT PRIMARY KEY,
    student_id INT,
    subject_id INT,
    grade FLOAT,
    FOREIGN KEY (student_id) REFERENCES students(student_id) ON DELETE CASCADE,
    FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);

2. Configurar el Módulo Go
   
Si no existe un archivo go.mod, créalo utilizando:
go mod init nombre-del-proyecto

Luego, instala las dependencias necesarias:
go mod tidy

Ejecutar el Proyecto

1. Compilar y Ejecutar el Servidor Backend
2. Ejecuta el siguiente comando para iniciar el servidor:
   
go run main.go

Esto iniciará el servidor en http://localhost:8080.

3. Probar el Servidor
Puedes usar herramientas como Postman para probar los endpoints de la API.


Frontend

El proyecto incluye un frontend en JavaScript Vanilla que permite interactuar con el sistema escolar.

1. Ubicación del Frontend
El frontend se encuentra en el directorio frontend/ dentro del proyecto.

3. Configuración del Frontend
No se requiere configuración adicional para el frontend. Todos los archivos necesarios están en el directorio frontend/.

4. Ejecutar el Frontend
Para ejecutar el frontend, simplemente abre el archivo index.html en tu navegador web preferido.

Navega al directorio del frontend:
cd frontend

Abre el archivo index.html:
Windows: Doble clic en index.html o haz clic derecho y selecciona "Abrir con" seguido de tu navegador preferido.

macOS: Doble clic en index.html o haz clic derecho y selecciona "Abrir con" seguido de tu navegador preferido.

Linux: Puedes usar el comando:

xdg-open index.html

Nota: Al abrir el archivo directamente, algunas funcionalidades pueden verse limitadas debido a restricciones de seguridad del navegador (CORS). Si encuentras problemas, considera utilizar un servidor local simple o ajustar la configuración de seguridad de tu navegador.

API Endpoints

Estudiantes
Crear Estudiante: POST /api/students
Obtener Todos los Estudiantes: GET /api/students
Obtener Estudiante por ID: GET /api/students/{student_id}
Actualizar Estudiante: PUT /api/students/{student_id}
Eliminar Estudiante: DELETE /api/students/{student_id}

Materias
Crear Materia: POST /api/subjects
Obtener Todas las Materias: GET /api/subjects
Obtener Materia por ID: GET /api/subjects/{subject_id}
Actualizar Materia: PUT /api/subjects/{subject_id}
Eliminar Materia: DELETE /api/subjects/{subject_id}

Calificaciones
Agregar Calificación: POST /api/grades
Obtener Calificaciones por Estudiante: GET /api/students/{student_id}/grades
Actualizar Calificación: PUT /api/grades/{grade_id}
Eliminar Calificación: DELETE /api/grades/{grade_id}

Base de Datos

La base de datos sistema_escolar contiene las siguientes tablas:

Tabla de Estudiantes (students)
student_id: INT AUTO_INCREMENT PRIMARY KEY
name: VARCHAR(100) NOT NULL
group: VARCHAR(50)
email: VARCHAR(100)

Tabla de Materias (subjects)
subject_id: INT AUTO_INCREMENT PRIMARY KEY
name: VARCHAR(100) NOT NULL

Tabla de Calificaciones (grades)
grade_id: INT AUTO_INCREMENT PRIMARY KEY
student_id: INT
subject_id: INT
grade: FLOAT

Estructura del Proyecto

proyecto_go/
├── main.go
├── go.mod
├── go.sum
├── handlers/
│   ├── h_students.go
│   ├── h_subjects.go
│   └── h_grades.go
├── models/
│   ├── student.go
│   ├── subject.go
│   └── grade.go
├── database/
│   ├── connection.go
├── frontend/
│   ├── index.html
│   ├── styles.css
│   └── script.js
└── README.md

1. main.go: Punto de entrada de la aplicación.
2. handlers/: Controladores para los endpoints de la API.
3. models/: Definición de las estructuras de datos.
4. database/: Conexión de la base de datos.
5. frontend/: Archivos del frontend (HTML, CSS, JavaScript).

   TRABAJO FINAL PARA EL TALLER DE GO HECHO POR Rodrigo Lara Velazquez, Arely Garcia Duran e Ivanna Dominguez.
   Universidad Autonoma de Queretaro - Facultad de Informatica
