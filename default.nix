{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
    pname = "gopro";
    version = "0.1";
    src = pkgs.lib.cleanSource ./.;
    nativeBuildInputs = builtins.attrValues {inherit (pkgs) go; };
    buildPhase = ''
        export  HOME=$(pwd)
        go build -mod vendor -o gp
        '';

        installPhase = ''
        cp gp $out
    '';
}