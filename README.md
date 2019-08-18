# QCaller
The simplfied truecaller

![](https://github.com/jenu9417/QCaller/blob/master/external/others/pics/QCaller.png)


## Features
- High performance, horizontally scalable platform
- CRUD APIs :  for better contact management
- Support for bulk operations
- ElasticSearch as the primary data store for efficient storage and search
- Aerospike as the cache layer for fast contact lookup in Search API
- Support for extended contact validation
- Supports HTTP Basic Auth


## Requirements
- Go (1.8)
- ElasticSearch (2.4.6+)
- Aerospike (3.13+)

## Data Model
```
ID            -   String    // Compound id made up of source-id and number
Name          -   String    // Contact name
SourceID      -   String    // Source from which the contact was uploaded
Country       -   String    // Country for which the contact is valid
CountryCode   -   String    // International Calling Code for the country
Number        -   String    // Number
LastUpdated   -   Long      // UTC timestamp of last modification
```

## Documentation
More detailed documentation is present [here](https://github.com/jenu9417/QCaller/tree/master/external/documentation).
## APIs
- > [<code>GET</code> /contact](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#get-contact)
- > [<code>POST</code> /contact](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#create-contact)
- > [<code>PUT</code> /contact](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#update-contact)
- > [<code>DELETE</code> /contact](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#delete-contact)
- > [<code>POST</code> /contact/bulk](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#bulk-create-contact)
- > [<code>PUT</code> /contact/bulk](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#bulk-update-contact)
- > [<code>GET</code> /contact/search](https://github.com/jenu9417/QCaller/blob/master/external/documentation/API.md#search-contact)

## How to use?
Run the make file to create package
> make all

For creating a package with all dependencies
> make fatty

Find the tar ball under the folder `/target`
```
--QCaller.tar.gz
  |-bin
  |   |-QCaller
  |   |-config.json
  |-deps                    // optional
  |   |-as
  |   |  |-aerospike.tar.gz
  |   |-es
  |   |  |-elasticsearch.tar.gz
  |-postman
  |   |-QCaller_apis.json
  |-scripts
  |   |-as
  |   |  |-InitAS.sh
  |   |  |-StartAS.sh
  |   |  |-StopAS.sh
  |   |  |-StatusAS.sh
  |   |-es
  |   |  |-InitES.sh
  |   |  |-StartES.sh
  |   |-Server.sh
  |-README.md
```
Run the QCaller
> bin/scripts/Server.sh
