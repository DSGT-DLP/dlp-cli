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

Step 2: Run cli in debug mode with `go run main.go`. For example, if you wanted to run the frontend install command in debug mode, run `go run main.go frontend install`.

Step 3: Make sure to add package paths of all files as imports in main.go

## FAQs

### panic: input/output error

![input-output-error](https://github.com/DSGT-DLP/dlp-cli/assets/54150946/b9acfcca-4646-4086-9aea-e4f262520d87)
If you are on WSL, append the `-w` flag to tell the dlp-cli to run bash cmds without pseudoterminals, which are handled differently on windows and WSL. The way that the dlp-cli uses pseudoterminals currently only work in unix environments.

### Path does not exist: dlp-cli/manifests/d/DSGT-DLP/...
Make sure to pull the latest changes from the `dlp-cli` submodule. This can be done by `cd dlp-cli` (from the project directory) and running `git pull origin main`.

### Error starting cmd: "mamba" executable file not found in %PATH%
Windows: Rerun the mamba installer and check the box that says that it will add mamba to PATH.
