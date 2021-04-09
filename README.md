# Monolit

This project is to inform you how my code style. This project build using Beego and MySql. You didnt need to import the sql file because the code using remote databe from my website. You also can see the docker registre in (https://hub.docker.com/repository/docker/asahro/test-monolit)

## What You Need

This project required :
- golang v 1.15
- mySql v 10.1.38 (mariaDb)

## Installation

first clone project in a folder named monolit. 
After clone the project, you will get on folder that contant db and golang code.

### Installation db

You didnt need to install bd because I use remote db.

### Installation Golang 

first runing golang  

```bash
go run main.go
```

if there any dependency needed, install them using

```bash
go get git-url
```

## Setup

You can change db setting using from config/app.config
```bash

appname = monolit
httpport = 8080
runmode = dev

db.host         = 45.84.204.205
db.port         = 3306
db.user         = u877696391_test
db.password     = 4Hmad-sahro
db.name         = "u877696391_test"
db.charset      = utf8mb4
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
