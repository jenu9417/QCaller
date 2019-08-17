package as

// Conf : holds the params needed to initialize AS client
type Conf struct {
	host string
	port int
}

// NewConf : initializes and returns Conf struct
func NewConf(host string, port int) *Conf {
	return &Conf{
		host: host,
		port: port,
	}
}

// Host : returns the AS host
func (c *Conf) Host() string {
	return c.host
}

// Port : returns the AS port
func (c *Conf) Port() int {
	return c.port
}
