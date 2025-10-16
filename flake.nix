{
    description = " gopro is an app i am building to basically get hands on with nix and go and to use sqlite which can be later replaced to postgres";
    inputs = {nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05"; };
    outputs = {self, nixpkgs} :
    let
        supportedSystems = ["x86_64-linux" "aarch64-darwin"];
        forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
        pkgsFor = nixpkgs.legacyPackages;
        in {
            packages = forAllSystems (system :  {
                default = pkgsFor.${system}.callsPackage ./default.nix { };
            });
            devShells = forAllSystems
                (system: { default = pkgsFor.${system}.callPackage ./shell.nix { }; });
        };
}