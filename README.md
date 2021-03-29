# SupermarketAPI

A restful API to fetch, create and delete produce items. Items are stored in memory and are not preserved.

## Running the app

Run natively
```
go run .
```
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


## API

| Method | Path          | Response                                                                                                                                                                                                                                                                                                                                              |
|--------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GET    | /items        | Returns an object containing an items array with all items.                                                                                                                                                                                                                                                                                           |
| GET    | /items/{code} | Returns an object containing the requested produce code, name and price.<br>Returns status 404 if code is invalid or doesn't exist in database.                                                                                                                                                                                                         |
| POST   | /items        | Inserts all items in post body array into the database.<br>Returns a 200 status code if at least one item is successfully inserted.<br>Returns a 400 status code if at least one item in array is invalid (see POST body*)<br>Returns an array of errors if insertion of at least one item fails (duplicate produce code).<br>Returns a 409 status code if insertion of all items fails. |
| DELETE | /items/{code} | Deletes the item with provided code from the database.<br>Returns a 400 status code if code is invalid. <br>Returns 404 status code if item with code isn't found.                                                                                                                                                                                    |

**/items  POST Body**
```
{
	"items" : [
        {
            "code": "ABC4-HJC9-SKL9-21KD",
            "name": "Bread",
            "price": 3.46
        }, 
        ...
    ]
}
```

## Project Structure

- [main.go](https://github.com/kristiancharb/SupermarketAPI/blob/main/main.go)
    - Contains routing and route handling logic
    - Contains structs for generic error responses
- [service.go](https://github.com/kristiancharb/SupermarketAPI/blob/main/service.go)
    - Provides an interface for produce item operations
    - Service functions return any errors that come up when accessing item store
    - Handles concurrency i.e. spawning goroutines to access item store 
- [database.go](https://github.com/kristiancharb/SupermarketAPI/blob/main/database.go)
    - Contains item store and functions to access/modify item store 
    - Functions send data to caller through channels
- [item.go](https://github.com/kristiancharb/SupermarketAPI/blob/main/item.go)
    - Contains structs used for storage and request/response payloads
    - Helper functions for validating items
- [supermarket_test.go](https://github.com/kristiancharb/SupermarketAPI/blob/main/supermarket_test.go)
    - Unit tests for service functions and other helper functions 

## External Dependencies
- github.com/go-chi/chi: Provides a lightweight router
- github.com/go-chi/render: Provides helpers for decoding request payloads into structs and encoding structs into response payloads

## Continuous Integration Pipeline
- Uses GitHub Actions
- Defined in [ci.yml](https://github.com/kristiancharb/SupermarketAPI/blob/main/.github/workflows/ci.yml)
- On push to main:
    - Builds app container
    - Builds test container
    - Runs test container
    - Uploads app/test images to Docker Hub
- [Docker Hub Repository](https://hub.docker.com/repository/docker/kristiancharb/supermarket-api)

## To Do
- Increase unit test coverage
    - So far we only have test coverage for the service functions, we should include the database functions
    - Look into mocking request/response structs for unit testing route handlers
- Extract concurrency logic from database functions
    - Currently the functions send data through channels
    - It might be cleaner to only have the goroutine logic inside the service functions
    - The service function could spawn a goroutine with an inline function that calls the database function and sends the return values to the channel
- Improve error handling for /items POST route 
    - Currently we return an error for each "bad" produce item but don't specifically acknowledge the items that were inserted
    - It might be clearer to specify which produce codes were successfully inserted (if any) in the response
    - If we were using an actual database it might be more intuitive to rollback successful insertions if there were any "bad" produce items 
