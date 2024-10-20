# link-shortener

## How to run

- Change config in [config/config.yaml](config/config.yaml), if necessary

### With golang
```bash
go build cmd/app/app.go

./app --storage=in-memory
# or
./app --storage=postgres
# default storage: in-memory
```

### With Docker
```bash
docker build -t shorter . 

# ports should be similar with ones in config/config.yaml
docker run -p 8080:8080 -p 8081:8081 shorter in-memory
# or
docker run -p 8080:8080 -p 8081:8081 shorter postgres
```

## How to use
### Pay attention !Original URLs must contain schema!

### GRPC
Base url(Host and Port) for grcp you can find in [config/config.yaml](config/config.yaml)

Endpoints you can find in [api/shortener_v1/shortener.proto](api/shortener_v1/shortener.proto)

### HTTP
Base url(Host and Port) for HTTP you can find in [config/config.yaml](config/config.yaml)

Endpoints:

- GET - redirect to original page
```
/{identifier}
```
- POST - create short url
```
/?url="http://example.com"
```

## Shorting algorithm
Service just generate random string of character until find unique. I'm not using hashing algs 
because there might be collisions.
