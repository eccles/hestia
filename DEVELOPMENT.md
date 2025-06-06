# Development environment

This repo uses asdf to manage tool versions and just as a 'make' replacement.
The use of asdf is not compulsory. If asdf is is not used then manually install
'just' from https://github.com/casey/just

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

Initialise tools and modules

```bash
just tools
just qa
```
## Changing code

Edit or add code or other development activity.

Quality check code:

```bash
just qa
```

