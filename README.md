# dlp-cli installation

## Unix installation (Linux, MacOS)

Step 1: Install homebrew (homebrew only works on mac and linux, so windows users may need WSL, may look into supporting winget in the future). You can find installation instructions for homebrew at [https://brew.sh/](https://brew.sh/)

Step 2: Run `brew tap DSGT-DLP/dlp-cli` in the terminal

Step 3: Run `brew install dlp-cli` in the terminal

## Windows installation

Step 1: If you don't have the `dlp-cli` submodule, run `git submodule init dlp-cli` in the project directory. Then run `git submodule update`.

Step 2: In the project directory, run `winget install -m dlp-cli/manifests/d/DSGT-DLP/dlp-cli/{version}`. Replace `{version}` with the version number that you want (typically the latest version).

## Usage

Run `dlp-cli` in DLP's project directory in the terminal

## Cli Development Info

Run cli in debug mode with `go run main.go`

Make sure to add package paths of all files as imports in main.go
