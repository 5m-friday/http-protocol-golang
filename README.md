# HTTP Protocol in Go

### What it is?

HTTP is just a communication protocol over TCP. In other words, just one of the ways
in which clients communicate with servers via some specially organized string.

### What it is NOT?

Some fancy alien technology which requires vast knowledge of quantum physics :D. Anyone can easily implement
the HTTP protocol either as a client or as a server side.

### RFC (Protocol)

For HTTP to be valid and acceptable for any client/server it should have the following:

#### HTTP MESSAGE STRUCTURE

`REQUEST/STATUS LINE\r\n`
<br>
`HEADER1\r\n`
<br>
`HEADER2\r\n`
<br>
`...\r\n`
<br>
`HEADERX\r\n`
<br>
`\r\n`
<br>
`OPTIONAL REQUEST/RESPONSE`

#### REQUEST LINE (REQUEST)
`HTTP VERB / METHOD` `␣` `URL/URI` `␣` `HTTP VERSION`
```
GET /some/uri HTTP/1.1
```

#### STATUS LINE (RESPONSE)
`HTTP VERSION` `␣` `STATUS CODE` `␣` `REASON PHRASE`
```
HTTP/1.1 200 OK
```

`Note:` Also in order for the HTTP request/response to be valid it also need to
have `Content-Length` header provided. In the request it equals to the length of the request
body, and should be present only if request body is present. In case of response should be equal
to the length of the response body if any.

##### Useful Links

- [RFC 7230 - HTTP Spec](https://tools.ietf.org/html/rfc7230)
- [RFC 7230 - Content Length](https://tools.ietf.org/html/rfc7230#section-3.3.2)
- [RFC 7231 - Request Methods](https://tools.ietf.org/html/rfc7231#section-4)
- [RFC 2616 - Status Codes](https://tools.ietf.org/html/rfc2616#section-10)

### HTTP Verbs

Generally speaking HTTP `Verbs` or `Methods` are ways to describe how
the client is going to have an effect on some resource located on the server,
like `GET` or `POST` for example.

Because HTTP is just a communication protocol, nobody restricts you to use specific methods aka verbs.
However, there are different verbs for different kinds of actions on resources on the server.

Here are the most widely used and known verbs:

- `GET` - typically used for getting resources
- `POST` - typically used for creating a resource or executing an action with side effects 
- `PATCH` - typically used for partially updating resources
- `PUT` - typically used for fully updating resources
- `DELETE` - typically used for deleting resources

However, there are other verbs/methods, such as:

`OPTIONS`, `CONNECT`, `TRACE`, `HEAD`

which we will not care to use, or rarely do so.

##### Useful Links

- [HTTP Verbs/Request Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)

### HTTP Statuses

HTTP Statuses are used for the server to communicate the client the result of its request,
meaning if the client made a request, it can generally know how the server responded and what's
the result of his action by simply getting back a `Status Code` in the response.

Status Codes usually come hand in hand with Reason Phrases. So a Status Code is made up of:
`NUMBER` `␣` `REASON PHRASE`, something like:

- `200 OK`
- `400 Bad Request`

There are many status codes, and you can even create your own `Custom Status Codes`, however generally
speaking you should know:

- `1.X.X` - Represents Informational statuses
- `2.X.X` - Represents Success statuses
- `3.X.X` - Represents Redirects statuses
- `4.X.X` - Represents Client Error statuses
- `5.X.X` - Represents Server Error statuses

##### Useful Links

- [HTTP Statuses/Response Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [HTTP Cat](https://http.cat/)

### HTTP Versions

- `HTTP 1 - Good old HTTP`
- `HTTP 2 - SPDY`
- `HTTP 3 - QUIC`

### HTTP/1.0 vs HTTP/1.1

##### HTTP/1.0

- `connection` is `closed` after every request/response (results in inefficiency)

##### HTTP/1.1

- requires `Host` header
- allows `persistent connections` (multiple request/response on the same connection)
- introduces the `OPTIONS` method mainly used for CORS
- has improved caching mechanism
- introduces `100` status code which lets the client know whether to proceed sending a large request body
- introduces chuncked transfer encoding
- introduces `Connection` header
- introduces enhanced compression support
- Much more (Generally stick with HTTP/1.1)

##### Useful Links

- [HTTP/1.0 vs HTTP/1.1](https://stackoverflow.com/questions/246859/http-1-0-vs-1-1#:~:text=HTTP%201.1%20also%20allows%20you,the%20connection%20would%20be%20closed.)
- [Differences between HTTP/1.0 & HTTP/1.1](http://www.ra.ethz.ch/cdstore/www8/data/2136/pdf/pd1.pdf)
- [Evolution of HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Evolution_of_HTTP)

### REST

Generally speaking REST is nothing, but a community effort & instruction around HTTP on how to
properly structure URLs and use HTTP Verbs & Statuses. So REST is the same old HTTP
used properly, with a little less freedom and at the end of the day it's a lot of debate
on how to properly use it, hence it's just a recommendation which most companies follow.

### HTTP Alternatives

- `Web Sockets`
- `HTTP2/HTTP3`
- `gRPC`
- `GraphQL`
