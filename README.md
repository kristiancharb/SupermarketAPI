# SupermarketAPI

## Running the app

- Build the container
```
docker build -t supermarket-api .
```
- Run the container in the background
```
docker run -p 8080:8080 -d supermarket-api
```

## Running tests
Build the testing container
```
docker build -t supermarket-tests -f Dockerfile.test .
```
Run the container
```
docker run supermarket-tests
```
