# dlp-cli installation

## Unix installation (Linux, MacOS, WSL)

Step 0: If on Windows, install [WSL](https://learn.microsoft.com/en-us/windows/wsl/install)

Step 1: Install homebrew. You can find installation instructions for homebrew at [https://brew.sh/](https://brew.sh/)

Step 2: Run `brew tap DSGT-DLP/dlp-cli` in the terminal

Step 3: Run `brew install dlp-cli` in the terminal

## Windows installation (Do this if WSL doesn't work)

Step 1: If you don't have the `dlp-cli` submodule, run `git submodule init dlp-cli` in the project directory. Then run `git submodule update`. If it says that you are in a detached HEAD state, cd to the dlp-cli submodule directory and `git checkout main`.

Step 2: Run `winget settings --enable LocalManifestFiles` as administrator. Then, in the project directory, run `winget install -m dlp-cli/manifests/d/DSGT-DLP/dlp-cli/{version}`. Replace `{version}` with the version number that you want (typically the latest version).

Step 3: Add the path to the install location to your PATH environment variable (somewhere within `%localappdata%\microsoft\winget\packages`, may need file explorer to find the exact folder)

## Usage

Run `dlp-cli` in DLP's project directory in the terminal

## Cli Development Info

Run cli in debug mode with `go run main.go`

Make sure to add package paths of all files as imports in main.go
