# hestia

Home projects. General area for mucking about with Golang using ideas from
articles on the internet.

# Preparation

## asdf

https://asdf-vm.com/guide/getting-started.html

Install asdf following the above instructions
and ensure the following is in ~/.bashrc:

```bash
export PATH="${ASDF_DATA_DIR:-$HOME/.asdf}/shims:$PATH"
. <(asdf completion bash)
```

## Reset environment

In order to make these settings permanent, logout or reboot your PC.

## Installing tools

This wil install go and other tools including 'just'.

```bash
./scripts/asdf-install.sh
```

and then install go tools:

```bash
just tools
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


