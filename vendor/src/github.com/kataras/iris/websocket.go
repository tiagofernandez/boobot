package iris

import (
	irisWebsocket "github.com/iris-contrib/websocket"
	"github.com/kataras/go-websocket"
)

// conversionals
const (
	// All is the string which the Emmiter use to send a message to all
	All = websocket.All
	// NotMe is the string which the Emmiter use to send a message to all except this websocket.Connection
	NotMe = websocket.NotMe
	// Broadcast is the string which the Emmiter use to send a message to all except this websocket.Connection, same as 'NotMe'
	Broadcast = websocket.Broadcast
)

// newUnderlineWsServer returns a new go-websocket.Server from configuration, used internaly by Iris.
func newUnderlineWsServer(c WebsocketConfiguration) websocket.Server {
	wsConfig := websocket.Config{
		WriteTimeout:    c.WriteTimeout,
		PongTimeout:     c.PongTimeout,
		PingPeriod:      c.PingPeriod,
		MaxMessageSize:  c.MaxMessageSize,
		BinaryMessages:  c.BinaryMessages,
		ReadBufferSize:  c.ReadBufferSize,
		WriteBufferSize: c.WriteBufferSize,
	}

	return websocket.New(wsConfig)
}

// Note I keep this code only to no change the front-end API, we could only use the go-websocket and set our custom upgrader

type (
	// WebsocketServer is the iris websocket server, expose the websocket.Server
	// the below code is a wrapper and bridge between iris-contrib/websocket and kataras/go-websocket
	WebsocketServer struct {
		websocket.Server
		config   WebsocketConfiguration
		upgrader irisWebsocket.Upgrader
	}
)

// NewWebsocketServer returns an empty WebsocketServer, nothing special here.
func NewWebsocketServer() *WebsocketServer {
	return &WebsocketServer{}
}

// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
//
// The responseHeader is included in the response to the client's upgrade
// request. Use the responseHeader to specify cookies (Set-Cookie) and the
// application negotiated subprotocol (Sec-Websocket-Protocol).
//
// If the upgrade fails, then Upgrade replies to the client with an HTTP error
// response.
func (s *WebsocketServer) Upgrade(ctx *Context) error {
	return s.upgrader.Upgrade(ctx)
}

// Handler is the iris Handler to upgrade the request
// used inside RegisterRoutes
func (s *WebsocketServer) Handler(ctx *Context) {
	if err := s.Upgrade(ctx); err != nil {
		if ctx.framework.Config.IsDevelopment {
			ctx.Log("Websocket error while trying to Upgrade the connection. Trace: %s", err.Error())
		}
		ctx.EmitError(StatusBadRequest)
	}
}

// RegisterTo creates the client side source route and the route path Endpoint with the correct Handler
// receives the websocket configuration and  the iris station
func (s *WebsocketServer) RegisterTo(station *Framework, c WebsocketConfiguration) {
	s.config = c // save the configuration, we will need that on the .OnConnection
	clientSideLookupName := "iris-websocket-client-side"
	station.Get(s.config.Endpoint, s.Handler)
	// check if client side already exists
	if station.Lookup(clientSideLookupName) == nil {
		// serve the client side on domain:port/iris-ws.js
		station.StaticContent("/iris-ws.js", contentJavascript, websocket.ClientSource)(clientSideLookupName)
	}
}

// WebsocketConnection is the front-end API that you will use to communicate with the client side
type WebsocketConnection interface {
	websocket.Connection
}

// OnConnection this is the main event you, as developer, will work with each of the websocket connections
func (s *WebsocketServer) OnConnection(connectionListener func(WebsocketConnection)) {
	// let's initialize here the ws server, the user/dev is free to change its config before this step.
	if s.Server == nil {
		s.Server = newUnderlineWsServer(s.config)
	}

	if s.upgrader.Receiver == nil {
		s.upgrader = irisWebsocket.Custom(s.Server.HandleConnection, s.config.ReadBufferSize, s.config.WriteBufferSize, false)
		// run the ws server
		s.Server.Serve()
	}

	s.Server.OnConnection(func(c websocket.Connection) {
		connectionListener(c)
	})
}
