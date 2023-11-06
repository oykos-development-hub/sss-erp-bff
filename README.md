# BFF - Backend for Frontend Gateway Service

## Description

BFF is used as a gatweay between Frontend and microservices

## Installation

### Prerequisites

- Go version 1.20 or higher
- Docker

### Clone the repository

```bash
git clone git@gitlab.sudovi.me:erp/bff-api.git
```

## Usage

BFF is built using Go and provides a GraphQL API for backend services communication. To make use of the API, follow the steps outlined below:

### Running the Service

From the root directory, run the following commands to start the service:

```bash
docker-compose up -d
make start
```

The service will be available on http://localhost:8080/.

### Interacting with the GraphQL API

Access the GraphQL API endpoint at /. For testing and exploring the API, we recommend using Postman.

### Postman collection

You can easily set up Postman to work with our API by clicking the "Run in Postman" button below:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/22875601-b0f708b1-1137-464d-906f-93548cc71d47?action=collection%2Ffork&collection-url=entityId%3D22875601-b0f708b1-1137-464d-906f-93548cc71d47%26entityType%3Dcollection%26workspaceId%3De27156de-19d4-4c8a-ae2d-9b845d9ea484)

With the Postman collection, you can:

- Create new requests to query or mutate data.
- View and run pre-defined requests from our collection.
- Update the request body to experiment with different queries and mutations.
- Inspect the responses to ensure your frontend can consume them correctly.

## Contributing

Please follow the guidelines below for branch naming and commits.

### Commit and branch naming

#### Branch naming convention

1. state the type of change you are making: `build, fix, refactor, feat`
2. add forward slash `/`
3. state the task ID (if applicable) - TSK-123
4. add forward slash `/`
5. add 2-3 words separated by `-` that describe changes you are making
6. Example: `fix/TSK-123/fixing-border-radius`

#### Commit & Push

We use the same convention as for Branch naming.

Only difference is that we use `:` instead of `/` in the commit message. And we describe in the message what we did without `-` between words.

Example: `fix: changed border radius from 4px to 2px`

## License

This project is licensed under the [Internal License](LICENCE).

## Credits

- [Unidoc v3](github.com/unidoc/unipdf/v3)
