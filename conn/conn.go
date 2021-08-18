package conn

// ConnType is the connection type
// that will be set between this
// bot and Twitch
type ConnType int

const (
	// IRCType defines the IRC connection
	// between the bot and Twitch's chat
	IRCType ConnType = iota
	// WSType defines the WebSocket connection
	// between the bot and Twitch's chat
	WSType
)

// Conn interface represents
// all action related to the
// connections types
type Conn interface {
	FetchAll()
}

// FetchAll fetch each message
func FetchAll(c Conn) {
	c.FetchAll()
}
