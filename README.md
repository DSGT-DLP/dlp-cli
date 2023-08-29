# dlp-cli installation

Step 1: Install homebrew (homebrew only works on mac and linux, so windows users may need WSL, may look into supporting winget in the future). You can find installation instructions for homebrew at [https://brew.sh/](https://brew.sh/)

Step 2: Run `brew tap DSGT-DLP/dlp-cli` in the terminal

Step 3: Run `brew install dlp-cli` in the terminal

Usage: Run `dlp-cli` in DLP's project directory in the terminal

## Cli Development Info

Run cli in debug mode with `go run main.go`

Make sure to add package paths of all files as imports in main.go
