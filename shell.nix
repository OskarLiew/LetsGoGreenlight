{
  pkgs ? import <nixpkgs> { },
}:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    go-tools
    gopls
    hey
    (go-migrate.overrideAttrs (oa: {
      tags = [ "postgres" ];
    }))
  ];
  shellHook = ''
    echo "Ready!"
  '';
}
