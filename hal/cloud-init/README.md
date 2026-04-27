# Cloud-Init And FDE Workflow

This directory contains the host-side workflow for preparing an Ubuntu `qcow2`
image with `cloud-init`, plus the Buildroot initramfs used by the current full
disk encryption flow.

The flow has two main phases:

1. Prepare a source Ubuntu image with `cloud-init`.
2. Boot a small Buildroot initramfs that copies the source image into a fresh
   LUKS2 destination disk, measures the copied disk, and switches into the
   encrypted root filesystem.

The `cloud-init` scripts prepare the guest OS image. The actual FDE work happens
in the Buildroot early userspace init script at
[buildroot/board/rootfs-overlay/init](./buildroot/board/rootfs-overlay/init).

## Files

- [package-services.yaml](./package-services.yaml): cloud-init user-data.
- [create-seed-iso.sh](./create-seed-iso.sh): renders the user-data and builds a
  NoCloud `seed.iso`.
- [qemu.sh](./qemu.sh): creates `seed.iso`, recreates the writable qcow2 overlay,
  and launches the preparation VM.
- [../cloud/.env](../cloud/.env): shared defaults currently reused by the
  `cloud-init` scripts.
- [../cloud/meta-data](../cloud/meta-data): shared NoCloud meta-data currently
  reused by the `cloud-init` scripts.
- [buildroot](./buildroot): Buildroot external tree used for the FDE initramfs.
- [buildroot/configs/cocos_defconfig](./buildroot/configs/cocos_defconfig):
  builds the FDE initramfs with NBD, `cryptsetup`, `kpartx`, `udev`, TPM2 tools,
  and `tdx-rtmr-extend`.
- [buildroot/package/tdx-rtmr-extend](./buildroot/package/tdx-rtmr-extend):
  helper used to extend TDX RTMR2 or RTMR3 with a SHA-384 digest.
- [../../scripts/nbd/src_start_nbd.sh](../../scripts/nbd/src_start_nbd.sh):
  external helper that exports a qcow2 source image over NBD as export name
  `src`.

## Requirements

Install these on the host:

- `qemu-system-x86_64` or whatever binary is set in the shared [`.env`](../cloud/.env)
- `qemu-img`
- `wget`
- one of `xorriso`, `genisoimage`, or `mkisofs`

This workflow does not have its own `.env` file under `hal/cloud-init`. Today,
both [qemu.sh](./qemu.sh) and [create-seed-iso.sh](./create-seed-iso.sh) reuse
the shared defaults from [../cloud/.env](../cloud/.env), and the default
NoCloud meta-data comes from [../cloud/meta-data](../cloud/meta-data).

For the FDE source-image export flow, the host also needs QEMU NBD support,
usually provided by `qemu-utils` or `qemu-tools`.

## Create The Seed ISO

To create `seed.iso` only:

```bash
cd ./cocos/hal/cloud-init
./create-seed-iso.sh
```

To write the ISO somewhere else:

```bash
./create-seed-iso.sh ./out/seed.iso
```

Supported environment overrides for [create-seed-iso.sh](./create-seed-iso.sh):

- `USER_DATA_SOURCE`
- `META_DATA_SOURCE`
- `NETWORK_CONFIG_SOURCE`

Package toggles for the `hal/linux/package` package set are also supported:

- `COCOS_INSTALL_AGENT`
- `COCOS_INSTALL_ATTESTATION_SERVICE`
- `COCOS_INSTALL_CC_ATTESTATION_AGENT`
- `COCOS_INSTALL_COCO_KEYPROVIDER`
- `COCOS_INSTALL_COMPUTATION_RUNNER`
- `COCOS_INSTALL_EGRESS_PROXY`
- `COCOS_INSTALL_INGRESS_PROXY`
- `COCOS_INSTALL_LOG_FORWARDER`
- `COCOS_INSTALL_WASMEDGE`

Defaults:

- `COCOS_INSTALL_CC_ATTESTATION_AGENT=false`
- `COCOS_INSTALL_COCO_KEYPROVIDER=false`
- all other `COCOS_INSTALL_*` toggles default to `true`

Runtime dependencies are resolved automatically during provisioning:

- `COCOS_INSTALL_AGENT=true` also enables `attestation-service`, `log-forwarder`, `computation-runner`, `ingress-proxy`, and `egress-proxy`
- `COCOS_INSTALL_COMPUTATION_RUNNER=true` also enables `log-forwarder`
- `COCOS_INSTALL_COCO_KEYPROVIDER=true` also enables `cc-attestation-agent`

Example:

```bash
USER_DATA_SOURCE=./package-services.yaml ./create-seed-iso.sh
```

Example with package selection:

```bash
COCOS_INSTALL_AGENT=false COCOS_INSTALL_WASMEDGE=false ./create-seed-iso.sh
```

## Boot The Prep VM

