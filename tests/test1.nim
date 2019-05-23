import unittest

include rect/util

suite "toClassifiedString":
  test "Half width":
    check "123".toClassifiedString == @[ClassifiedString(data: "1", indices: @[0]), ClassifiedString(data: "2", indices: @[1]), ClassifiedString(data: "3", indices: @[2])]
    check " 23".toClassifiedString == @[ClassifiedString(data: " ", indices: @[0]), ClassifiedString(data: "2", indices: @[1]), ClassifiedString(data: "3", indices: @[2])]
  test "Full width":
    check "あい".toClassifiedString == @[ClassifiedString(data: "あ", indices: @[0, 1]), ClassifiedString(data: "い", indices: @[2, 3])]
    check "1い2う".toClassifiedString == @[ClassifiedString(data: "1", indices: @[0]), ClassifiedString(data: "い", indices: @[1, 2]), ClassifiedString(data: "2", indices: @[3]), ClassifiedString(data: "う", indices: @[4, 5])]

suite "pad":
  test "Half width":
    var
      a = "a"
      b = "bbbbb"
    pad(a, b)
    check a == "a    "
    check b == "bbbbb"
  test "Half width, set pos":
    var
      a = "a"
      b = "bbbbb"
    pad(a, b, 2)
    check a == "a      "
    check b == "  bbbbb"
  test "Full width, set pos":
    var
      a = "a"
      b = "あいう"
    pad(a, b, 1)
    check a == "a      "
    check b == " あいう"

suite "continuedInts":
  test "normal":
    check continuedInts(0, 2) == @[0, 1, 2]
    check continuedInts(1, 3) == @[1, 2, 3]

suite "split3":
  let data = @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])]
  test "center":
    check data.split3(2, 3) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3])], @[ClassifiedString(indices: @[4, 5])]]
  test "left center":
    check data.split3(2, 4) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]
  test "center right":
    check data.split3(1, 3) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3])], @[ClassifiedString(indices: @[4, 5])]]
  test "center only":
    check data.split3(0, 4) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]
    check data.split3(0, 5) == [@[], @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[]]

  let data2 = @[ClassifiedString(indices: @[0, 1]), ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5]), ClassifiedString(indices: @[6, 7])]
  test "center":
    check data2.split3(2, 4) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[ClassifiedString(indices: @[6, 7])]]
    check data2.split3(2, 5) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5])], @[ClassifiedString(indices: @[6, 7])]]
  test "left center":
    check data2.split3(2, 6) == [@[ClassifiedString(indices: @[0, 1])], @[ClassifiedString(indices: @[2, 3]), ClassifiedString(indices: @[4, 5]), ClassifiedString(indices: @[6, 7])], @[]]

suite "paste":
  test "normal":
    discard

# suite "pasteLine":
#   test "Half width":
#     check "".pasteLine("abcde") == "abcde"
#     check "123".pasteLine("abcde") == "abcde"
#   test "Half width, set X pos":
#     check "".pasteLine("abcde", x = 1) == " abcde"
#     check "123".pasteLine("abcde", x = 1) == "1abcde"
#     check "123".pasteLine("abcde", x = 3) == "123abcde"
#     check "123".pasteLine("abcde", x = 4) == "123 abcde"
#     check "".pasteLine("abcde", x = 4) == "    abcde"
#   test "Full width":
#     check "".pasteLine("あいうえお") == "あいうえお"
#     check "123".pasteLine("あい") == "あ3"
