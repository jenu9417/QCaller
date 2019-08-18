# QCaller


## Auth
QCaller uses HTTP basic auth for api authorization.
The following header params should be part of each api request for authentication.

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |


## Create Contact
Create a single contact

```http
POST    /contact
```

#### Request Body
```js
{
	"Name" : "<name>",                     // **Required**
	"SourceID" : "<sourceid>",             // **Required**
	"Country" : "<country>",               // **Required**
	"CountryCode" : "<countrycode>",       // **Required**
	"Number" : "<number>"                  // **Required**
}
```


## Get Contact
Get a single contact

```http
GET    /contact?sourceID=<sourceid>&number=<number>&country=<country>
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `sourceID` | `string` | **Required**. |
| `number` | `string` | **Required**. |
| `country` | `string` | **Required**. |


## Search Contact
Search for a contact. Set `immedidate` to `true` for enabling fast lookup from Aerospike.

```http
GET    /contact?number=<number>&country=<country>&immediate=<true|false>&size=<size>
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `number` | `string` | **Required**. |
| `country` | `string` | **Required**. |
| `immediate` | `string` | **Optional**.  Set the param to true to enable fast lookup. Default false. |
| `size` | `int` | **Optional**.  Numbervof contacts to return. Defaults to 5. No significance if `immediate` is set to true |


## Update Contact
Update a single contact. Two scenarios possible.
a) Update name - If name is updated, the contact detail will be updated.
b) Update number - If number is update, a new contact is upserted. This is because, the id is a compound of number and sourceid.

```http
PUT    /contact
```

#### Request Body
```js
{
	"Name" : "<name>",                     // **Required**
	"SourceID" : "<sourceid>",             // **Required**
	"Country" : "<country>",               // **Required**
	"CountryCode" : "<countrycode>",       // **Required**
	"Number" : "<number>"                  // **Required**
}
```


## Delete Contact
Delete a single contact.

```http
DELETE    /contact?sourceID=<sourceid>&number=<number>&country=<country>
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `sourceID` | `string` | **Required**. sourceID |
| `number` | `string` | **Required**. number |
| `country` | `string` | **Required**. country |


## Bulk Create Contact
Create a list of contacts

```http
POST    /contact/bulk
```

#### Request Body
```js
[{
	"Name" : "<name1>",                     // **Required**
	"SourceID" : "<sourceid1>",             // **Required**
	"Country" : "<country1>",               // **Required**
	"CountryCode" : "<countrycode1>",       // **Required**
	"Number" : "<number1>"                  // **Required**
 },
...
]
```


## Bulk Update Contact
Update a list of contacts

```http
PUT    /contact/bulk
```

#### Request Body
```js
[{
	"Name" : "<name1>",                     // **Required**
	"SourceID" : "<sourceid1>",             // **Required**
	"Country" : "<country1>",               // **Required**
	"CountryCode" : "<countrycode1>",       // **Required**
	"Number" : "<number1>"                  // **Required**
 },
...
]
```


## Response

The response from QCaller api's conforms to the following standard.

### Standard APIs

```js
{
    "HTTPCode": <http-code>,
    "Method": "<http-method>",
    "Status": "<Success | Failure>",
    "Error": {
          "Code": "<error code>",
          "Message": "<error message>",
          "Description": "<error description>"
    },
    "Response": <api response>
}
```

### Bulk APIs

```js
{
    "HTTPCode": <http-code>,
    "Method": "<http-method>",
    "Status": "<Success | Failure>",
    "Error": {
          "Code": "<error code>",
          "Message": "<error message>",
          "Description": "<error description>"
    },
    "Response": {
          "Success": <true|false>,
          "FailedResponses" : [
              {
                  "Code" : "<error-code>",
                  "ID" : "<id of the resource>",
                  "Err" : "error"
              },
            ...
          ]
    }
}
```


## Status Codes

QCaller returns the following status codes in its API:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 206 | `PARTIAL CONTENT` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |
