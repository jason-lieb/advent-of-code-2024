# Advent of Code 2024

Advent of Code challenges completed with Go and built with Nix

## Use

- To install dev tools, run `nix develop`
- To build the packages (and run tests), run `nix build`
- After building to run the executables, run `./result/bin/x`
- To run tests directly, go into the project directory and run `go test`
- To generate a new day, build the project and run `./result/bin/newDay dayx`
- Remember that nix will ignore new files if you don't run `git add` on them
