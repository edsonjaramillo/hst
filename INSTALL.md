# Install `hst`

`hst` is distributed through GitHub Releases.

## Quick install (latest release)

```bash
curl -fsSL https://github.com/edsonjaramillo/hst/releases/latest/download/install.sh | sh
```

By default this installs to `~/.local/bin`.

## Install a pinned version

```bash
curl -fsSL https://github.com/edsonjaramillo/hst/releases/latest/download/install.sh | sh -s -- --version v0.2.0
```

## Install to a custom directory

```bash
curl -fsSL https://github.com/edsonjaramillo/hst/releases/latest/download/install.sh | sh -s -- --install-dir /usr/local/bin
```

If `/usr/local/bin` needs elevated permissions:

```bash
curl -fsSL https://github.com/edsonjaramillo/hst/releases/latest/download/install.sh -o /tmp/hst-install.sh
sudo HST_INSTALL_DIR=/usr/local/bin sh /tmp/hst-install.sh
```

## Manual install

1. Download the matching archive and `SHA256SUMS` from the release page.
2. Verify checksum:

```bash
sha256sum -c SHA256SUMS --ignore-missing
```

On macOS without `sha256sum`:

```bash
shasum -a 256 hst_<version>_<os>_<arch>.tar.gz
```

3. Extract and install:

```bash
tar -xzf hst_<version>_<os>_<arch>.tar.gz
install -m 0755 hst_<version>_<os>_<arch>/hst ~/.local/bin/hst
```

## Supported platforms

- `linux/amd64`
- `linux/arm64`
- `darwin/amd64`
- `darwin/arm64`
