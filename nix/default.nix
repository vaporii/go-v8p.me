{ pkgs ? (import <nixpkgs> {}), ... }:

pkgs.buildGoModule {
  pname = "v8p.me";
  version = "0.1.0";

  src = ../.;

  vendorHash = "sha256-yRgT6vp/EGWaglN7LKgKkaGP/hFzC86EncyZMQ6iRcA="; # not updated! i don't have access to my nix setup at the moment
}