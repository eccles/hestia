# zed editor

Installing zed on my Pop OS laptop following this link:

https://zed.dev/download

and using this script:

```bash
set -euxo pipefail
cd ~/Downloads
mkdir -p zed
cd zed
curl -fO https://zed.dev/install.sh
chmod +x install.sh
./install.sh
```
led to an error when launching zed.

```
This Zed opens with a message "Unsupported GPU" and offers "Troubleshoot and Quit" or "Skip"
```

Investigating showed these links:

https://github.com/zed-industries/zed/issues/22864
https://github.com/zed-industries/zed/issues/22978

Editing /etc/prime-discrete and setting 'on' to 'off' fixed the problem.

This error seems to be specific to Pop OS.

# Pop OS 24.04

This is an experimental COSMIC desktop. Setting the /etc/prime-discrete value as
described above led to system76-power dying and the NIC losing its IP address.

On this platform setting

```
export ZED_ALLOW_EMULATED_GPU=1
```

in .bashrc worked.

See https://github.com/pop-os/cosmic-epoch/issues/1976.
