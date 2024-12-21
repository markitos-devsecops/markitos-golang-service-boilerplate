# Golang Microservice Boilerplate - Marco Antonio - markitos

Este repositorio es una plantilla de microservicio en Golang, diseñada para facilitar la creación de microservicios robustos, escalables y fáciles de mantener. Incluye un conjunto básico de herramientas y patrones para ayudar a los desarrolladores a empezar rápidamente con microservicios en Go.

## Características

- **Microservicio base en Golang**: Configuración lista para empezar con Golang.
- **Dockerizado**: Contenedor Docker para facilitar la ejecución y despliegue.
- **Manejo de dependencias**: Usando Go Modules.
- **Middleware básico**: Implementación de middlewares comunes para autenticación y logging.
- **Configuración centralizada**: Para manejar diferentes entornos de desarrollo, pruebas y producción.
- **Tests**: Configuración de pruebas unitarias y de integración.
- **Documentación de API**: Usando Swagger para generar y mantener documentación.
- **Estructura modular**: Organizado para escalar fácilmente.

## Estructura del Proyecto

El proyecto sigue una estructura modular que separa claramente las responsabilidades:

```
golang-microservice-boilerplate/
├── cmd/
│   └── main.go               # Punto de entrada del microservicio
├── internal/
│   ├── app/                  # Lógica de negocio
│   ├── api/                  # Endpoints y rutas
│   ├── config/               # Configuración del microservicio
│   ├── db/                   # Conexiones a bases de datos
│   ├── middleware/           # Middlewares
│   ├── models/               # Modelos de datos
│   ├── service/              # Lógica de servicio
│   └── utils/                # Funciones utilitarias
├── Dockerfile                # Dockerfile para construir la imagen
├── Makefile                  # Comandos de Make para facilitar tareas comunes
├── go.mod                    # Gestión de dependencias
├── go.sum                    # Suma de las dependencias
└── README.md                 # Este archivo
```

## Requisitos

- Go 1.18 o superior.
- Docker (para ejecutar el microservicio en contenedor).

## Instalación

### 1. Clona el repositorio

```bash
git clone https://github.com/markitos-devsecops/golang-microservice-boilerplate.git
cd golang-microservice-boilerplate
```

### 2. Instala las dependencias

Este proyecto usa Go Modules para gestionar las dependencias. Ejecuta el siguiente comando para descargarlas:

```bash
go mod tidy
```

### 3. Ejecuta el servicio

Puedes ejecutar el microservicio en tu máquina local de la siguiente manera:

```bash
go run cmd/main.go
```

El servicio estará disponible en [http://localhost:8080](http://localhost:8080).

### 4. Docker (opcional)

Si prefieres ejecutar el microservicio en un contenedor Docker, puedes construir la imagen y ejecutar el contenedor con los siguientes comandos:

```bash
docker build -t golang-microservice .
docker run -p 8080:8080 golang-microservice
```

## Documentación de la API

Puedes consultar la documentación interactiva de la API generada por Swagger en [http://localhost:8080/swagger](http://localhost:8080/swagger).

## Contribución

Si deseas contribuir al proyecto, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una rama nueva para tu funcionalidad o corrección de errores (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza los cambios y asegúrate de que los tests pasen (`go test`).
4. Haz un commit de tus cambios (`git commit -am 'Agrega nueva funcionalidad'`).
5. Haz push a la rama (`git push origin feature/nueva-funcionalidad`).
6. Crea un pull request.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.

---

Si necesitas alguna modificación o detalle adicional en el README, no dudes en decírmelo. ¡Estoy aquí para ayudarte!
