<div align="center">

<img src = ".assets/cli-example.png" alt="MarketXIV Terminal Example" height="320"/>

### MarketXIV
Lightning-fast delivery of FFXIV market data to your Command Line 🚀🔥

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
- 🔥 Super Fast Responses: Get the information you need in a split second
- 🔧 Built-in Updater: All executable updates are handled for you, always use the best.
- 📦 Lightweight: Uses hardly any system storage or resources.
- ⚡ Powerful: Can fetch all the information you'd ever need.

## Installing
### Using Scoop (Windows)
Install [scoop](https://scoop.sh/) and then run the following commands in a non-escalated command prompt:
```
scoop bucket add marketxiv https://github.com/BitsOfAByte/MarketXIV.git
scoop install marketxiv/marketxiv
```

### Using Homebrew (MacOS & Linux)
Install [homebrew](https://brew.sh/) on your system and then run the following:
```
brew tap bitsofabyte/marketxiv https://github.com/BitsOfAByte/MarketXIV.git
brew install marketxiv
```

### Go Install (All Platforms)
⚠️ Support is not provided if the application does not work when installing with this method.

Install [Go](https://go.dev/) on your system and then run the following command, which will install MarketXIV to wherever your GoPath binaries folder is located. 
```
go install github.com/BitsOfAByte/marketxiv@latest 
```
When using Go to install the app certain build-time flags will not be set which may cause the application to say there is an update available. Running the update command after installing via Go is recommended.

### GitHub Releases (All Platforms, Manual)
Download the [newest release](https://github.com/BitsOfAByte/MarketXIV/releases/latest) and extract it somewhere inside of your system path, or alternatively add the executable to your path wherever you'd like it to be installed. [(Windows Path Guide)](https://www.maketecheasier.com/what-is-the-windows-path/) [(UNIX Path Guide)](https://www.computerhope.com/issues/ch001647.htm)

If you are confused about what architecture to download, then you will most likely want `amd64` on Windows/Intel-Processor Macs/Linux, `arm64` on new M-Processor Macs and `arm` on RasberryPi. The updater tool will download for your architecture once installed. 

*Some devices may use different architectures that are not listed here, it is best to check your processor info to confirm what you need to download*

### From Source (All Platforms)
If you would like to build MarketXIV from source, install both Go & GoReleaser on to your system, then run the one-liner below:
```
git clone https://github.com/BitsOfAByte/MarketXIV && cd MarketXIV && goreleaser build --single-target --rm-dist --snapshot
```
You will find the compiled binary for your OS & Arch inside of the `/dist` folder. Please keep in mind building manually will create a snapshot version instead of a full release, but you can still update as usual. 
