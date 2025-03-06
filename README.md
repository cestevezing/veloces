# Order Management API
## Descripción del Proyecto
Este proyecto es una API RESTful desarrollada en Go que permite gestionar órdenes de compra y productos. La aplicación sigue la arquitectura hexagonal (Clean Architecture), lo que garantiza una separación clara entre las capas de negocio, datos y presentación. Además, implementa características avanzadas como idempotencia, manejo de transacciones y concurrencia para asegurar un sistema robusto y eficiente.

La API permite:

- Crear órdenes de compra.
- Consultar productos y sus detalles.
- Actualizar el stock de productos de manera segura.
- Garantizar la consistencia mediante el uso de Redis para idempotencia y MySQL para almacenamiento persistente.

# Tecnologías Utilizadas
1. Go (Golang) :
Lenguaje principal utilizado para desarrollar la API debido a su rendimiento, simplicidad y soporte nativo para concurrencia.
2. Fiber :
Framework web rápido y ligero inspirado en Express.js. Se utiliza para manejar las solicitudes HTTP y proporcionar endpoints RESTful.
3. MySQL :
Base de datos relacional utilizada para almacenar información sobre productos, órdenes e ítems de órdenes.
4. Redis :
Base de datos en memoria utilizada para implementar idempotencia y mejorar el rendimiento de operaciones repetitivas.
5. GORM :
ORM (Object Relational Mapper) para interactuar con MySQL. Simplifica las consultas SQL y proporciona herramientas para manejar relaciones entre tablas.
6. Docker y Docker Compose :
Herramientas utilizadas para contenerizar la aplicación y sus dependencias (MySQL y Redis). Facilitan la configuración y ejecución del proyecto en diferentes entornos.

# Características Principales
1. Idempotencia :
 - Implementada mediante Redis para evitar duplicados en la creación de órdenes. Cada solicitud incluye un encabezado Idempotency-Key que se verifica antes de procesar la solicitud.
2. Arquitectura Hexagonal :
- Separación clara entre capas:
  - Controladores : Manejan las solicitudes HTTP.
  - Servicios : Implementan la lógica de negocio.
  - Repositorios : Interactúan con la base de datos.

# Cómo Ejecutar el Proyecto
## Prerrequisitos
1. Docker :
    - Asegúrate de tener Docker instalado en tu máquina. Puedes descargarlo desde aquí .
2. Docker Compose :
    - Docker Compose está incluido en Docker Desktop para macOS y Windows. Para Linux, puedes instalarlo siguiendo las instrucciones [aquí](https://docs.docker.com/compose/install/?spm=2b75ac3d.6b29c267.0.0.186710b95c1mq6) .

## Pasos para Ejecutar
1. Clonar el Repositorio: 
    ```bash
    git clone https://github.com/cestevezing/veloces.git
    cd veloces
    ```
2. Iniciar los Servicios con Docker Compose:
    ```bash
    docker compose up -d
    ```
    Esto iniciará los siguientes servicios:
     - MySQL : Base de datos para almacenar productos y órdenes.
     - Redis : Base de datos en memoria para idempotencia.
     - API : La aplicación Go que expone los endpoints RESTful.

3. Verificar los Logs :
    - Los logs de los servicios estarán disponibles en la terminal. Si todo funciona correctamente, verás mensajes indicando que los servicios están listos.

4. Acceder a la API :
    - Una vez que los servicios estén en funcionamiento, puedes acceder a la API en http://localhost:8080.

# Endpoints Disponibles
## Productos
 - GET /products :
    - Obtiene la lista de todos los productos.
 - GET /products/:id :
    - Obtiene los detalles de un producto específico.
 - PUT /products/{id}/stock :
    - Actualiza el stock de un producto específico. Requiere un encabezado Idempotency-Key.
## Órdenes
 - POST /orders :
    - Crea una nueva orden. Requiere un encabezado Idempotency-Key.
 - GET /orders/:id :
    - Obtiene los detalles de una orden específica, incluyendo sus ítems.

# Contacto
Si tienes preguntas o sugerencias, no dudes en contactarme:

 - Email: cestevezing@gmail.com
 - GitHub: [cestevezing](https://github.com/cestevezing)