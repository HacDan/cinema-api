# Cinema Ticket Booking System API


## Notes
- I structured this how I would have a full-blown api. I feel this is more in the spirit of the project, as opposed to just the presented challenge. 
- The code could easily be expanded to use a database. For now, it is just using an in-memory `map`.
- I would have rather seen the specifications request a full json object on post with the id instead of just returning the id. I feel this is more versatile for consumers of the API.

## Running the project
Be sure to have Go installed. If you don't, please check out [webinstall.dev](https://webinstall.dev/go/) for a painless install

```sh
go build .

./cinema-api
```
## Summary of API Specification

### Endpoint: Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.

#### Description:

Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt was awarded.

How many points should be earned are defined by the rules below.

Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

#### Description
A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:
```json
{ "points": 32 }
```

