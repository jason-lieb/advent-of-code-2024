{
  pkgs,
  buildGoApplication,
  script,
}:

let
  srcDirs = builtins.attrNames (builtins.readDir ./src);
  mkPackage =
    name:
    buildGoApplication {
      pname = name;
      version = "0.0";
      go = pkgs.go_1_23;
      src = ./src + "/${name}";
      modules = ./src + "/${name}/gomod2nix.toml";
    };
  src = builtins.listToAttrs (
    map (name: {
      inherit name;
      value = mkPackage name;
    }) srcDirs
  );
in
src.${script}
