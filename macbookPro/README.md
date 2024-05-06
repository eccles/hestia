# Installing Linux on Macbook Pro

I have recently acquired a Macbook Pro vintage 2014 and installed Pop-OS on it.
Unfortunately the wifi did not activate. After much googling and encountering many posts on 
this issue that were confusing:

- overcomplicated - one very helpful dev had the original OP recompile the kernel to try and solve this.
- confusion about which software tools to install and what the actual cause was.

# Cause

The macBook Pro used a broadcom bcm4360 network chip which is incompatible with the standard broadcom linux 
drivers.

Executing the following will show the problem - using bcma-pci-bridge kernel driver:

```bash
sudo lspci -vnn -d 14e4:
...
03:00.0 Network controller [0280]: Broadcom Inc. and subsidiaries BCM4360 802.11ac Wireless Network Adapter [14e4:43a0] (rev 03)
Subsystem: Apple Inc. BCM4360 802.11ac Wireless Network Adapter [106b:0112]
Flags: bus master, fast devsel, latency 0, IRQ 18, IOMMU group 13
Memory at b0600000 (64-bit, non-prefetchable) [size=32K]
Memory at b0400000 (64-bit, non-prefetchable) [size=2M]
Capabilities: [48] Power Management version 3
Capabilities: [58] MSI: Enable- Count=1/1 Maskable- 64bit+
Capabilities: [68] Vendor Specific Information: Len=44 <?>
Capabilities: [ac] Express Endpoint, MSI 00
Capabilities: [100] Advanced Error Reporting
Capabilities: [13c] Device Serial Number 60-f8-00-ff-ff-00-00-01
Capabilities: [150] Power Budgeting <?>
Capabilities: [160] Virtual Channel
Capabilities: [1b0] Latency Tolerance Reporting
Capabilities: [220] Physical Resizable BAR
Kernel driver in use: bcma-pci-bridge
Kernel modules: bcma
```

```
** NOTE **
There are 2 variants: 14e4:4360 and 14e4:43a0
```

# Solution

Install the non-free drivers (Pop OS has the non-free repo enabled by default):

```bash
sudo apt-get install broadcom-sta-dkms
```

This should blacklist the bcma and b43 drivers. Inspect /etc/modprobe.d to find the appropriate statements.

The initramfs must be rebuilt:

```bash
sudo update-initramfs -u
```

Reboot and the wifi setup should be available in the settings app.
Reinspecting the pci bus should show( note the driver has changed to 'wl' and not 'bcma-pci-bridge'):

```bash
sudo lspci -vnn -d 14e4:
.....
03:00.0 Network controller [0280]: Broadcom Inc. and subsidiaries BCM4360 802.11ac Wireless Network Adapter [14e4:43a0] (rev 03)
Subsystem: Apple Inc. BCM4360 802.11ac Wireless Network Adapter [106b:0112]
Flags: bus master, fast devsel, latency 0, IRQ 18, IOMMU group 13
Memory at b0600000 (64-bit, non-prefetchable) [size=32K]
Memory at b0400000 (64-bit, non-prefetchable) [size=2M]
Capabilities: [48] Power Management version 3
Capabilities: [58] MSI: Enable- Count=1/1 Maskable- 64bit+
Capabilities: [68] Vendor Specific Information: Len=44 <?>
Capabilities: [ac] Express Endpoint, MSI 00
Capabilities: [100] Advanced Error Reporting
Capabilities: [13c] Device Serial Number 60-f8-00-ff-ff-00-00-01
Capabilities: [150] Power Budgeting <?>
Capabilities: [160] Virtual Channel
Capabilities: [1b0] Latency Tolerance Reporting
Capabilities: [220] Physical Resizable BAR
Kernel driver in use: wl
Kernel modules: bcma, wl
```
