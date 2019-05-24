const doc = """
rect is a command to crop/paste rectangle text

Usage:
  rect crop  [options]
  rect crop  [options] <srcfile>
  rect paste [options] <dstfile>
  rect paste [options] <srcfile> <dstfile>
  rect (-h | --help)
  rect (-v | --version)

Options:
  -h --help               Show this screen
  -v --version            Show version
  -X --debug              Debug mode ON
  -x <x>                  Cropping/pasting position X [default: 0]
  -y <y>                  Cropping/pasting position Y [default: 0]
  -W --width <width>      Cropping width              [default: 1]
  -H --height <height>    Cropping height             [default: 1]
"""

import docopt
import rect/[paste, crop]

from strutils import parseInt, parseBool
from strformat import `&`

const
  version = "v1.0.0"

var
  useDebug: bool

proc logDebug(msg: string) =
  if useDebug:
    stderr.writeLine "[DEBUG] " & msg

proc logErr(msg: string) =
  stderr.writeLine "[ERR] " & msg

proc readLines(f: File): seq[string] =
  var line: string
  while f.readLine line:
    result.add line

proc execCrop(args: Table[string, Value]) =
  logDebug &"Execute `crop` subcommand: args:{args}"
  let 
    x = parseInt($args["-x"])
    y = parseInt($args["-y"])
    w = parseInt($args["--width"])
    h = parseInt($args["--height"])
    srcFile = $args["<srcfile>"]
    f = if srcFile == "" or srcFile == "nil": stdin
        else: open(srcFile)
  defer: f.close
  let lines = f.readLines.crop(x=x, y=y, width=w, height=h)
  for line in lines:
    echo line

proc execPaste(args: Table[string, Value]) =
  logDebug &"Execute `paste` subcommand: args:{args}"
  let 
    x = parseInt($args["-x"])
    y = parseInt($args["-y"])
    srcFile = $args["<srcfile>"]
    dstFile = $args["<dstfile>"]
    src = if srcFile == "" or srcFile == "nil": stdin
          else: open(srcFile)
    dst = open(dstFile)
  defer:
    src.close
    dst.close
  let
    srcData = src.readLines
    dstData = dst.readLines
    lines = dstData.paste(srcData, x=x, y=y)
  for line in lines:
    echo line

when isMainModule:
  let args = docopt(doc, version = version)
  useDebug = parseBool($args["--debug"])

  if args["crop"]:
    execCrop(args)
    quit 0

  if args["paste"]:
    execPaste(args)
    quit 0
  
  logErr "illegal options"
  stderr.writeLine doc
  quit 1