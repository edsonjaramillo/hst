# hst

`hst` is a CLI helper for Atuin history management.

## Install

```bash
curl -fsSL https://github.com/edsonjaramillo/hst/releases/latest/download/install.sh | sh
```

For pinned versions and manual install steps, see `INSTALL.md`.

## Usage

### Commands

- `hst remove search`: interactively search and delete commands from Atuin history.
- `hst remove errors`: delete commands that returned non-zero exit codes.
- `hst remove fewer [fewer]`: delete commands occurring less than or equal to `fewer` times (default `1`) using interactive selection.
- `hst sync`: rewrite `$HISTFILE` with deduplicated commands from Atuin.

### Interactive dependency

`hst remove search` and `hst remove fewer` require `fzf`.
If `fzf` is missing, `hst` fails with an explicit dependency error.
In interactive selection, use `Alt-a` to select all and `Alt-d` to deselect all.

### Shell completions

Current session examples:

```bash
source <(hst completion zsh)
source <(hst completion bash)
hst completion fish | source
```
