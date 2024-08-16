# Tender Mail Sender in Go
Multithreaded mail sender in Go.

### Update database models:
- Install jet generator to specific folder:
```sh
git clone https://github.com/go-jet/jet.git
cd jet && go build -o dir_path ./cmd/jet
```
*Make sure `dir_path` folder is added to the PATH environment variable.*
- Generate models:
```sh
jet -dsn="newtendex:devdbaccess@tcp(127.0.0.1:3307)/newtendex" -path=./.gen -source="mariadb"
```


