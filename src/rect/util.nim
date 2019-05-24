proc continuedInts*(s, e: int): seq[int] =
  for v in s..e:
    result.add v