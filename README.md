# Frog
Simple Go Web Application

## Getting Started
### clone in $GO_PATH
```
cd $GO_PATH
mkdir frog
cd frog
git clone https://github.com/voyageth/frog.git
```

### change directory
```
cd $GO_PATH/frog
```

### setting conf file
```
cp server/conf/app.conf.sample server/conf/app.conf
```
#### edit app.conf
```
...
# Application
app.name = frog
app.secret = #Secret here#
...
db.spec = #userId:password:db-url#
...
```

### go get dependencies
```
go get -t github.com/revel/revel
go get -t github.com/revel/cmd/revel
go get -t github.com/go-gorp/gorp
go get -t github.com/go-sql-driver/mysql
```

### build
```
revel build frog/server tmp
```

## run application
```
revel run frog/server dev 8080
```

## License
Apache License, Version 2.0
http://www.apache.org/licenses/LICENSE-2.0