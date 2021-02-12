# Cinode
![Run tests for package](https://github.com/polarsquad/cinode/workflows/Run%20tests%20for%20package/badge.svg)

This is a Go module to function as a REST API wrapper for [Cinode](https://cinode.com) service.

It is trying to fullfil all the features listed in the Cinode API reference [documentation](https://api.cinode.com/docs/index.html).

## Usage
To initialise the client, one needs to have the following information from Cinode available:
- AccessID
- AccessSecret
- Company ID

The initialisation of the client is done by:  
```golang
client, err := cinode.NewClient([AccessID], [AccessSecret], [CompanyID]
```