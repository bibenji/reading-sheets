```
go build -o {name}
./{name}
```

```
go build -o chat
./chat
```

# 1 Chat Application with Web Sockets

go get github.com/gorilla/websocket

P. 34

P. 39

go run main go -addr=":3000"

Tracing code to get a look under the hood

P. 44

`Trace(...interface{}) => take multiple args`

`go test`

`go test -cover`

Clean package APIs

# 2 Adding User Accounts

P. 57

for advanced routing situations, check: goweb, pat, routes or mux

Getting started with OAuth2

P. 66

1. user selects the provider
2. user is redirected to provider's website (with a URL that includes the client app ID) where they are asked to give permission to the client app
3. user signs in from the OAuth2 service provider and accepts the permissions requested by the third-party application
4. user is redirected to the client app with a request code
5. in the background, the client app sends the grant code to the provider, who sends back an authentication token
6. client app uses access token to make authorized requests to provider, such as to get user information or wall posts

check golang/oauth2 or stretchr/gomniauth 

P. 69

Try a watcher to work better...
canthefason / go-watcher

[...]

# 3 Three Ways to Implement Profile Picture

P. 81

Refactoring and optimizing our code

P. 106

Tidying up and testing

P. 116

# 4 Command-Line Tools to Find Domain Names

P. 117
 
redirection metacharacters

echo -n "Hello" | md5

create the programs inside the $GOPATH/src directory

use the current time as a random seed.

Computers can't actually generate random numbers, but changing the seed number of random algorithms gives the illusion that it can

bufio.ScanWords

crypto/rand

`go build -o sprinkle ./sprinkle`

echo "chat" | ./sprinkle

TLD = top level domain

./synonyms/synonyms | ./sprinkle/sprinkle | ./coolify/coolify | ./domainify/domainify 

Available

P. 134

./synonyms/synonyms | ./sprinkle/sprinkle | ./coolify/coolify | ./domainify/domainify | ./available/available

exec.Command

and cmd.Start (run in the background instead of cmd.Run)

# 5 Building Distributed Systems and Working with Flexible Data

P. 144

Horizontal scaling
Vertical scaling

nsqlookupd => manages topology information about the distributed NSQ environment, keeps track of all the nsqd producers for specific topics and provides interfaces for clients to query such information

nsqd => does the heavy lifting for NSQ, receiving, queuing, delivering messages from and to interested parties

Reading from MongoDB

P. 156

a lot of things to review for oauth in twitter.go
and wait for validation of twitter's dev account

Signal channels

P. 160

deadlines, cancelation and stopping for cross API boundaries (P. 162)

you can close a channel to mean things are done or send something in it

Look: A Channel Compendium by John Graham-Cumming

p. 164

pour suivre en direct the messagin queue topic and outputs any messages that it notices :

`nsq_tail --topic="votes" --lookupd-http-address=localhost:4161`

```
use ballots
db.polls.find().pretty()
```

# 6 Exposing Data and Functionality through a RESTful Data Web Service API

P. 177

use things like that for CORS : https://github.com/fasterness/cors

DRY: Don't Repeat Yourself

For path parsing, looks Goweb or Gorrillz's mux package

To add polls in the database :
```
use ballots
db.polls.insert({"title":"Test poll","options":["one","two","three"]})
db.polls.insert({"title":"Test poll two","options":["kiwi","pomme","poire"]})
```

http://localhost:8070/polls/?key=abc123

P. 193

Test with curl:
```
curl -X GET http://localhost:8070/polls/?key=abc123
curl --data '{"title":"test","options":["one","two","three"]}' -X POST http://localhost:8070/polls/?key=abc123
curl -X GET http://localhost:8070/polls/{id}?key=abc123
curl -X DELETE http://localhost:8070/polls/{id}?key=abc123
```

A web client that consumes the API

7 Random Recommendations Web Service

P. 209

Public views of Go structs

P. 216

can use tags to control fiel names and implement your own []string type that provides a MarshalJSON method which tells the encoder how to marshal your type

Enumerators in Go:
- define a new type based on a primitive integer type
- use that type whenever you need users to specify one of the appropriate values
- use the iota keyword to set the values in a const block, disregarding the first zero value
- implement a map of sensible string representations to the values of your ennumerator
- implement a String method on the type that returns the appropriate string representation from the map
- implement a ParseType function that converts from a string to your type using the map

Querying the Google Places API

P. 225

Building recommendations

P.226

