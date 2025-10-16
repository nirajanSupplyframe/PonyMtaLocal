{ pkgs ? import <nixpkgs> {} } :
pkgs.mkShell {
    inputsFrom = [ (pkgs.callPackage ./default.nix {}) ];
    buildInputs = with pkgs; [ go git gopls gosimports sqlite air];
    GIT_TERMINAL_PROMPTS = "1";
}
