# administrativa_NoSQL_api
administrativa_NoSQL_api, CRUD para el negocio de novedades, el proyecto está programado en el lenguaje Go y creado con el [framework beego](https://beego.me/).

***Instlaciones Previas:***
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Mongo](https://www.mongodb.com/)

*Recomendado*: 
* [MongoDB Compass](https://www.mongodb.com/products/compass?lang=es-es) o [Robo T3](https://robomongo.org/): Para verificar que los datos se insertan o editan bien.


## Datos
La base de datos está desarrollada en [mongodb](https://www.mongodb.com/) y el backup es: [backupArgo](https://drive.google.com/file/d/1Ybg2hZn2EM72e73s89PoRed_m3kitjiA/view?usp=sharing)

#### * Deplegar base de datos con Contenedor:
Para desplegar en un contenedor se requiere:
* [Docker](https://www.docker.com/get-started)
* [Docker-compose](https://docs.docker.com/compose/install/)

Relizar una copia local del siguiente archivo (Descargar o copiar y pegar en un archo .yml): [Imagen de Mongodb](https://gitlab.com/babermudezb/pruebaoas_brayanbermudez/blob/master/docker-compose.yml)

y subir el contenedor desde visual studio code, por comando o con PORTAINER.

Y crear una base de datos y subir el backup del proyecto: [backupArgo](https://drive.google.com/file/d/1Ybg2hZn2EM72e73s89PoRed_m3kitjiA/view?usp=sharing)

## Configuración del Proyecto.

Para instalar el proyecto realizar los siguientes pasos:
- Para clonar el proyecto en la carpeta local go/src/github.com/udistrital ir a la consola y ejecutar:
```shell 
    cd go/src/github.com/udistrital
```
- Ejecutar:
```shell 
    git clone https://github.com/udistrital/administrativa_NoSQL_api.git
```

- Ir a la carpeta del proyecto:
```shell 
    cd administrativa_NoSQL_api
```

- Instalar dependencias del proyecto:
```shell 
    go get
```

## Ejecución del proyecto

* Ubicado en la raíz del proyecto, ejecutar:
```shell 
    bee run
```
* O si se quiere ejecutar el swager:
```shell 
    bee run -downdoc=true -gendoc=true
```

### Puertos

* El servidor se expone en el puerto: 127.0.0.1:8083

* Para ver la documentación de swagger: [127.0.0.1:8083/swagger/](http://127.0.0.1:8083/swagger/)
    *Nota*: En el swagger sale un error, hacer caso omiso.

