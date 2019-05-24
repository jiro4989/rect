import eastasianwidth
import classifiedstring
from sequtils import repeat, filterIt, mapIt
from strutils import join

proc pasteLine*(dst, src: string, x = 0): string =
  ## dstにsrcを貼り付ける。全角文字幅を考慮する。
  runnableExamples:
    doAssert "123".pasteLine("a") == "a23"
    doAssert "abc".pasteLine("あ", x = 1) == "aあ"
    doAssert "あいう".pasteLine("あ1", x = 3) == "あ あ1"

  var
    dst2 = dst
    src2 = src
  let ds2 = pad(dst2, src2, x)
  dst2 = ds2[0]
  src2 = ds2[1]
  if dst2.stringWidth < src2.stringWidth:
    let diff = src2.stringWidth - dst2.stringWidth
    # Classifiedで位置をセットするためにパディング
    dst2 = dst2 & " ".repeat(diff).join
  let
    dst3 = dst2.toClassifiedString
    src3 = src2.toClassifiedString[x..^1]
    minIndex    = x
    maxIndex    = src3.last
    # 3分割
    s3     = dst3.split3(minIndex, maxIndex)
    left   = s3[0]
    center = s3[1] # centerの箇所をsrc3に置き換える
    right  = s3[2]

  result.add left.mapIt(it.data).join
  if src3.first != center.first:
    result.add " "
  result.add src3.mapIt(it.data).join
  if src3.last != center.last:
    result.add " "
  result.add right.mapIt(it.data).join
  
proc paste*(dst, src: seq[string], x = 0, y = 0): seq[string] =
  ## dstにsrcを貼り付ける。全角文字幅を考慮する。
  ## dstの幅を超過しようとした場合は行や文字を追加する。
  runnableExamples:
    let data = @["12345", "あいうえお", "678"]
    doAssert data.paste(@["abc"]) == @["abc45", "あいうえお", "678"]
    doAssert data.paste(@["abc", "def"], y = 2) == @["12345", "あいうえお", "abc", "def"]

  result = dst
  if src.len < 1: return
  for i, line in src:
    let n = i + y
    if result.len <= n:
      let diff = n + 1 - result.len
      for _ in 1..diff:
        result.add ""
    result[n] = result[n].pasteLine(line, x = x)
