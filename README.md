# link-shortener

## How to run

- Change config in config/config.yaml, if necessary

### With golang
```bash
go build cmd/app/app.go

./app --storage=in-memory
# or
./app --storage=postgres
```

### With Docker
```bash
docker build -t shorter . 

# ports should be similar with ones in config/config.yaml
docker run -p 8080:8080 -p 8081:8081 shorter in-memory
```