+ tester requête à Google API

Go to mygeoposition.com

http://localhost:8040/recommendations?lat=51.520707&lng=-0.153809&radius=5000&journey=cafe|bar|casino|restaurant&cost=$...$$$

Web application

P. 232

# 8 Files System Backup

P. 234

See fsnotify for go

Checking for changes and initiating a backup

P. 242

The tool should be used the following ways:
```
backup -db=/path/to/db add {path} [paths...]
backup -db=/path/to/db remove {path} [paths...]
backup -db=/path/to/db list
```

The daemon backup tool

P. 251

How to use backupd:

`./backupd -db="db_path" -archive="archive_path" - interval=5s`

# 9 Building a Q&A Application for Google App Engine

P. 259

https://cloud.google.com/appengine/downloads

https://console.cloud.google.com

```
goapp serve
goapp deploy
```

Transactions in Google Cloud Datastore

P. 278

- Align the happy path to the left edge so that you can scan down in a single column and see the expected flow of execution
- Don't hide the happy path logic inside a nest of indented braces
- Exit early from your function
- Indent only to handle errors or edge cases
- Extract functions and methods to keep bodies small and readable

```
if obj, ok := v.(interface{ OK() error }); ok {
	// v has OK() method
} else {
	// v does not have OK() method
}
```

but not so good because it hides the secret functionnality from users of the code
so you must either document the function very well or perhaps promote the method to its own first-class interface and insist that all objects implement it
always seek clear code over clever code

P. 301

`goapp serve dispatch.yaml default/app.yaml api/app.yaml web/app.yaml`

`goapp deploy default/app.yaml api/app.yaml web/app.yaml`

then use : appcfg.py

`appcfg.py update_dispatch .`

`appcfg.py update_indexes -A YOUR_APPLICATION_ID_HERE ./default`

# 10 Micro-services in Go with the Go kit Framework

P. 310

https://gokit.io/

to solve SOA (service-oriented architecture) problems, such as service discovery, metrics, monitoring, logging, load balancing, circuit breaking, and others

gRPC
protocol buffers

protoc vault.proto --go_out=plugins=grpc:.

or

```
docker run -v `pwd`:/defs namely/protoc-all -f vault.proto -l go

docker run --rm -v `pwd`:/defs -v `pwd`:/output namely/protoc-all -f vault.proto -o /output -l go
docker run -v ${pwd}:/defs namely/protoc-all -f vault.proto -l go
```

Building the service

P. 317

!!!:

Note that the receiver in the Hash method is just (vaulService); we don't capture the variable because there is no way we can store state on an empty struct.

An HTTP server in Go kit

P. 327

Creating a server command

P. 331

Add TLS (Transport Layer Security) for every service

`go run main.go`

```
curl -XPOST -d '{"password":"hernandez"}' http://localhost:8080/hash

curl -XPOST -d '{"password":"hernandez", "hash":"THE_HASH_HERE"}' http://localhost:8080/validate
```

#### Building a gRPC client

P. 337

`vaultcli hash MyPassword`

`vaultcli hash MyPassword HASH_GOES_HERE`

to build a script and put it in your path: `go install`

inside cmd/vaultd:  `go run main.go`

#### Rate limiting with service middleware

The token bucket is an algorithm used in packet switched computer networks and telecommunications networks.

The general idea: we have a bucket of tokens, each request need a token, if no tokens, we have reached our limit. Buckets refill over time at a specific interval.

```
e := getEndpoint(srv)
{
	e = getSomeMiddleware()(e)
	e = getLoggingMiddleware(logger)(e)
	e = getAnotherMiddleware(something)(e)
}
```

`go run main.go`

`vautcli hash bourgon`

11 Deploying Go Applications Using Docker

P. 349

Building Go binaries for different architectures

`CGO_ENABLED=0 GOOS=linux go build -a ./cmd/vaultd/`

`docker build -t vaultd`

`docker run -p 6060:8080 -p 6061:8081 --name localtest --rm vaultd`

Deploying to Docker Hub [...]

Deploying to Digital Ocean [...]

Platform as a Service (PaaS)

Creating a droplet [...]

### Good Practices for a Stable Go Environment

GOPATH can contain a list of colon-separated folders, you can even have a different value for GOPATH depending on which project you are working
but strongly recommended that you use a single GOPATH location for everything

`go fmt -w`

`go vet`

go get golang.org/x/tools/cmd/goimports

goimports -w *.go

Finish.
