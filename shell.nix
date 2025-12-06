{
  pkgs ? import <nixpkgs> { },
}:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
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
