# Development environment

This repo uses just as a 'make' replacement.

Install go tools:

```bash
just tools
```

# Development workflow

## On a rebase

Initialise tools and modules

```bash
just tools
just generate
just qa
```
## Changing code

Edit or add code or other development activity.

Quality check code:

```bash
just generate
just qa
just unittest
```

