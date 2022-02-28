# mock-fixtures-server
Golang HTTP server for responding with fixtures files to incoming requests.

This project provides a simple interface as e for solving the problem
of serving static mock fixtures in response to requests intended for an API.

The core functionality allows the mapping of arbitrary urls to be matched against
arbitrary fixture file responses.

```
/api/v1/dogs -> all-dogs.json
/api/v1/dogs/detail -> all-dogs-detail.json
```

