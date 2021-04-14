# explorer-backend
Spacemesh explorer backend designed to provide data for explorer-frontends

## Explorer Software Architecture
![](https://raw.githubusercontent.com/spacemeshos/product/master/resources/explorer_arch_chart.png)

## Using the Explorer Backend API
The explorer backend provides a public REST API that can be used to get data about a Spacemesh network.
Following these steps to get data about a public Spacemesh network:

1. Obtain a currently available explorer API endpoint for a network from the [permanent Spacemesh public web services url](https://discover.spacemesh.io/networks.json).
2. Build a REST API url using the endpoint. For example, if the url is `https://explorer-api-28.spacemesh.io/` then the network-info is available at `https://explorer-api-28.spacemesh.io/network-info`.
3. Call the RESTP API url to get the json results.

### Paging and pagination
Use the `pagesize` and `page` params to get paginated results. The first page number is 1, so for example, to get the first 20 accounts on TN 128 call: `https://explorer-api-28.spacemesh.io/accounts?pagesize=20&page=1` and to get the next 20 accounts use: `https://explorer-api-28.spacemesh.io/accounts?pagesize=20&page=2`

### API Capabilities
The API is not properly documented yet. To best way to identity the supported API methods is via the api server [source code](https://github.com/spacemeshos/explorer-backend/blob/master/api/httpserver/httpserver.go).
