# nvim-updater
Little tool to easily update to latest nightly nvim version.

## Install
```bash
go install github.com/olacin/nvim-updater@latest
```

## Usage

Default binary destination is `/usr/bin/nvim`. 

```bash
$ nvim-updater -h
Usage of nvim-updater:
  -dest string
    	executable directory destination (default "/usr/bin/nvim")
```

You can also specify a custom output destination with `-dest` flag.

```bash
$ nvim-updater -dest $HOME/.local/bin/nvim
2022/05/23 16:16:12 Fetching latest version of neovim
2022/05/23 16:16:12 Latest neovim nightly version is 9e1ee9fb1
2022/05/23 16:16:12 Current neovim version is 9e1ee9fb1
2022/05/23 16:16:12 Already at the latest version: latest=9e1ee9fb1 current=9e1ee9fb1
```
