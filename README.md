# Databricks SDK for Go

[![Build Status](https://travis-ci.org/syedatifakhtar/databricks-sdk-go.svg?branch=master)](https://travis-ci.org/syedatifakhtar/databricks-sdk-go)

databricks-sdk-go is an unofficial Databricks SDK for the Go programming language. It currently only supports a subset of the operations provided by the [Databricks REST API](https://docs.databricks.com/api/latest/index.html). At this point, there is only support for the following APIs:

* Clusters API
* Workspace API
* Groups API

## Installation

To install the SDK run the following command:

```bash
go get -u github.com/syedatifakhtar/databricks-sdk-go
```

## Usage

The SDK provides an API client as well as multiple endpoints to interact with the different parts of the Databricks REST API.

### Creating an API Client

All operations interacting with Databricks REST API need an API Client. The following example shows how to create a client.

```golang
cl, err := client.NewClient(client.Options{Domain: &domain, Token: &token})
if err != nil {
    panic(err)
}
```

`domain` refers to the name of the Databricks deployment (e.g., `<your-account>.cloud.databricks.com`), and `token` needs to contain the authentication token. See Databricks [authentication documentation](https://docs.databricks.com/api/latest/examples.html) for more details.

### Making Requests

Requests for a given API can be sent using the appropriate endpoint. For instance, the following example shows how to upload a notebook to the workspace.

```golang
endpoint := workspace.Endpoint{
    Client: cl,
}

language := models.PYTHON
content := base64.StdEncoding.EncodeToString([]byte("print('hello world')"))

err := endpoint.Import(&models.WorkspaceImportRequest{
    Path:     path,
    Language: &language,
    Content:  content,
})
if err != nil {
    panic(err)
}
```

See the `examples` folder for more examples on how to use the SDK.

## Development

The information below is only of interests for someone who wants to do development work on the Databricks SDK for Go.

Check the [Go programming language documentation](https://golang.org/doc/) to setup a development environment for Go. 

### Generating Models

Request and response models are autogenerated using [swagger-codegen](https://github.com/swagger-api/swagger-codegen). To generate the models run the following command from the project's root.

```bash
go generate
```

Please note that you need to have `swagger-codegen` installed for the previous command to work.
