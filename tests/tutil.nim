import unittest

include rect/util

suite "continuedInts":
  test "normal":
    check continuedInts(0, 2) == @[0, 1, 2]
    check continuedInts(1, 3) == @[1, 2, 3]
