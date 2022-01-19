# OpenFaaS Gin Demo

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

now test api in
```
http://localhost:31112/function/demo/record/1/
```


---

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

now test api in
```
http://localhost:8000/record/1/
```
