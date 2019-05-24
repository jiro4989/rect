import classifiedstring
from sequtils import filterIt, mapIt
from strutils import join

proc cropLine*(src: string, x = 0, width = 1): string =
  if src.len <= 0 or width == 0: return
  let indicesRange = continuedInts(x, x + width - 1)
  let matched = src.toClassifiedString.filterIt(0 < it.indices.filterIt(it in indicesRange).len)
  for v in matched:
    let m = v.indices.filterIt(x <= it and it < x + width)
    if m.len == v.indices.len:
      result.add v.data
    else:
      result.add " "

proc crop*(src: seq[string], x = 0, y = 0, width = 1, height = 1): seq[string] =
  if width == 0: return src
  for y2 in y..<y+height:
    if src.len <= y2: return
    result.add src[y2].cropLine(x = x, width = width)