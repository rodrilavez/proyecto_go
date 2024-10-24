# proyecto_go

Sistema Escolar - Proyecto en Go
Indice
1. Requisitos Previos
2. Instalacion
3. Configuracion del Proyecto 4. Ejecutar el Proyecto
5. API Endpoints
6. Frontend
7. Estructura del Proyecto
Requisitos Previos
- Go >= 1.18: Descargar desde https://golang.org/dl/
- Git: Descargar desde https://git-scm.com/
- Docker (Opcional, para MySQL en Docker): Descargar desde https://www.docker.com/get-started - MySQL: Instalado localmente o usando Docker.
Instalacion
1. Clonar el Repositorio
Clona este repositorio en tu maquina local:
git clone https://github.com/rodrilavez/proyecto_go.git
2. Navegar al Directorio del Proyecto
Sistema Escolar - Proyecto en Go
cd proyecto_go
3. Configurar el Proyecto Go
Crea un archivo go.mod si no existe, usando:
go mod init nombre-del-proyecto
Luego instala las dependencias necesarias:
go mod tidy
Configuracion del Proyecto
1. Configuracion de la Base de Datos
- Puedes usar MySQL local o usar Docker para lanzar una instancia de MySQL.
Usando Docker:
docker run --name sistema_escolar_db -e MYSQL_ROOT_PASSWORD=proyectogo -e MYSQL_DATABASE=sistema_escolar -p 3306:3306 -d mysql:latest
Esto creara una base de datos llamada sistema_escolar con el usuario root y contraseña

Sistema Escolar - Proyecto en Go
proyectogo.
- Crear Tablas:
Asegurate de tener las tablas necesarias en tu base de datos ejecutando los scripts SQL proporcionados en database/schema.sql.
Ejecutar el Proyecto
1. Compilar y Ejecutar el Proyecto
go run main.go
Esto iniciara el servidor en http://localhost:8080.
2. Prueba del Servidor
Puedes usar herramientas como Postman o curl para probar los endpoints API.
API Endpoints
Estudiantes:
- Crear Estudiante: POST /api/students
- Obtener Todos los Estudiantes: GET /api/students
- Obtener Estudiante por ID: GET /api/students/:student_id

- Actualizar Estudiante: PUT /api/students/:student_id
- Eliminar Estudiante: DELETE /api/students/:student_id
Materias:
- Crear Materia: POST /api/subjects
- Obtener Todas las Materias: GET /api/subjects
- Obtener Materia por ID: GET /api/subjects/:subject_id - Actualizar Materia: PUT /api/subjects/:subject_id
- Eliminar Materia: DELETE /api/subjects/:subject_id
