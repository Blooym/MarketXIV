<div align="center">

<img src = ".assets/cli-example.png" alt="MarketXIV Terminal Example" height="320"/>

### MarketXIV

Delivery of FFXIV market data to your Command Line

#### Built With & Powered By

[![GoLang-Badge](https://img.shields.io/badge/GoLang-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Universalis-Badge](https://img.shields.io/badge/Universalis-orange?style=for-the-badge&logo=swagger&logoColor=white)](https://universalis.app)
[![Universalis-Badge](https://img.shields.io/badge/XIVAPI-purple?style=for-the-badge&logo=swagger&logoColor=white)](https://xivapi.com)

</div>

---

## About

### Why?

MarketXIV was created due to a lack of command line tools for getting market data from any FFXIV API. The command line is always super light-weight to keep open and fast to use, making it perfect for getting the necessary information about the market in a short amount of time.

### Features

- 🔧 Built-in Updater: All executable updates are handled for you (Manual Installs)
- 📦 Lightweight: Uses hardly any system storage or resources.

## Installation Methods

> [!IMPORTANT]  
> The previously available repository at `packages.bitsofabyte.dev` is deprecated and will not be used for futher updates. Please remove it from your system's repositories, thanks!


### Manual Installation

Manually install the binary from the release archives.

<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/Blooym/marketxiv/releases/latest) for your system/architecture
2. Extract the tar archive and place it somewhere inside of your `$PATH`

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common.

</details>

---

### From Source

Build MarketXIV directly from the source for your system. Only recommended for advanced users that prefer to build from source or contributors.
<details>  
<summary>Show Steps</summary>

1. Make sure you have [Go](https://go.dev/) installed on your system and setup properly, alternatively use the [Devcontainer](./.devcontainer) setup.
2. Install [GoReleaser](https://goreleaser.com/) if you want to build using the supported buildsystem (Optional unless contributing)
3. Run `make build` to build the binary for your system, or `make build-all` to build for all supported systems. You can optionally use `./build/scripts/upx.sh <file>` to compress the binary with UPX (This is done automatically when using make and having GoReleaser installed with `SKIP_COMPRESS=false` set)
4. You will find all the binaries in the `./dist` directory alongside any other build artifacts.

</details>  
