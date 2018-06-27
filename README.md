
# Aplicación básica basada en microservicios

## Caso de uso de la aplicación
Crear un endpoint REST(En el diagrama tiene el nombre de Phoenix, para esto puede usar Go standard o frameworks como Gin o Echo) que permite guardar la
información de un usuario en la base de datos. Debe permitir guardar los siguientes
datos para el usuario: Nombre(Requerido), Email(Requerido), Password(Requerido,
Hash), Verificado, Número de teléfono(Requerido), País, Ciudad, Dirección. Se debe
retornar de una manera clara para el consumo del api cuando se cree el usuario o
los errores de sucedan en el backend para las validaciones y otros tipos de errores,
por favor apegarse a los estándares definidos en la especificación de HTTP para
dichos códigos.


## Requerimientos
1. Golang 1.9
2. Docker 18.03
3. Docker compose 1.21

## NOTA

Agregar en `/etc/host` la siguiente linea para configurar nuestros dominios virtuales utilizados en este proyecto:

```
127.0.0.1   localhost  user.local  phoenix.local
```

## Iniciar servicios

Posicionarse en la carpeta principal donde se encuentra el archivo `docker-compose.yml` y ejecutar el siguiente comando:

```
  docker-compose up -d
```

## Documentación

### UserMicroservice API

Este microservicio se encarga de manejar la gestión de usuarios en la aplicación.

#### Endpoint

1. Mostrar todos los usuarios registrados:
```
 GET /api/users
```
2. Crear un registro de usuario:
```
 POST /api/users
```

### Phoenix API
Este microservicio se encarga de proveer el frontend a los clientes que se conectan a la aplicación.

#### Rutas

1. Pagina de inicio:
```
  GET /
```
2. Pagina para mostrar todos los usuarios:
```
  GET /users
```
3. Formulario para crear usuarios:
```
  GET /users/new
```
