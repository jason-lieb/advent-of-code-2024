{
  description = "Advent of Code 2024";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
    gomod2nix.inputs.flake-utils.follows = "flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      gomod2nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        callPackage = pkgs.callPackage;
        buildPackage =
          script:
          callPackage ./build.nix {
            inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
            inherit script;
          };

        srcDirs = builtins.attrNames (builtins.readDir ./src);
        packages = builtins.listToAttrs (
          map (name: {
            inherit name;
            value = buildPackage name;
          }) srcDirs
        );

        packageOutputs = builtins.attrValues packages;
      in
      {
        packages = packages // {
          default = pkgs.runCommand "default" { buildInputs = packageOutputs; } ''
            mkdir -p $out/bin
            for pkg in ${builtins.concatStringsSep " " (map (pkg: "${pkg}/bin/*") packageOutputs)}; do
              cp $pkg $out/bin
            done
          '';
        };

        devShells.default = callPackage ./shell.nix {
          inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
        };
      }
    );
}
