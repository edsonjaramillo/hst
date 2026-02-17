# Releasing `hst`

## Automated release flow

Releases are tag-driven via GitHub Actions.

1. Ensure `main` is green (tests pass in CI).
2. Create and push a semver tag (`vX.Y.Z`).
3. GitHub Actions builds binaries and creates a GitHub Release.

## Quick command

```bash
scripts/release-tag.sh v0.2.0
```

## Published artifacts

- `hst` tarballs for:
  - `linux/amd64`
  - `linux/arm64`
  - `darwin/amd64`
  - `darwin/arm64`
- `SHA256SUMS`
- `install.sh`

## Versioning in code

The CLI version is injected at build time:

```bash
-ldflags="-X github.com/edsonjaramillo/hst/internal/cli.Version=<version>"
```

Default local version is `dev` when not provided.
