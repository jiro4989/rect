import eastasianwidth

import unicode
from sequtils import repeat
from strutils import join

type
  ClassifiedString* = object
    data*: string
    indices*: seq[int]

proc toClassifiedString(s: string): seq[ClassifiedString] =
  var i: int
  for v in s.toRunes:
    let w = v.`$`.stringWidth
    var seqi: seq[int]
    for j in i..<i+w:
      seqi.add j
    result.add ClassifiedString(data: $v, indices: seqi)
    i += w

proc pad(src, dst: var string, x = 0) =
  dst = " ".repeat(x).join & dst
  let diff = dst.stringWidth - src.stringWidth
  src = src & " ".repeat(diff).join

proc paste*(src, dst: seq[string], x, y: int): seq[string] =
  discard

proc pasteLine*(dst, src: string, x = 0): string =
  var rets = dst.toRunes
  let srcRunes = src.toRunes
  var managePos: int
  for i, _ in srcRunes:
    let
      i2 = i + managePos
      v = srcRunes[i2]
      n = i2 + x
    if rets.len <= n:
      let diff = n + 1 - rets.len
      for _ in 1..diff:
        rets.add Rune(' ')
    rets[n] = v
    if v.`$`.stringWidth == 2:
      inc managePos
  result = $rets
  