order
--
Due to connecitivity issue, I couldn't connect the containerized DB from inside the container. Therefore, the API is running locally (support Mac or Linux).

## File structure
`db/`: contains migration file used by `goose` to handler db migration. It also contains the `Dockerfile` to build db container
`order`: contains the business logic of the three API services.

## System architecture
Database: postgres in container
API: running locally

`handler`: responsible for parsing http requests and writing http response
`usecase`: preceded by `handler`, usecase has the business logic of creating, updating and querying orders. The dependencies are `DBAdaptor` and `DistanceService` which are used to store data and calculate distances.
`pg_adaptor`: is a thin layer above postgres db in which I use `gorm` library to simplify object manipulation.
`mock*.go`: mock files generated by `mockery`

## Limitations
1. Due to time limitation, tests are not complete. Router, handler and business logic (`usecase`) have good testing using mocks. But I don't have time to test the database adaptor.
2. It uses lock to prevent mutliple users taking the same order. This poses performance constraint to the service also doesn't scale very well.
3. Only db is running in container. API should also be containerized once communication issues between containers is resolved
4. No authentication

