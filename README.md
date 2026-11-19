# Go Template

CLI Utility for opening git repo's

## Usage

- Install with go
  ```sh
  go install github.com/ChrisPJohnstone/git-open@latest
  ```
- Get help using the tool
  ```sh
  git open -h
  ```
- Run the command without args to open the repo you are currently inside on it's current branch
  ```sh
  git open
  ```
- Optionally you can provide a path with the first positional argument
  ```sh
  git open README.md
  ```

  - If specified will open the directory / file in question
  - By default will open at repo root
