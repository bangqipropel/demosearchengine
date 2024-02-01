# Demo Search Engine - In Memory

## Introduction

This project is a Go application designed for managing and querying metadata. It provides an API to store metadata records and allows searching through them based on various criteria. The application uses a concurrent-safe in-memory storage for metadata, making it suitable for environments where thread safety is a concern.

## Features

- Store metadata records with comprehensive details including title, version, maintainers, company, etc.
- Enable searching for metadata based on fields such as title, version, company, website, and license.
- Implement thread-safe in-memory storage using `sync.Map` for efficient concurrent access.

## Getting Started

### Prerequisites

- Go (Version go1.18.1 linux/amd64)

### Installation

Clone the project repository to your local machine:

```bash
git clone https://github.com/bangqipropel/demosearchengine.git
cd demosearchengine
```

### Running the Application
To start the server, run:

```bash
go run main.go
```
The server will start on http://localhost:8080.

## Usage

### Upload Metadata
To upload metadata, send a POST request to /metadata with the metadata content in YAML format.

Example using curl:
```bash
curl -X POST --data-binary @payload2.yml -H "Content-Type: application/x-yaml" http://localhost:8080/metadata
```
I have some basic validations for the payload, such as null checks, email validations and will return bad request errors if payload cannot pass validations.
Once payload is successfully uploaded, the 200 response will contain a uuid, representing the id for this doc, and user can use this id for specific doc query in the future.

### Searching Metadata

#### ID Search:
If user knows the specific doc id:
```bash
curl "http://localhost:8080/search?id=13acb5de-bf61-40fd-84a1-8285724bb95d"
```

#### General search:
To search for metadata, send a GET request to /search with query parameters.
`base_query` takes some text and searches this text against some text fields like title, company and description. For other parameters (`company`, `version`, `license`, `website`), the project will apply them as filters. If filter parameters are left blank, the project will only search with base query.


Example using curl:
```bash
curl "http://localhost:8080/search?title=Valid%20App&version=0.0.1"
```

## Run Tests

To run all unit tests, run the following command in the root path.

```bash
go test ./...
```

## Other Notes

This demo app is finished withiin a very short time. Given more time, I would introduce some inverted index mechanism to have better searching perfomance or leverage some existing index and search module such as Bleve. 

In structure, I would have a Makefile to run the feature and all the test cases if I have more time.

In the test part, I would like to have higher test coverage if more time available.