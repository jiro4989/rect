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
    check pad("bbbbb", "a") == ["bbbbb", "a"]
    check pad("bbbbb", " a") == ["bbbbb", " a"]
    check pad("bbbbb", "a", 1) == ["bbbbb", " a"]
    check pad("a", "a", 1) == ["a", " a"]
  test "Full width":
    check pad("あ", "a", 1) == ["あ", " a"]
    check pad("あ", "a", 2) == ["あ", "  a"]

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

suite "first":
  test "1 element":
    check [ClassifiedString(indices: @[1, 2])].first == 1
  test "2 element":
    check [ClassifiedString(indices: @[1, 2]), ClassifiedString(indices: @[3, 4])].first == 1
  test "0 element":
    check [].first == 0

suite "last":
  test "1 element":
    check [ClassifiedString(indices: @[1, 2])].last == 2
  test "2 element":
    check [ClassifiedString(indices: @[1, 2]), ClassifiedString(indices: @[3, 4])].last == 4
    check [ClassifiedString(indices: @[1, 2]), ClassifiedString(indices: @[3])].last == 3
  test "0 element":
    check [].last == 0

suite "paste":
  test "normal":
    discard

suite "pasteLine":
  test "Half width":
    check "".pasteLine("abcde") == "abcde"
    check "123".pasteLine("abcde") == "abcde"
    check "123".pasteLine("a") == "a23"
  test "Half width, set X pos":
    check "".pasteLine("abcde", x = 1) == " abcde"
    check "123".pasteLine("abcde", x = 1) == "1abcde"
    check "123".pasteLine("abcde", x = 3) == "123abcde"
    check "123".pasteLine("abcde", x = 4) == "123 abcde"
    check "".pasteLine("abcde", x = 4) == "    abcde"
  test "Full width":
    check "".pasteLine("あいうえお") == "あいうえお"
    check "123".pasteLine("あ") == "あ3"
    check "abc".pasteLine("あ", x = 1) == "aあ"
    check "abc".pasteLine("あ", x = 2) == "abあ"
    check "abc".pasteLine("あ", x = 3) == "abcあ"
    check "abc".pasteLine("あ", x = 4) == "abc あ"
  test "Full width half":
    check "あいう".pasteLine("あ", x = 1) == " あ う"
    check "あいう".pasteLine("あ", x = 2) == "ああう"
    check "あいう".pasteLine("あ", x = 3) == "あ あ "
