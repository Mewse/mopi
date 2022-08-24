# MOPI - Mock API

## Purpose
MOPI is a basic API simulator. You can register a json response to an endpoint and it will return 
that response anytime the URL is called. This can be used in a deployed system where parts of it
need to be mocked out in an interactive way.

## API

### [POST] /register
--------------------
This endpoint is to register responses against URLs. Re-registering the same URL will replace it.

#### Body
- url       : The url you would this response to be returned from
- body      : The json response to return
- status    : The HTTP code to respond with

#### Example
```
{
    "url": "/api/servers/test",
    "body": {
        "serverId": "test",
        "supportedServiceTypes": ["live_encoding"],
        "tags": [],
        "timeout": 10
    },
    "status": 200
}
```

### [POST,PUT,GET,DELETE] /<any_endpoint>
-------------------------------------
Anything other than the `/register` endpoint will be matched against an internal map of urls. 
If that url has been previously registered, it will return the registered response, else it 
will return a `404`
