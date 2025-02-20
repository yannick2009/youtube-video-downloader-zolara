{
  description = "Un environnement de développement avec yt-dlp et Go 1.24";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config = {
            allowUnfree = true;
          };
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            yt-dlp
            go_1_24
          ];

          shellHook = ''
            echo "Environnement de développement initialisé avec:"
            echo "- yt-dlp version: $(yt-dlp --version)"
            echo "- Go version: $(go version)"
          '';
        };
      }
    );
}
