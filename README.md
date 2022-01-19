OpenFaaS Gin Demo
===

- prod

use faas secrets

```bash
vim openfaas-postgresql-password.txt
```

edit postgresql password

```bash
faas-cli secret create openfaas-postgresql-password --from-file openfaas-postgresql-password.txt
faas-cli up -f demo.yml
```

- dev

```bash
cd demo/conf
cp app.ini app.dev.ini
vim app.dev.ini
```

edit config

```bash
cd ..
export run_mode=debug
go run main.go
```