To create the seed ISO and boot the Ubuntu prep VM:

```bash
cd ./cocos/hal/cloud-init
sudo ./qemu.sh
```

By default, [qemu.sh](./qemu.sh) downloads the Ubuntu Noble cloud image from:

- <https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img>

To download the image manually with `curl`:

```bash
cd ./cocos/hal/cloud-init
curl -L https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img \
  -o noble-server-cloudimg-amd64.img
```

To use an Ubuntu cloud image that is already downloaded, point
`BASE_IMAGE_PATH` at the existing image file. When that file exists,
[qemu.sh](./qemu.sh) skips the download step:

```bash
cd ./cocos/hal/cloud-init
sudo BASE_IMAGE_PATH="$PWD/noble-server-cloudimg-amd64.img" \
  CUSTOM_IMAGE_PATH="$PWD/noble-cocos-prep.qcow2" \
  ./qemu.sh
```

What [qemu.sh](./qemu.sh) does:

- loads defaults from the shared [`.env`](../cloud/.env) in `hal/cloud`
- downloads the Ubuntu Noble cloud image if it is missing
- recreates `seed.iso`
- deletes and recreates the writable qcow2 overlay image
- boots QEMU with the seed ISO attached as a CD-ROM and the writable qcow2 attached as the VM disk

Important:

- [qemu.sh](./qemu.sh) must be run as `root`
- the writable qcow2 at `CUSTOM_IMAGE_PATH` is removed and recreated on each run
- the VM powers itself off after `cloud-init` finishes, so the altered image is left on disk for later use

Useful environment overrides for [qemu.sh](./qemu.sh):

- `SEED_ISO`
- `META_DATA`
- `BASE_IMAGE_PATH`
- `CUSTOM_IMAGE_PATH`
- `OVMF_FILE`
- all `COCOS_INSTALL_*` package toggles listed above

Example:

```bash
sudo BASE_IMAGE_PATH=./noble-server-cloudimg-amd64.img CUSTOM_IMAGE_PATH=./noble-custom.qcow2 ./qemu.sh
```

Example with package selection:

```bash
sudo COCOS_INSTALL_AGENT=false COCOS_INSTALL_WASMEDGE=false ./qemu.sh
```

Example with a single firmware file:

```bash
sudo OVMF_FILE=./OVMF.fd BASE_IMAGE_PATH=./noble-server-cloudimg-amd64.img ./qemu.sh
```

## What The Guest Configures

On first boot, [package-services.yaml](./package-services.yaml) will:

- grow the root partition and filesystem to use extra disk space when the qcow2 is larger than the base image
- install build dependencies with `apt`
- build and install the Cocos package-defined services
- optionally build and install `attestation-agent` and `coco_keyprovider`
- optionally install WasmEdge
- configure `/etc/ocicrypt_keyprovider.conf`
- prepare 9P mountpoints and `/etc/fstab` entries for `certs_share` and `env_share`
- enable and restart the configured systemd services
- power the VM off when provisioning is complete

The package install list can be reduced with the `COCOS_INSTALL_*` environment variables when creating the seed ISO or running [qemu.sh](./qemu.sh).

With the default package selection, the configured services are:

- `cocos-agent`
- `log-forwarder`
- `computation-runner`
- `egress-proxy`
- `attestation-service`

Optional services when enabled are:

- `attestation-agent`
- `coco-keyprovider`

## 9P Note

The altered Ubuntu image is prepared to use 9P mounts through `/etc/fstab`, but
[qemu.sh](./qemu.sh) only prepares the image. The final runtime launcher must
attach the `certs_share` and `env_share` 9P devices if those mounts are needed.

## FDE Source Image

The prepared qcow2 is the trusted source image that the FDE initramfs later
clones into an encrypted destination disk.

Before booting the FDE initramfs, the source qcow2 must be reachable over NBD.
The helper in
[../../scripts/nbd/src_start_nbd.sh](../../scripts/nbd/src_start_nbd.sh)
currently does this with:

```bash
sudo qemu-nbd --read-only --persistent --fork --export-name=src --port="$SRC_PORT" -f qcow2 "$SRC_QCOW"
```

Important details:

- the export name is `src`;
- the default source port is `10809`;
- the export is read-only;
- the initramfs expects the source endpoint on the kernel command line as
  `src_ip=<host-or-service-ip>` and optionally `src_port=<port>`.

If `src_ip` is missing, the initramfs stops and drops to a shell.

## FDE Initramfs

The Buildroot image is built from
[buildroot/configs/cocos_defconfig](./buildroot/configs/cocos_defconfig). The
configuration produces a compressed CPIO root filesystem and includes the tools
needed by the early FDE script:

- `nbd-client`
- `cryptsetup`
- `kpartx` from multipath tools
- `udev`
- `tpm2_pcrextend`
- `tdx-rtmr-extend`
- core utilities for hashing, copying, and key cleanup

