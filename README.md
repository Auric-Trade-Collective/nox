# Nox
Nox is a general purpose webserver written in Go. Providing
robust and simple ways for you to host your websites and web 
APIs, Nox is a powerful replacement to tools like Caddy, Nginx, 
or Apache.

**Supported Platforms**
- Linux: Nox has complete and first class support for Linux, all
testing and development happens natively for it.
- MacOS: Nox reasonably has support for MacOS, there's no testing
or builds officially made for it, but as of now nox is known to build
and run on MacOS.
- Windows: There is no Nox support for Windows. Furthermore, Nox
not plan to add support for Windows. The best way to run Nox if
you are on a Windows machine is to use WSL.

# How to Build
Pull the current repository, and run ``go build``! You notably
will need gcc installed.

# How to Run
Simply create a TOML file by running `nox init` in a directory.

After this, you can run nox in the current directory, it
will read your config and run the server when you run `nox spin`!

# Building an API
Nox has support for hosting APIs, this is done through a native
ABI.

Right now Nox supports C first class through the include/nox.h
file. But there are also official SDKs for other languages that
wrap the ABI.

- https://github.com/Auric-Trade-Collective/nox-d
- https://github.com/Auric-Trade-Collective/nox-go [Discontinued]
- https://github.com/Auric-Trade-Collective/nox-py
- Flux wrapper coming!
