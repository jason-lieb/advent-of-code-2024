{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [ (import "${fetchTree gomod2nix.locked}/overlay.nix") ];
    }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
  script,
}:

let
  src = {
    day1 = buildGoApplication {
      pname = "day1";
      version = "0.1";
      go = pkgs.go_1_23;
      src = ./src/day1;
      modules = ./src/day1/gomod2nix.toml;
    };
  };
in
src.${script}
