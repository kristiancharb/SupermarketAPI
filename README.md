# SupermarketAPI

A restful API to fetch, create and delete produce items. Items are stored in memory and are not preserved.

## API

| Method | Path          | Response                                                                                                                                                                                                                                                                                                                                              |
|--------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GET    | /items        | Returns an object containing an items array with all items.                                                                                                                                                                                                                                                                                           |
| GET    | /items/{code} | Returns an object containing the request produce code, name and price.<br>Returns status 404 if code is invalid or doesn't exist in database.                                                                                                                                                                                                         |
| POST   | /items        | Inserts all items in array into the database.<br>Returns a 200 status code if at least one item is successfully inserted.<br>Returns a 400 status code if at least one item in array is invalid (see POST body*)<br>Returns an array of errors if insertion of at least one item fails.<br>Returns a 409 status code if insertion of all items fails. |
| DELETE | /items/{code} | Deletes the item with provided code from the database.<br>Returns a 400 status code if code is invalid. <br>Returns 404 status code if item with code isn't found.                                                                                                                                                                                    |

\* */items  POST Body*
```
{
	"items" : [
        {
            "code": "ABC4-HJC9-SKL9-21KD",
            "name": "Bread",
            "price": 3.46
        },
    ]
}
```
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
