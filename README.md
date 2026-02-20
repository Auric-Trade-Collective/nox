# Nox
Nox is a general purpose webserver written in Go. Providing
robust and simple ways for you to make your webapps come to
life, nox aims to live up to Nginx/Apache in a modern and 
complete way.

# How to Build
Pull the current repository, and run ``go build``!

# How to Run
Since nox is in alpha versions, there will be many changes to
its configuration, however currently it needs a directory and
configuration file to run. Simply create a TOML file like so:

```toml
[nox]
addr = ":5432"
root = "/path/to/my/dir"
api = "/path/to/my/api"

# TLS not required, and not yet supported
[nox.tls]
enabled = true
cert_file = "/some/dir/file.crt"
key_file = "/some/dir/file.key"
ciphers = "N/A"
```
