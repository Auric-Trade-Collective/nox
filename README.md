# Nox
Nox is a general purpose webserver written in Go. Providing
robust and simple ways for you to make your webapps come to
life, nox aims to live up to Nginx/Apache in a modern and 
complete way.

Features include:
- Serving directories and files
- Hosting APIs
- Using TLS
- Hosting quick and large scale web projects

# How to Build
Pull the current repository, and run ``go build``!

# How to Run
Since nox is in alpha versions, there will be many changes to
its configuration, however currently it needs a directory and
configuration file to run. Simply create a TOML file by running
`nox init` in a directory.

After this, you can run nox in the current directory, and it
will read your config and run the server when you run `nox spin`!

# Building an API
Right now Nox supports C first class through the native/webapi.h
file. But there are also official Python and Golang SDKs for the
native ABI.

- https://github.com/Auric-Trade-Collective/nox-go
- https://github.com/Auric-Trade-Collective/nox-py
