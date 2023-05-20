# hestia

Home projects. General area for mucking about with Golang using ideas from
articles on the internet.

# Preparation

## just

https://github.com/casey/just

which can be installed by executing:

```bash
export BIN="~/bin"
curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to ${BIN}
```

```bash
just
```

to see which commands are available.

## Shell setup

Create all local bin directories:

```bash
mkdir -p ~/bin
mkdir -p ~/go/bin
mkdir -p ~/.local/bin
mkdir -p ~/.local/go/bin
```

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

# set PATH so it includes user's go bin if it exists
if [ -d "$HOME/go/bin" ] ; then
    PATH="$HOME/go/bin:$PATH"
fi
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

[ -s "$HOME/.cargo/env" ] && \. "$HOME/.cargo/env"
```

## Reset environment

In order to make these settings permanent, logout or reboot your PC.

## Installing tools

First check for the presence of Go and other tools.

```bash
just tools which
```

Note: installation of Go is done locally in ~/.local/go as opposed to the normal
location of /usr/local/go. This is so sudo is not required. Uninstall Go in /usr/local/go
before attempting to install any tools.

Install tools like so:

```bash
just tools install
```

# Development workflow

## On a rebase

Initialise generated code and modules

```bash
just generate
just qa
```
## Changing code

Edit or add code or other development activity.

If any source files for generated code has changed:

```bash
just generate
```

Quality check code:

```bash
just qa
```

And build any new executables:

```bash
just build
```


