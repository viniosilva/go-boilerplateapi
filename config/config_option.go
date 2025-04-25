package config

type Option func(*client)

func WithPath(path string) Option {
	return func(c *client) {
		c.path = path
	}
}
