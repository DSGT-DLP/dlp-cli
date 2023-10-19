# dlp-cli installation

## Unix installation (Linux, MacOS)

Step 1: Install homebrew. You can find installation instructions for homebrew at [https://brew.sh/](https://brew.sh/)

Step 2: Run `brew tap DSGT-DLP/dlp-cli` in the terminal

Step 3: Run `brew install dlp-cli` in the terminal

## Windows installation

Step 1: Install scoop. You can find installation instructions for scoop at [https://scoop.sh/](https://scoop.sh/)

Step 2: Run `scoop bucket add scoop-dlp-cli https://github.com/DSGT-DLP/scoop-dlp-cli.git` in the terminal

Step 3: Run `scoop install scoop-dlp-cli/dlp-cli` in the terminal

## Usage

Run `dlp-cli` in DLP's project directory in the terminal

## READ IF YOU WANT TO CONTRIBUTE TO THE dlp-cli:

Step 1: If you don't have the `dlp-cli` submodule, run `git submodule init dlp-cli` in the project directory. Then run `git submodule update`. If it says that you are in a detached HEAD state, cd to the dlp-cli submodule directory and `git checkout main`.

Step 2: Run cli in debug mode with `go run main.go`

Step 3: Make sure to add package paths of all files as imports in main.go
