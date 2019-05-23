import eastasianwidth

import unicode
from sequtils import repeat, mapIt, filterIt
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

proc pad(dst, src: string, x = 0): array[2, string] =
  result[0] = dst
  result[1] = " ".repeat(x).join & src

proc continuedInts(s, e: int): seq[int] =
  for v in s..e:
    result.add v

proc split3(self: openArray[ClassifiedString], minIndex, maxIndex: int): array[3, seq[ClassifiedString]] =
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

proc first(self: openArray[ClassifiedString]): int =
  if self.len < 1: return
  self[0].indices[0]

proc last(self: openArray[ClassifiedString]): int =
  if self.len < 1: return
  let li = self[self.len-1].indices
  result = li[li.len-1]

proc pasteLine*(dst, src: string, x = 0): string =
  var
    dst2 = dst
    src2 = src
  let ds2 = pad(dst2, src2, x)
  dst2 = ds2[0]
  src2 = ds2[1]
  if dst2.stringWidth < src2.stringWidth:
    let diff = src2.stringWidth - dst2.stringWidth
    dst2 = dst2 & " ".repeat(diff).join
  let
    dst3 = dst2.toClassifiedString
    src3 = src2.toClassifiedString[x..^1]
    minIndex    = x
    maxIndex    = src3.last
    s3     = dst3.split3(minIndex, maxIndex)
    left   = s3[0]
    center = s3[1]
    right  = s3[2]
  result.add left.mapIt(it.data).join
  if src3.first != center.first:
    result.add " "
  result.add src3.mapIt(it.data).join
  if src3.last != center.last:
    result.add " "
  result.add right.mapIt(it.data).join
  
proc paste*(dst, src: seq[string], x = 0, y = 0): seq[string] =
  result = dst
  if src.len < 1: return
  for i, line in src:
    let n = i + y
    if result.len <= n:
      let diff = n + 1 - result.len
      for _ in 1..diff:
        result.add ""
    result[n] = result[n].pasteLine(line, x = x)
