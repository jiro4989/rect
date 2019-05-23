# Package

version       = "0.1.0"
author        = "jiro4989"
description   = "rect is a command to paste text rectangle."
license       = "MIT"
srcDir        = "src"
binDir        = "bin"
bin           = @["rect"]


# Dependencies

requires "nim >= 0.19.4"
requires "eastasianwidth >= 0.1.0"

task docs, "Generate documents":
  exec "nimble doc src/rect.nim -o:docs/rect.html"
  exec "nimble doc src/rect/util.nim -o:docs/util.html"

task ci, "Run CI":
  exec "nim -v"
  exec "nimble -v"
  exec "nimble test -y"
  exec "nimble docs -y"
  exec "nimble build -y"