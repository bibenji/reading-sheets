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
```asp
use ballots
db.polls.insert({"title":"Test poll","options":["one","two","three"]})
db.polls.insert({"title":"Test poll two","options":["kiwi","pomme","poire"]})
```

http://localhost:8070/polls/?key=abc123

P. 193

Test with curl:
```asp
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
