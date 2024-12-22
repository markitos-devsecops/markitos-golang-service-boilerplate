# Golang Markitos Service Boilerplate

Este repositorio es una **plantilla para crear servicios en Golang**, diseñada con **prácticas de DevSecOps** en mente. La plantilla está orientada a la seguridad desde el primer momento del ciclo de vida del desarrollo, lo que incluye análisis estáticos de seguridad (SAST), pruebas automatizadas y un diseño modular que facilita la escalabilidad.

Este boilerplate tiene como objetivo ser una solución robusta para comenzar a desarrollar servicios de manera eficiente, asegurando las mejores prácticas de **seguridad** y **pruebas**. Además, se ha diseñado para ser fácilmente extensible y adaptable.

**¡Contribuciones, sugerencias y críticas son bienvenidas!** Si te interesa aprender más sobre **DevSecOps** y seguridad en el ciclo de vida del software, puedes seguir mi canal de YouTube:

[markitos_devsecops](https://www.youtube.com/@markitos_devsecops)

---

## Características Principales

- **Estructura Modular**: Separa las funcionalidades en capas, lo que facilita el mantenimiento y la escalabilidad del servicio.
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
golang-service-boilerplate/
├── internal/                      # Lógica interna del servicio
│   ├── domain/                    # Entidades y lógica de dominio
│   ├── infrastructure/            # Implementación de infraestructuras (API, base de datos, etc.)
│   ├── services/                  # Lógica de los servicios/casos de uso/features/application/...
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
- **Docker**: Se recomienda usar Docker para ejecutar el servicio en contenedores.
- **PostgreSQL**: Si deseas usar una base de datos PostgreSQL para tu servicio.

---

## Instalación

### 1. Clonar el Repositorio

```bash
git clone https://github.com/markitos-devsecops/golang-service-boilerplate.git
cd golang-service-boilerplate
```

### 2. Instalar Dependencias

Este proyecto usa **Go Modules** para gestionar dependencias. Para descargarlas, ejecuta:

```bash
go mod tidy
```

### 3. Ejecutar el Servicio

Puedes iniciar el servicio de forma local con el siguiente comando:

```bash
go run internal/app/main.go
```

Esto ejecutará el servicio en **http://localhost:3000**.

### 4. Usar Docker (opcional)

Si prefieres ejecutar el servicio dentro de un contenedor Docker, puedes construir la imagen y ejecutar el contenedor con:

```bash
docker build -t golang-service .
docker run -p 3000:3000 golang-service
```

---

## Comandos Makefile

El `Makefile` incluye una serie de comandos útiles para automatizar tareas comunes como pruebas, migraciones de bases de datos, análisis de seguridad y construcción de imágenes Docker.

### Comandos de Ejecución

- **Ejecutar el servicio**:
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

---

## Comentario de main.go

El archivo `main.go` es el punto de entrada del servicio. Aquí es donde se inicializa la aplicación y se configuran los diferentes componentes necesarios para su funcionamiento, como el servidor HTTP, las rutas de la API y las conexiones a la base de datos.

```go
// main.go
package main

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  // Importar otros paquetes necesarios
)

func main() {
  //...
  //...
  //------------------------------------------------
  // Migrar el esquema (migrate)
  // Solo usarlo en caso de no hacer uso de migrate
  // Comentar este bloque si hacemos uso de migrate
  //------------------------------------------------
  err := db.AutoMigrate(&domain.Boiler{})
  if err != nil {
    log.Fatal(err)
  }
  //------------------------------------------------
  //...
  //...
}
```

---

## Documentación del Makefile

El `Makefile` es una herramienta poderosa para automatizar tareas comunes en el desarrollo del servicio. A continuación se describen los comandos disponibles en el `Makefile`:

### Comandos de Ejecución

- **run**: Ejecuta el servicio localmente.
  ```bash
  make run
  ```

### Comandos de Pruebas

- **test**: Ejecuta todas las pruebas.
  ```bash
  make test
  ```

- **testv**: Ejecuta todas las pruebas con salida detallada.
  ```bash
  make testv
  ```

### Comandos de Migraciones de Base de Datos

- **createdb**: Crea la base de datos.
  ```bash
  make createdb
  ```

- **dropdb**: Elimina la base de datos.
  ```bash
  make dropdb
  ```

- **migrate-up**: Ejecuta las migraciones de la base de datos.
  ```bash
  make migrate-up
  ```

- **migrate-down**: Deshace la última migración de la base de datos.
  ```bash
  make migrate-down
  ```

### Comandos de Seguridad (DevSecOps)

- **appsec-sast-sca**: Ejecuta el análisis estático de seguridad con Semgrep.
  ```bash
  make appsec-sast-sca
  ```

- **appsec-gitleaks**: Detecta secretos expuestos con Gitleaks.
  ```bash
  make appsec-gitleaks
  ```

### Comandos de Construcción de Imágenes Docker

- **image-build**: Construye la imagen Docker del servicio.
  ```bash
  make image-build
  ```

- **image-push**: Sube la imagen Docker al GitHub Container Registry.
  ```bash
  make image-push
  ```

### Ejemplos de Uso con Variables

Algunos comandos del `Makefile` permiten pasar parámetros opcionales, como los comandos de migración que pueden aceptar una `VERSION` para especificar una versión de la migración:

### Ejemplos de Uso con Variables

Algunos comandos del `Makefile` permiten pasar parámetros opcionales, como los comandos de migración que pueden aceptar una `VERSION` para especificar una versión de la migración. Si no se especifica `VERSION`, se usará `VERSION=1` por defecto:

- **migrate-up** con `VERSION`:
  ```bash
  make migrate-up VERSION=2
  ```

- **migrate-up** sin `VERSION`:
  ```bash
  make migrate-up
  ```

- **migrate-down** con `VERSION`:
  ```bash
  make migrate-down VERSION=1
  ```

- **migrate-down** sin `VERSION`:
  ```bash
  make migrate-down
  ```

### Ejemplos de Uso con Variables en Construcción

También puedes pasar parámetros opcionales a los comandos de construcción de imágenes Docker. Si no se especifica `TAG`, se usará `TAG=1.0.0` por defecto:

- **image-build** con `TAG`:
  ```bash
  make image-build TAG=v1.0.1
  ```

- **image-build** sin `TAG`:
  ```bash
  make image-build
  ```

- **image-push** con `TAG`:
  ```bash
  make image-push TAG=v1.0.1
  ```

- **image-push** sin `TAG`:
  ```bash
  make image-push
  ```

### Ejemplos de Uso con Variables en Construcción

También puedes pasar parámetros opcionales a los comandos de construcción de imágenes Docker:

- **image-build** con `TAG`:
  ```bash
  make image-build TAG=v1.0.1
  ```

- **image-push** con `TAG`:
  ```bash
  make image-push TAG=v1.0.1
  ```

---

## Contribuir

Si deseas contribuir a este proyecto, por favor sigue los siguientes pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y haz commit (`git commit -am 'Añadir nueva funcionalidad'`).
4. Sube tus cambios (`git push origin feature/nueva-funcionalidad`).
5. Abre un Pull Request.

## Licencia

Este proyecto está licenciado bajo la Licencia GNU General Public License v3.0. Consulta el archivo [LICENSE](LICENSE) para más detalles.