At runtime, the current init script assumes these block devices:

- source image: `/dev/nbd2`
- destination disk: `/dev/sda`
- opened LUKS mapping: `/dev/mapper/sda_crypt`
- expected root partition: `/dev/mapper/sda_cryptp1` or
  `/dev/mapper/sda_crypt1`

The destination disk is formatted. Any previous contents of `/dev/sda` are
destroyed. The destination must be large enough to hold the source image plus
LUKS metadata overhead.

## Early Init FDE Sequence

[buildroot/board/rootfs-overlay/init](./buildroot/board/rootfs-overlay/init)
runs as PID 1 and performs these steps:

1. Sets up console I/O, `/dev`, `/proc`, `/sys`, `/run`, `/tmp`, devtmpfs,
   devpts, and configfs.
2. Parses `src_ip` and `src_port` from `/proc/cmdline`; `src_port` defaults to
   `10809`.
3. Acquires networking with `udhcpc`.
4. Connects to the source NBD export:

   ```bash
   nbd-client -N src "$SRC_IP" "$SRC_PORT" /dev/nbd2
   ```

5. Generates a 64-byte random key at `kk.bin`.
6. Creates a LUKS2 container on `/dev/sda` using `cryptsetup luksFormat` with
   `--cipher aes-xts-plain64`.
7. Opens the container as `/dev/mapper/sda_crypt` using the generated key.
8. Copies the entire source block device into the encrypted mapper:

   ```bash
   dd if=/dev/nbd2 bs=16M iflag=fullblock of=/dev/mapper/sda_crypt oflag=direct conv=fsync
   ```

9. Computes SHA-256 and SHA-384 by reading the opened plaintext mapper at
   `/dev/mapper/sda_crypt` after the source copy completes.
10. Disconnects the NBD source.
11. Extends attestation measurements when the relevant device support is
    present:
    - vTPM: extends PCR16 with the SHA-256 and SHA-384 disk hashes through
      `tpm2_pcrextend`;
    - TDX: extends RTMR3 with the SHA-384 disk hash through
      `tdx-rtmr-extend -rtmr 3 -sha384`.
12. Runs `kpartx -av /dev/mapper/sda_crypt` to expose partitions inside the
    encrypted disk.
13. Waits for devices to settle, preferring `udevadm settle`.
14. Mounts the first root partition as ext4 at `/root`.
15. Moves `/proc` and `/sys` into the new root.
16. Replaces the cloned image's `/etc/fstab` with:
    - the detected encrypted root partition mounted at `/`;
    - preserved or default 9P entries for `certs_share` at `/etc/certs`;
    - preserved or default 9P entries for `env_share` at `/etc/cocos`.
17. Wipes `kk.bin` with `shred` when available, falling back to overwriting the
    file with zeros.
18. Executes:

    ```bash
    switch_root /root /sbin/init
    ```

After `switch_root`, the prepared Ubuntu system boots from the already-open
encrypted root mapping.

## Measurement Behavior

The disk hash is measured after the source image has been copied into the open
encrypted mapper. The measured bytes are read from `/dev/mapper/sda_crypt`, so
the value represents plaintext as exposed by dm-crypt, not the LUKS ciphertext
on `/dev/sda`.

The current script reads the mapper until EOF. If the destination disk is larger
than the copied source image, the extra mapper tail also contributes to the
hash.

Measurements are best-effort:

- If no TPM device or `tpm2_pcrextend` binary is present, PCR extension is
  skipped.
- If the TDX configfs RTMR interface or `tdx-rtmr-extend` is missing, RTMR
  extension is skipped.
- Failed PCR or RTMR extension logs a warning but does not stop boot.

## Key Lifetime

The LUKS key is generated inside the initramfs and is not written into the
cloned root filesystem. The script wipes the key file before switching root.

The encrypted root remains usable after the wipe because the dm-crypt mapping is
already open. A later reboot cannot unlock the destination disk from persisted
state unless an external flow repeats provisioning or supplies an unlock
mechanism.

## Failure Behavior

Most fatal failures clean up what they can, then drop to `/bin/sh` for
debugging. Cleanup commonly includes closing the LUKS mapping, wiping `kk.bin`,
and disconnecting NBD.

Non-fatal measurement failures only emit warnings and continue.

## Current Launch Assumptions

The current code documents and implements the initramfs-side flow, but the
launcher that boots a confidential VM into this initramfs must still provide the
runtime wiring:

- attach the blank destination disk as `/dev/sda`;
- make the NBD source reachable from the guest;
- pass `src_ip=<source-ip>` and optionally `src_port=<source-port>` on the
  kernel command line;
- include vTPM or TDX RTMR support if measurements are required.

The generated Buildroot `start-qemu.sh` template is a generic Buildroot test
launcher and does not currently encode the complete FDE runtime command line.
