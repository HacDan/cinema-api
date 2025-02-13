# Cinema Ticket Booking System API


## Notes
- I structured this how I would have a full-blown api. I feel this is more in the spirit of the project, as opposed to just the presented challenge. 
- I would have rather seen the specifications request a full json object on post with the id instead of just returning the id. I feel this is more versatile for consumers of the API.

## Running the project
Be sure to have Go installed. If you don't, please check out [webinstall.dev](https://webinstall.dev/go/) for a painless install

```sh
go build .

./cinema-api
```
## Summary of API Specification

### Endpoint: Get Movies

* Path: `/v1/movies`
* Method: `GET`
* Response: JSON containing all movies

#### Description:


##### Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

### Endpoint: Get Movie

* Path: `/v1/movie/{id}
* Method: `GET`
* Response: JSON containing {id} movie data

#### Description:


##### Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

### Endpoint: Post Movies

* Path: `/v1/movie`
* Method: `POST`
* Response: JSON containing all movies

#### Description:


##### Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

### Endpoint: Put Movie

* Path: `/v1/movie/{id}
* Method: `PUT`
* Response: JSON containing all movies

#### Description:


##### Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

### Endpoint: Delete Movie

* Path: `/v1/movie/{id}
* Method: `DELETE`
* Response: Confirmation of movie with {id} is deleted

#### Description:


##### Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```
