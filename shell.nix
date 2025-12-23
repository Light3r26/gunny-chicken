{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  ),
  mkGoEnv ? pkgs.mkGoEnv,
  gomod2nix ? pkgs.gomod2nix,
  lib,
  ...
}:

#let
  #goEnv = mkGoEnv { pwd = ./.; };
#in
pkgs.mkShell {
  packages = [
      #goEnv
    gomod2nix
    pkgs.go
    pkgs.gcc

    pkgs.pkg-config
    pkgs.libx11
    pkgs.libxcursor
    pkgs.libxrandr
    pkgs.libxinerama
    pkgs.libxi
    pkgs.libxxf86vm
    pkgs.mesa
    pkgs.libGL
    pkgs.xorg_sys_opengl
  ];

  LD_LIBRARY_PATH = lib.makeLibraryPath [
    pkgs.libGL
    pkgs.mesa
    pkgs.xorg_sys_opengl
  ];
}
