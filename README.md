# little-sender

## Descripción
**little-sender** es un proyecto en Go diseñado para gestionar endpoints HTTP de manera sencilla, permitiendo crear, consultar, actualizar y eliminar endpoints almacenados en una base de datos para luego ejecutarlos.

## Estructura del proyecto

```
little-sender/
├── cli/                # Cliente CLI y servidor principal
│   └── server/
│       └── main.go     # Punto de entrada del servidor
├── internal/           # Lógica interna y organización principal
│   ├── config/         # Configuración y conexión a la base de datos
│   ├── model/          # Modelos de datos y DTOs
│   │   └── dto/        # Objetos de transferencia de datos (DTO)
│   ├── routes/         # Definición de rutas HTTP
│   └── service/        # Lógica de negocio y servicios
├── go.mod              # Dependencias del proyecto
├── go.sum              # Checksum de dependencias
└── README.md           # Este archivo
```

## Componentes principales

- **Modelos y DTOs:** Definidos en `internal/model/` y `internal/model/dto/`, representan la estructura de los datos y los objetos de transferencia entre capas.
- **Servicios:** En `internal/service/`, contienen la lógica de negocio para manipular los endpoints.
- **Rutas:** En `internal/routes/`, definen los endpoints HTTP disponibles.
- **Configuración:** En `internal/config/`, se gestiona la conexión a la base de datos y otras constantes.

## Uso básico

1. Clona el repositorio:
	```sh
	git clone https://github.com/Segian/little-sender.git
	cd little-sender
	```
2. Instala las dependencias:
	```sh
	go mod download
	```
3. Ejecuta el servidor:
	```sh
	go run cli/server/main.go
	```

## Notas
- El proyecto utiliza GORM para la gestión de la base de datos.
- Los endpoints y su información se almacenan en una base de datos SQLite (`test.db` por defecto).
- Los DTOs se encuentran en `internal/model/dto/`.

## Estructura de carpetas relevante

- `internal/model/`: Modelos de datos principales.
- `internal/model/dto/`: Objetos de transferencia de datos.
- `internal/service/`: Lógica de negocio.
- `internal/routes/`: Definición de rutas HTTP.
- `internal/config/`: Configuración y conexión a la base de datos.

## Licencia
MIT