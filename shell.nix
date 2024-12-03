{
  pkgs,
  mkGoEnv,
  gomod2nix,
}:

let
  srcDirs = builtins.attrNames (builtins.readDir ./src);
  goEnvs = map (dir: mkGoEnv { pwd = ./src + "/${dir}"; }) srcDirs;
in
pkgs.mkShell { packages = [ gomod2nix ] ++ goEnvs; }
