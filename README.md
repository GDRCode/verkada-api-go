
# Verkada API Go Client

This module implements a Go wrapper for Verkada's public API endpoints. Methods across the different product lines are provided as well as structs for query parameters, body parameters, and responses. The public API documentation can be found here: https://apidocs.verkada.com/reference/introduction

Usage of Verkada's public API depends on a properly scoped API key generated within the Command platform.

## Installation

Download the latest version of this package using the 'go get' command:

```go
go get github.com/GDRCode/verkada-api-go
```

Then, import the package for use in your source code
```go
import "github.com/GDRCode/verkada-api-go/pkg/client"
```
## Usage/Examples

All methods are scoped to the Client struct, which contains additional metadata like the API key and access tokens. Start by initializing a client and passing configuration options.

New() will attempt to gather an API key from a ".env" file from the program's working directory. If one is not found, it can be set directly from the Client struct after initialization.

Initializing a client:
```go
c, err := client.New("prod1", true)
//if no .env file
c.Key = "api-key-here"
```
(Optional) Contents of .env
```
API_KEY=api-key-here
```

## Using the API

Methods are organized on a per-product basis using the client's struct fields. Structs for query and body options used in method calls to provide type safety. Some input validation is performed based on requirements found in the API documentation, but .

```go
res, err := client.Camera.GetCameraDevices(*GetCameraDevicesOptions{})

res2, err2 := client.Access.GetAllAccessUsers()

res3, err3 := client.Core.CreateUser(*CreateUserOptions{})
```

## Maintenance, Bug Fixes, and Feature Requests

The client package implements all Verkada public API methods as of August 2025. Tracking API updates and bug fixes done on a best effort basis.