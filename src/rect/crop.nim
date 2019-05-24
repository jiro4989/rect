import classifiedstring
from sequtils import filterIt, mapIt
from strutils import join

proc cropLine*(src: string, x = 0, width = 1): string =
  var
    x2 = x
    width2 = width
  if x2 < 0:
    x2 = 0
    width2 = width2 - (x2 - x)
  if src.len <= 0 or width2 <= 0: return

  let indicesRange = continuedInts(x2, x2 + width2 - 1)
  let matched = src.toClassifiedString.filterIt(0 < it.indices.filterIt(it in indicesRange).len)
  for v in matched:
    let m = v.indices.filterIt(x2 <= it and it < x2 + width2)
    if m.len == v.indices.len:
      result.add v.data
    else:
      result.add " "

proc crop*(src: seq[string], x = 0, y = 0, width = 1, height = 1): seq[string] =
  var
    x2 = x
    y2 = y
    width2 = width
    height2 = height
  if x2 < 0:
    x2 = 0
    width2 = width2 - (x2 - x)
  if y2 < 0:
    y2 = 0
    height2 = height2 - (y2 - y)
  if width2 < 1 or height2 < 1: return

  for yy in y2..<y2+height2:
    if src.len <= yy: return
    result.add src[yy].cropLine(x = x2, width = width2)