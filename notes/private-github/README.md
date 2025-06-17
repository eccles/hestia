Accessing private github golang repos
-------------------------------------

If you want to access github private repos without hassle then please try the following:

install git-credential-manager:  https://github.com/git-ecosystem/git-credential-manager

  - extract the executable and copy to your path

    The tarball contains:

```
        git-credential-manager  libHarfBuzzSharp.so  libSkiaSharp.so  NOTICE
```

```bash
        mkdir ~/bin/gcm
        cp * ~/bin/gcm
        chmod +x ~/bin/gcm/git-credential-maneger
```

  - configure:

```bash
        ~/bin/gcm/git-credential-manager configure
```

  - check ~/.gitconfig:

```
        [credential]
        helper = 
        helper = /home/eccles/bin/git-credential-manager
```

  - also delete any 'insteadOf entries - they are not required.

  - now check out a repo using HTTPS. This is required because that is what 'go get' uses.

```bash
        # there are more options available - lets just use RAM...
        export GCM_CREDENTIAL_STORE=cache 
        git clone https://github.com/eccles/hestia.git
```

    A browser window should appear - choose accordingly.

> [!NOTE]
> This is an official Microsoft product which is available for all platforms.


