# API Productos

crear un API en GraphQL donde puedas autenticar por medio de JWT, una
vez autenticado poder acceder a los endpoints de products y categories y hacer
el respectivo CRUD de cada una de las entidades, debes guardar esto en una
base de datos PostgreSQL y realizar las pruebas unitarias.

## Contenido

- [Arquitectura de la Solución](#arquitectura)
- [Stack Tecnológico](#stack-tecnológico)
- [Cómo ejecutar este proyecto](#cómo-ejecutar-este-proyecto)
- [Posibles improvements](#posibles-improvements)
- [Comentarios](#comentarios)

----

# Arquitectura

Se ha seleccionado el concepto de diseño de Clean Architecture como la solución nuclear, dicho concepto permite:
la independencia de cualquier framework, testeable, idependencia de la UI e independencia de cualquier elemento externo.

![plot](./images/cleanArc.jpg)

----
La solución tiene 4 capas :

* ###### Domain Layer
  El núcleo de la arquitectura. Entidades
* ###### Aplication Layer
  La segunda capa del núcleo. Casos de uso
* ###### Adapter Layer
  La tercera capa del núcleo. Controladores / Gateways
* ###### External Layer
  La cuarta capa del núcleo. No se escribe mucho código en esta capa

### Consideraciones de diseño

* Las capas de Domain y Application son el centro del diseño, si en algún momento se requiere cambiar de graphql a
  REST, está decisión no tiene impacto en la arquitectura de la aplicación, el core del negocio queda aislado de conceptos
  técnicos.


* Se implementa lógica de validación en las capas: Domain, Application y Adapter.

----

# Stack Tecnológico

* [Go] - Lenguaje de programación
* [GitHub] - Sistema de control de versiones
* [Go Modules] - Herramienta para la gestión de dependencias
* [Echo] - Framework para la creación de API REST
* [graphql-go/graphql] - Framework para la creación de API GraphQl 
* [IntelliJ] - Ide de desarrollo
* [PostgreSQL] - Base de datos
* [Ent] - ORM
* [Testify] - Extensiones para pruebas: assertion, mocking, y test suite
* [Logrus] - Para manejo de logging
* [go-model] - Mapeo Objeto-Objeto

----

# Cómo ejecutar este proyecto

#### Ejecutar el Testing

```bash
$ make test
```

#### Corriendo la aplicación usando Docker

Estos son los pasos para ejecutarlo con `docker-compose`

```bash
# clonar el proyecto
$ git clone https://github.com/efrengarcial/bitsports

# mover al proyecto
$ cd bitsports

# Construir la imagen docker y Ejecutar la aplicación 
$ make run

# check if the containers are running
$ docker ps

# Ejeuctar la llamada
$ curl -X GET "http://localhost:8080/beers" -H  "accept: */*" -H  "Content-Type: application/json-patch+json" 

# Parar
$ docker-compose down
```

----

# Posibles improvements:

* Manejo amigle de errores  en el api de productos y categorias
* Pruebas unitarias en el microservicio de usuarios
* Unificar manejo de logs

----

# Comentarios

* La primera ejecución de las pruebas se demora un poco mientras descarga la imagen docker de postgres:
  postgres:14-alpine

----
