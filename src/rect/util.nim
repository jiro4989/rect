import unicode

proc paste*(src, dst: seq[string], x, y: int): seq[string] =
  discard

proc pasteLine*(dst, src: string, x = 0): string =
  var rets = dst.toRunes
  for i, v in src.toRunes:
    let n = i + x
    if rets.len <= n:
      let diff = n + 1 - rets.len
      for _ in 1..diff:
        rets.add Rune(' ')
    rets[n] = v
  result = $rets
  