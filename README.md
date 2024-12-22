# Golang Markitos Service Boilerplate

Este repositorio es una **plantilla para crear servicios en Golang**, diseñada con **prácticas de DevSecOps** en mente. La plantilla está orientada a la seguridad desde el primer momento del ciclo de vida del desarrollo, lo que incluye análisis estáticos de seguridad (SAST), pruebas automatizadas y un diseño modular que facilita la escalabilidad.

Este boilerplate tiene como objetivo ser una solución robusta para comenzar a desarrollar servicios de manera eficiente, asegurando las mejores prácticas de **seguridad** y **pruebas**. Además, se ha diseñado para ser fácilmente extensible y adaptable.

**¡Contribuciones, sugerencias y críticas son bienvenidas!** Si te interesa aprender más sobre **DevSecOps** y seguridad en el ciclo de vida del software, puedes seguir mi canal de YouTube:

[markitos_devsecops](https://www.youtube.com/@markitos_devsecops)

---

## Características Principales

- **Estructura Modular**: Separa las funcionalidades en capas, lo que facilita el mantenimiento y la escalabilidad del microservicio.
- **Seguridad Integrada**: Implementa prácticas de **DevSecOps** para garantizar la seguridad del código desde el principio, con herramientas de análisis de seguridad como **Semgrep** y **Gitleaks**.
- **Automatización de Pruebas**: Incluye pruebas unitarias, de integración y de seguridad automatizadas.
- **Contenedores Docker**: Se proporciona soporte para Docker, lo que facilita la creación de contenedores para el desarrollo y el despliegue.
- **PostgreSQL y Migraciones**: Implementación de base de datos con PostgreSQL y migraciones automatizadas mediante herramientas como **Migrate** y **sqlc**.
- **Arquitectura RESTful**: Exposición de API RESTful utilizando **Gin**, un framework web para Go.
- **CI/CD con Makefile**: El proyecto incluye un `Makefile` para facilitar la construcción, el análisis estático de seguridad, las migraciones de base de datos, la ejecución de pruebas, entre otras tareas.

---

## Estructura del Proyecto

El proyecto sigue las mejores prácticas de desarrollo modular para servicios. La estructura de directorios está organizada en capas que separan la lógica de dominio, infraestructura y servicios. Además, incluye un directorio `testsuite` para pruebas.

```
golang-microservice-boilerplate/
├── internal/                      # Lógica interna del microservicio
│   ├── domain/                    # Entidades y lógica de dominio
│   ├── infrastructure/            # Implementación de infraestructuras (API, base de datos, etc.)
│   ├── services/                  # Lógica de los servicios del microservicio
├── testsuite/                     # Tests, incluyendo pruebas de seguridad
│   ├── internal/infrastructure    # Tests de infraestructuras (API, base de datos, etc.)
│   ├── internal/domain/           # Tests para la capa de dominio
│   └── internal/services/         # Tests para los servicios
├── Makefile                       # Comandos útiles para construir, probar y analizar el proyecto
├── go.mod                         # Gestión de dependencias en Go
├── go.sum                         # Archivo de suma de dependencias
└── README.md                      # Este archivo
```

---

## Requisitos

- **Go 1.18+**: Necesitas tener Go 1.18 o superior instalado.
- **Docker**: Se recomienda usar Docker para ejecutar el microservicio en contenedores.
- **PostgreSQL**: Si deseas usar una base de datos PostgreSQL para tu microservicio.

---

## Instalación

### 1. Clonar el Repositorio

```bash
git clone https://github.com/markitos-devsecops/golang-microservice-boilerplate.git
cd golang-microservice-boilerplate
```

### 2. Instalar Dependencias

Este proyecto usa **Go Modules** para gestionar dependencias. Para descargarlas, ejecuta:

```bash
go mod tidy
```

### 3. Ejecutar el Microservicio

Puedes iniciar el microservicio de forma local con el siguiente comando:

```bash
go run internal/app/main.go
```

Esto ejecutará el microservicio en **http://localhost:3000**.

### 4. Usar Docker (opcional)

Si prefieres ejecutar el microservicio dentro de un contenedor Docker, puedes construir la imagen y ejecutar el contenedor con:

```bash
docker build -t golang-microservice .
docker run -p 3000:3000 golang-microservice
```

---

## Comandos Makefile

El `Makefile` incluye una serie de comandos útiles para automatizar tareas comunes como pruebas, migraciones de bases de datos, análisis de seguridad y construcción de imágenes Docker.

### Comandos de Ejecución

- **Ejecutar el microservicio**:
  ```bash
  make run
  ```

### Comandos de Pruebas

- **Ejecutar pruebas (todos los tests)**:
  ```bash
  make test
  ```

- **Ejecutar pruebas con salida detallada**:
  ```bash
  make testv
  ```

### Comandos de Migraciones de Base de Datos

- **Crear la base de datos**:
  ```bash
  make createdb
  ```

- **Eliminar la base de datos**:
  ```bash
  make dropdb
  ```

- **Ejecutar migraciones**:
  ```bash
  make migrate-up
  ```

- **Deshacer la última migración**:
  ```bash
  make migrate-down
  ```

### Comandos de Seguridad (DevSecOps)

- **Análisis estático de seguridad con Semgrep (SAST)**:
  ```bash
  make appsec-sast-sca
  ```

- **Detectar secretos expuestos con Gitleaks**:
  ```bash
  make appsec-gitleaks
  ```

### Comandos de Construcción de Imágenes Docker

- **Construir la imagen Docker**:
  ```bash
  make image-build
  ```

- **Subir la imagen a GitHub Container Registry**:
  ```bash
  make image-push
  ```

---

## Pruebas

El proyecto incluye pruebas unitarias e integradas que cubren tanto la lógica de negocio como las interacciones con la base de datos. Además, el **Makefile** permite ejecutar análisis de seguridad, pruebas de integración y otras pruebas relevantes de manera automatizada.

Para ejecutar las pruebas de seguridad y las pruebas regulares, puedes usar:

```bash
make test
make security
```

### Pruebas de Seguridad

El proyecto integra herramientas como **Semgrep** para el análisis estático del código y **Gitleaks** para detectar secretos expuestos en el código fuente.

---

## Contribución

Este proyecto es una propuesta abierta y cualquier contribución es bienvenida. Si tienes sugerencias, mejoras o correcciones de seguridad, siéntete libre de abrir un **issue** o **pull request**.

Para contribuir:

1. Haz un **fork** del repositorio.
2. Crea una rama para tu nueva funcionalidad o corrección de errores:
   ```bash
   git checkout -b feature/nueva-funcionalidad
   ```
3. Realiza tus cambios y asegúrate de que todas las pruebas pasen:
   ```bash
   go test ./...
   ```
4. Haz un **commit** de tus cambios:
   ```bash
   git commit -am 'Agrega nueva funcionalidad'
   ```
5. Haz **push** a tu rama:
   ```bash
   git push origin feature/nueva-funcionalidad
   ```
6. Crea un **pull request**.

---

## Licencia

Este proyecto está bajo la licencia **MIT**. Consulta el archivo [LICENSE](LICENSE) para más detalles.

---

## Canal de YouTube

Si te interesa aprender más sobre **DevSecOps**, **seguridad en aplicaciones** y buenas prácticas de desarrollo de software, puedes seguir mi canal de YouTube:

[markitos_devsecops](https://www.youtube.com/@markitos_devsecops)