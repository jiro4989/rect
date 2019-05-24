proc continuedInts*(s, e: int): seq[int] =
  ## Returns continued sequence.
  ##
  ## 連番のシーケンスを返す。
  runnableExamples:
    assert continuedInts(0, 3) == @[0, 1, 2, 3]
    assert continuedInts(1, 3) == @[1, 2, 3]

  for v in s..e:
    result.add v
