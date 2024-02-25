# Konsole: Web Console Library for Bedrock Servers
Konsole is a web console library designed specifically for Bedrock servers, providing a convenient interface for server administration and monitoring through a web browser.

## Features
Web Console: Access your Bedrock server's console directly from a web browser.
Real-time Updates: Receive real-time updates and logs from the server.
WebSockets Support: Utilize WebSockets for efficient communication between the server and the web interface.

## Getting Started
Hosting the Web Server
To host the web server for Konsole, follow these steps:

## Import the Konsole library in your Go project.

```go
a := app.New("/your_secret_endpoint_here")
```
## Start the web server, specifying the desired port.

```go
err := a.ListenAndServe(":6969")
if err != nil {
    panic(err)
}
```

## Starting the WebSocket Server
If you're using Dragofly, proceed with the following steps:

```go
ws := konsole.NewWebSocketServer(chat.StdoutSubscriber{}, "test", testFormatter{})
```
## Start the WebSocket server, specifying the desired port.

```go
err := ws.ListenAndServe(":8080")
if err != nil {
    panic(err)
}
```
Customization
Konsole allows for customization of behavior, and functionality as needed.
