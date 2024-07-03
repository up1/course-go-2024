# API Gateway Workshop
* Kong
  * Route
  * Service

## 1. Deploy Kong API Gateway
* Database mode => PostgreSQL
* [Kong UI](https://github.com/Kong/kong-manager)

Start database mode = PostgreSQL
```
$docker compose up -d db
$docker compose ps
NAME               IMAGE          COMMAND                  SERVICE   CREATED          STATUS                    PORTS
api-gateway-db-1   postgres:9.5   "docker-entrypoint.s…"   db        33 seconds ago   Up 33 seconds (healthy)   5432/tcp
```

Start Kong
```
$docker compose build kong
$KONG_DATABASE=postgres docker compose --profile database up -d kong
$docker compose ps
NAME                 IMAGE          COMMAND                  SERVICE   CREATED          STATUS                    PORTS
api-gateway-db-1     postgres:9.5   "docker-entrypoint.s…"   db        6 minutes ago    Up 6 minutes (healthy)    5432/tcp
api-gateway-kong-1   kong:latest    "/docker-entrypoint.…"   kong      13 seconds ago   Up 12 seconds (healthy)   0.0.0.0:8000->8000/tcp, 127.0.0.1:8001-8002->8001-8002/tcp, 0.0.0.0:8443->8443/tcp, 127.0.0.1:8444->8444/tcp


$docker compose logs --follow
```

Access to Kong Manage (UI)
* http://localhost:8002/

## 2. Manage Kong with Postman
* Create service
* Create router
* Create plugin
* Testing
* Delete route and service

## 3. Start auth service
```
$docker compose build auth-service
$docker compose up -d auth-service
```

Try step 2 again !!