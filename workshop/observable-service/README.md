# Workshop of observable services

## 1. Start LGTM stack (development mode)
```
$docker compose up -d collector

$docker compose ps             
NAME                             IMAGE                     COMMAND                  SERVICE     CREATED         STATUS                   PORTS
observable-service-collector-1   grafana/otel-lgtm:0.6.0   "/bin/sh -c ./run-al…"   collector   2 minutes ago   Up 2 minutes (healthy)   0.0.0.0:3000->3000/tcp, 0.0.0.0:4317-4318->4317-4318/tcp
```

Access to Grafana dashboard
* http://localhost:3000
  * user=admin
  * password=admin


## 2. Start service a
```
$docker compose build service-a
$docker compose up -d service-a
```

Call api
* http://localhost:8081/start

## 3. Start service b
```
$docker compose build service-b
$docker compose up -d service-b
```

Call api
* http://localhost:8082/data

