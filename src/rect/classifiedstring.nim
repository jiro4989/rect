import eastasianwidth
import util
import unicode
from sequtils import repeat, filterIt
from strutils import join

type
  ClassifiedString* = object
    data*: string
    indices*: seq[int]

proc toClassifiedString*(s: string): seq[ClassifiedString] =
  var i: int
  for v in s.toRunes:
    let w = v.`$`.stringWidth
    var seqi: seq[int]
    for j in i..<i+w:
      seqi.add j
    result.add ClassifiedString(data: $v, indices: seqi)
    i += w

proc pad*(dst, src: string, x = 0): array[2, string] =
  result[0] = dst
  result[1] = " ".repeat(x).join & src

proc continuedInts*(s, e: int): seq[int] =
  for v in s..e:
    result.add v

proc split3*(self: openArray[ClassifiedString], minIndex, maxIndex: int): array[3, seq[ClassifiedString]] =
  let
    lastElem = self[self.len-1].indices
    leftIndicesRange   = continuedInts(0, minIndex-1)
    centerIndicesRange = continuedInts(minIndex, maxIndex)
    rightIndicesRange  = continuedInts(maxIndex+1, lastElem[lastElem.len - 1])
  result[0] = self.filterIt(0 < it.indices.filterIt(it in leftIndicesRange).len)
  result[1] = self.filterIt(0 < it.indices.filterIt(it in centerIndicesRange).len)
  result[2] = self.filterIt(0 < it.indices.filterIt(it in rightIndicesRange).len)
  block:
    var v: seq[ClassifiedString]
    for r in result[0]:
      if r notin result[1]:
        v.add r
    result[0] = v
  block:
    var v: seq[ClassifiedString]
    for r in result[2]:
      if r notin result[1]:
        v.add r
    result[2] = v

proc first*(self: openArray[ClassifiedString]): int =
  if self.len < 1: return
  self[0].indices[0]

proc last*(self: openArray[ClassifiedString]): int =
  if self.len < 1: return
  let li = self[self.len-1].indices
  result = li[li.len-1]