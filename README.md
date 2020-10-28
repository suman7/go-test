# go-test

# To run go-docker env

## Create and run container

```
docker-compose up -d
```

## go to container shell

```
docker-compose exec go-dev sh
```

# To run the app

## Run the main app

### Go to the container env shell

```
go run app
```

## To run the test

```
go test app...
```

## To build executable
Note: For the sake of simplicity, it doesnot run with argument yet.
```
go build
```