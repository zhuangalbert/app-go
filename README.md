# app-go

# Installation

on this project, you need to install tool on your local :
- Git
- Docker
- Docker Compose
- Postgres (if you need to use local database)

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
this command take a long time.
If you get a screen like the one below, the setup is successful and you are ready to develop your project:

```bash
app_golang      | [2020-04-05 16:56:44]  [6.09ms]  INSERT INTO "users" ("username","name","role","password","created_at","updated_at") VALUES ('zhuangalbert','albert','user','/iH3s9LMv4MmzSk.Pnce17ZhaBsusDCu8CF.6q05HciG1Ie','2020-04-05 16:56:44','2020-04-05 16:56:44') RETURNING "users"."id"  
app_golang      | [1 rows affected or returned ] 
app_golang      | 
app_golang      | (/app/api/seed/seeder.go:37) 
app_golang      | [2020-04-05 16:56:45]  [12.49ms]  INSERT INTO "users" ("username","name","role","password","created_at","updated_at") VALUES ('John Doe','john','user','.H4EtS','2020-04-05 16:56:45','2020-04-05 16:56:45') RETURNING "users"."id"  
app_golang      | [1 rows affected or returned ] 
app_golang      | Running :8099
```

