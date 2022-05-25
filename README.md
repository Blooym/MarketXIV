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
- 🔧 Built-in Updater: All executable updates are handled for you (Manual Installs)
- 📦 Lightweight: Uses hardly any system storage or resources.
- ⚡ Powerful: Can fetch all the information you'd ever need.

## Installation Methods
### APT Package Manager
If you are using an Ubuntu-derivative system then use this installation method.

<details>
<summary>Show Steps</summary>

<br>
  
1. Add the repository hosting MarketXIV to your apt sources directory
```
echo "deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /" | sudo tee -a /etc/apt/sources.list.d/bitsofabyte.list
``` 

2. Fetch all sources again to detect the new list
```
sudo apt update
 ```

3. Install MarketXIV to your system
```
sudo apt install marketxiv
```
  
</details>  

---

### Yum/DNF Package Manager
If you are using Fedora, OpenSUSE, or any other system that supports the yum/dnf package manager then use this installation method.

<details>
<summary>Show Steps</summary>
<br>
  
1. Add the repository hosting MarketXIV to your yum repos directory
```
echo "[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0" | sudo tee -a /etc/yum.repos.d/bitsofabyte.repo
``` 
  
2. Fetch all sources again to detect the new list
```
sudo yum update
```

3. Install MarketXIV to your system
```
sudo yum install marketxiv
```

</details>  

---


### Using Scoop 
If you are using a Windows sytem, then use this installation method.

<details>
<summary>Show Steps</summary>
<br>

1. Install [scoop](https://scoop.sh/) using their guide
2. Add the bucket for MarketXIV
```
scoop bucket add marketxiv https://github.com/BitsOfAByte/MarketXIV.git
```
3. Install MarketXIV to your system
```
scoop install marketxiv/marketxiv
```

</details>

---

### Homebrew Package Manager
If your distributions package manager is not listed here or you wish to use [Homebrew](https://brew.sh).

<details>
<summary>Show Steps</summary>
<br>
  
1. Install homebrew if you haven't already got it
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
2. Add the tap for MarketXIV to homebrew
```
brew tap bitsofabyte/marketxiv https://github.com/BitsOfAByte/marketxiv.git
```
3. Install MarketXIV to your system
```
brew install marketxiv
```
  
</details>

---

### Manual Installation
Manually install the BInary from the release Archives.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/BitsOfAByte/marketxiv/releases/latest) for your system/architecture
2. Extract the binary into your system path or add the binary to your path.

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common for everyday hardware.

</details>

---

### From Source
Build MarketXIV directly from the GitHub source for any supported platform.
<details>  
<summary>Show Steps</summary>
  
Building MarketXIV from source is not recommended for beginners, but if you know what you're doing then follow these steps: 
1. Install [Go](https://go.dev/) on your system
2. Download the [GoReleaser](https://goreleaser.com/) package
3. Clone the repository to your system with `git clone https://github.com/BitsOfAByte/marketxiv`
4. Inside the repository directory, run `goreleaser build --single-target --rm-dist --snapshot` to build.

You will find the compiled binary for your OS & Arch inside of the `/dist` folder.

</details>  


## Analytics / Logs
MarketXIV collects some analytics to help identify issues and create patches for them. Any information submitted is done so in a way that provides as much privacy as possible, you can view all of the analytics code [here](./backend/analytics.go).

#### Information Collected

| Data              | Used For                                                      |
|-------------------|---------------------------------------------------------------|
| UUID              | Finding specific data for deletion requests                   |
| Command           | The command that triggered the log to be sent                 |
| Log Type          | Filtering out un-necessary logs when searching                |
| Log Message       | Knowing what happened and possible fixes                      |
| Architecture + OS | Diagnosing OS or Architecture specific bugs                   |
| MarketXIV Version | Diagnosing bugs in specific versions of the program           |
| Created At        | Knowing when issues occured for correlating with API downtime |

### Opting Out
`marketxiv config analytics false` will opt you out from all analytics.

#### Deletion
If you would like any previous logs deleted from the database, open an issue with your UUID (found using `marketxiv config uuid`).
