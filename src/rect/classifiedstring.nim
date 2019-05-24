import eastasianwidth
import util
import unicode
from sequtils import repeat, filterIt
from strutils import join

type
  ClassifiedString* = object
    ## 位置を保持する文字列。
    ## 半角文字はその文字自体と、その文字が行文字列におけるどのような位置を持つ
    ## 文字列であると、ここでは考える。
    ## 日本語のような全角文字列は、２つの位置を持つもの、としてここでは扱う。
    ##
    ## 例：
    ## * `a == data: a, indices: 0`
    ## * `あ == data: あ, indices: 0, 1`
    data*: string      ## String data
    indices*: seq[int] ## Position indices

proc toClassifiedString*(s: string): seq[ClassifiedString] =
  ## Converts to ClassifiedString.
  ##
  ## 文字列をClassifiedStringのシーケンスに変換する。
  runnableExamples:
    doAssert "abc".toClassifiedString == @[ClassifiedString(data: "a", indices: @[0]),
                                           ClassifiedString(data: "b", indices: @[1]),
                                           ClassifiedString(data: "c", indices: @[2])]
    doAssert "あc".toClassifiedString == @[ClassifiedString(data: "あ", indices: @[0, 1]),
                                           ClassifiedString(data: "c", indices: @[2])]

  var i: int
  for v in s.toRunes:
    let w = v.`$`.stringWidth
    var seqi: seq[int]
    for j in i..<i+w:
      seqi.add j
    result.add ClassifiedString(data: $v, indices: seqi)
    i += w

proc pad*(dst, src: string, x = 0): array[2, string] =
  ## srcをxの値の分だけ左に半角スペースを追加する。
  ##
  ## FIXME: dstが不要
  runnableExamples:
    discard

  result[0] = dst
  result[1] = " ".repeat(x).join & src

proc split3*(self: openArray[ClassifiedString], minIndex, maxIndex: int): array[3, seq[ClassifiedString]] =
  ## ClassifiedStringを`minIndex`未満、`minIndex`以上`maxIndex`以下、それ以外の３つの文字列を含む配列に変換する。
  runnableExamples:
    let data = @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])]
    doAssert data.split3(2, 3) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3])], @[ClassifiedString(indices: @[4, 5])]]
    doAssert data.split3(2, 4) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]
    doAssert data.split3(1, 3) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3])], @[ClassifiedString(indices: @[4, 5])]]
    doAssert data.split3(0, 4) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]
    doAssert data.split3(0, 5) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]

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
  ## Returns a first index of ClassifiedString.
  ##
  ## ClassifiedStringが保持する最初のインデックスを返す。
  runnableExamples:
    doAssert [ClassifiedString(indices: @[1, 2]), ClassifiedString(indices: @[3, 4])].first == 1
    doAssert [ClassifiedString(indices: @[3, 2]), ClassifiedString(indices: @[3, 4])].first == 3

  if self.len < 1: return
  self[0].indices[0]

proc last*(self: openArray[ClassifiedString]): int =
  ## Returns a last index of ClassifiedString.
  ##
  ## ClassifiedStringの最後のインデックスを返す。
  runnableExamples:
    doAssert [ClassifiedString(indices: @[1, 2]), ClassifiedString(indices: @[3, 4])].last == 4
    doAssert [ClassifiedString(indices: @[3, 2]), ClassifiedString(indices: @[3, 1])].last == 1

  if self.len < 1: return
  let li = self[self.len-1].indices
  result = li[li.len-1]
