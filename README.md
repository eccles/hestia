# hestia

Home projects. General area for mucking about with Golang using ideas from
articles on the internet.

# Preparation

## Install make

```bash
sudo apt install make    # ubuntu
sudo dnf install make    # fedora
```

then type

```bash
make help
```

to see which commands are available.

### Installing tools

First check for the presence of Go and other tools.

```bash
make tools-check
```

Note: installation of Go is done locally in ~/.local/go as opposed to the normal
location of /usr/local/go. This is so sudo is not required. Uninstall Go in /usr/local/go
before attempting to install any tools.

Install tools like so:

```bash
make tools-install
```

### Shell setup

The file ~/.profile should contain the following lines to setup PATH correctly:

```bash
# set these in the correct order
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/bin" ] ; then
    PATH="$HOME/bin:$PATH"
fi

# set PATH so it includes user's private local bin if it exists
if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

# set PATH so it includes user's local go bin if it exists
if [ -d "$HOME/.local/go/bin" ] ; then
    PATH="$HOME/.local/go/bin:$PATH"
fi
```
