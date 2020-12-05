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

3 Three Ways to Implement Profile Picture

P. 81

Refactoring and optimizing our code

P. 106

 