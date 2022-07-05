# wrdle
Golang implementation of the [wordle game](https://www.nytimes.com/games/wordle/index.html)

to start
```shell
go run cmd/main.go
```

run mysql container
```shell
docker run --name=mysql-local -p 3306:3306 -d mysql/mysql-server:latest
```