const doc = """
rect is a command to paste rectangle text

Usage:
  rect [Options] srcFile dstFile
  rect [Options] dstFile

Options:
  -x              Position to paste [default: 0]
  -y              Position to paste [default: 0]
  -h --help       Show this screen
  -v --version    Show version
  -X --debug      Debug mode ON
"""

import rect/util
import parseopt
from strutils import parseInt
from strformat import `&`

const
  version = "v1.0.0"

var
  useDebug: bool

proc logDebug(msg: string) =
  if useDebug:
    stderr.writeLine msg

proc readLines(f: File): seq[string] =
  var line: string
  while f.readLine line:
    result.add line

when isMainModule:
  var optParser = initOptParser()

  var
    x, y: int
    files: seq[string]
  for kind, key, val in optParser.getopt():
    case kind
    of cmdArgument:
      files.add key
    of cmdLongOption, cmdShortOption:
        case key
        of "help", "h":
          echo doc
          quit 0
        of "version", "v":
          echo version
          quit 0
        of "debug", "X":
          useDebug = true
        of "x": x = val.parseInt
        of "y": y = val.parseInt
    of cmdEnd: assert(false)  # cannot happen
  logDebug &"[DEBUG] command line parameters: x:{x}, y:{y}, files:{files}, debug:{useDebug}"

  var
    srcfile, dstfile: File
    src, dst: seq[string]
  try:
    case files.len
    of 2:
      logDebug &"[DEBUG] files.len:2"
      srcfile = open(files[0])
      dstfile = open(files[1])
    of 1:
      logDebug &"[DEBUG] files.len:1, use stdin"
      srcfile = stdin
      dstfile = open(files[0])
    else:
      stderr.writeLine "[ERR] a count of files must be 1 or 2"
      stderr.writeLine doc
      quit 1

    src = srcfile.readLines
    dst = dstfile.readLines
  except:
    stderr.writeLine getCurrentExceptionMsg()
  finally:
    logDebug &"[DEBUG] close files"
    if not srcfile.isNil: srcfile.close
    if not dstfile.isNil: dstfile.close
  logDebug &"[DEBUG] src:{src}, dst:{dst}"

  let lines = dst.paste(src, x = x, y = y)
  logDebug &"[DEBUG] lines:{lines}"

  for line in lines:
    echo line

  logDebug &"[DEBUG] Finish application"