{
  description = "Advent of Code";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { self
    , flake-utils
    , nixpkgs
    }:

    flake-utils.lib.eachDefaultSystem (system:
    let
      overlays = [ (self: super: { go = super."go_1_19"; }) ];
      pkgs = import nixpkgs { inherit overlays system; };
    in
    {
      devShells.default = pkgs.mkShellNoCC {
        buildInputs = with pkgs; [
          # go 1.19 (specified by overlay)
          go

          # goimports, godoc, etc.
          gotools

          # https://github.com/golangci/golangci-lint
          golangci-lint

          entr
          fd
        ];
        shellHook = ''${pkgs.go}/bin/go version'';
      };
    });
}
