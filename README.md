# SupermarketAPI

## Running the app

Build the container
```
docker build -t kristiancharb/supermarket-api .
```
Run the container in the background
```
docker run -p 8080:8080 -d kristiancharb/supermarket-api
```
Run the container in the foreground
```
docker run -p 8080:8080 -it kristiancharb/supermarket-api
```

## Running tests
Build the testing container
```
docker build -t kristiancharb/supermarket-api:test -f Dockerfile.test .
```
Run the container
```
docker run kristiancharb/supermarket-api:test
```
