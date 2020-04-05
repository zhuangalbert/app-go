# app-go

# Installation

on this project, you need to install tool on your local :
- Git
- Docker
- Docker Compose
- Postgres #if you need to use local database

# How to Setup Local

1. Clone repository using command : 
```bash
git clone git@github.com:zhuangalbert/app-go.git
```

2. enter the directory
```bash
cd app-go
```

3. Create file .env, by copying from .env.example, don't forget to modify .env value 
```bash
cp .env.example .env
```
you can use the mysql database on the local machine with the following configuration : 
```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=database_name
DB_USERNAME=database_username
DB_PASSWORD=database_password
```

4. run docker-compose with command : 
```bash
docker-compose up --build
```


