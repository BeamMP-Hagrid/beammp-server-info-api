# Introduction

Bsia was built from the ground-up a JSON API that makes it easy for developers to get data from BeamMP servers.

# Get data

To obtain information from a server, we will need to know the IP.

This api has a single endpoint which receives a query param named serverIp.

Example:

```http
GET uri/servers?serverIp=0.0.0.0
```

# Response (happy path)
#### StatusCode = 200
#### Body
```json
[
	{
	"Players": string,
	"PlayersList": string,
	"MaxPlayers": string,
	"IP": string,
	"Port": string,
	"Dport": string,
	"Map": string,
	"Private": boolean,
	"Sname": string,
	"Version": string,
	"Cversion": string,
	"Official": boolean,
	"Owner": string,
	"Sdesc": string,
	"Pps":  string,
	"ModList": string,
	"ModsTotal": string,
	"ModsTotalSize": string
	},
]
```

## Response errors
#### No serverIp query param error.
#### StatusCode = 400
#### Body
```json
{
	"Error": "No serverIp query param.",
	"Message": "You need to send a serverIp query param."
}
```

#### No servers found error.
#### StatusCode = 404
#### Body
```json
{
	"Error": "No servers found",
	"Message": "No servers found with the serverIp query param. Please check the serverIp query param."
}
```

#### Backend error.
#### StatusCode = 500
#### Body
```json
{
	"Error": "Error getting data from beam mp backend.",
	"Message": "Error getting data from beam mp backend."
}
```

#### Parsing data error.
#### StatusCode = 500
#### Body
```json
{
	"Error": "Error reading response body.",
	"Message": "Error reading response body."
}
```