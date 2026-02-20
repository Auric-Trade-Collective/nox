package webserver

type Config struct {
	Nox struct {
		Addr string `toml:"addr"`
		Root string	`toml:"root"`
		Api string	`toml:"api"`
		Tls struct {
			Enabled bool `toml:"enabled"`
			CertFile string `toml:"cert_file"`
			KeyFile string	`toml:"key_file"`
			Ciphers string `toml:"ciphers"`
		} `toml:"tls"`
		Dependencies map[string]string `toml:"dependencies"`
	} `toml:"nox"`
}
