# go-meli-integration

## Instalacion

- Descargar el proyecto

```
    $ go get -u github.com/KevinGenaro/go-meli-integration
```
- Instalar dependencias
  
```  
    $ go get
```
## Uso


- Cerrar sesion en mercado libre


- Ejecutar el programa

```
    $ go run main.go 
```
- Iniciar sesion en mercado libre con el usuario de test

```
    Vendedor
    {
        "id": 653305018,
        "nickname": "TETE4210128",
        "password": "qatest9737",
        "site_status": "active",
        "email": "test_user_12827026@testuser.com"
    }    
```     
- Abrir browser en la url

```
    https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=6589031130474375&redirect_uri=http://localhost:8080/auth/code/   
```
- Consultar listado de items, preguntas y ventas del usuario

```
    http://localhost:8080/dashboard
```        
## Informacion extra

- Usuario comprador

 ```     
    Comprador
    {
        "id": 671566680,
        "nickname": "TETE8756233",
        "password": "qatest7203",
        "site_status": "active",
        "email": "test_user_317810@testuser.com"
    }
 ```
- Diagrama

![Alt text](https://github.com/KevinGenaro/go-meli-integration/blob/master/assets/apiLab3%20(1).png?raw=true "Optional Title")   
     