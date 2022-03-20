# Background
Requirement:
- This service will accept a list of usernames in a specific website.
- This will also not terminate the request even if one of those requested username has an error.
- A maximum of 10 usernames can be searched by this service per request
- Usernames should be sorted alphabetically on the response
- Fetched user account details should be cached and invalidate after 2mins

Solution:
- REST webservice (Accepts: JSON Produces: JSON)
- Handle panics correctly by utilizing Gin's default behavior to panic
- Use hexagonal architecture
- We want this service to be scallable in such a way that it can also get user account details from other sources
    - We will base the source on the URL path provided (e.g. /github, /bitbucket, etc.)
    - We will create a factory to get the correct service base on the provided source
- Cache fetched account details using ristretto
    - Cached account detail should also expire after 2mins
- Response is sorted alphebetically
- Requested usernames should not exceed 10

# Installation details
How to install swaggo:
`go install github.com/swaggo/swag/cmd/swag@latest`

What to run when we need to update swagger documentation:
`swag init`

URL to test this application using swagger:
`http://localhost:8080/swagger/index.html`

To build the application:
`go build .`

To run the application:
`go run .`

To run unit tests with coverage:
`go test ./... -cover`

# Technologies used:
- Gin - HTTP router
- Swaggo - Interface Description Language for describing RESTful APIs expressed using JSON
- go-validators - input validation
- go-resty - REST API client to call other REST webservices
- ristretto - fast, concurrent cache library
    - Reason for using ristretto over other caching libraries:
        - https://github.com/dgraph-io/ristretto
        - https://www.start.io/blog/we-chose-ristretto-cache-for-go-heres-why

# Version
1.0.0

# Good to have in the future
Proposed future features:
- Additional request validation like access token
- Using of go flags for dynamic service configuration
- Graceful shutdown