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
    check "あいう".pasteLine("か1", x = 1) == " か1う"
    check "あいう".pasteLine("あ", x = 2) == "ああう"
    check "あいう".pasteLine("あ1", x = 2) == "ああ1 "
    check "あいう".pasteLine("あ", x = 3) == "あ あ "
    check "あいう".pasteLine("あ1", x = 3) == "あ あ1"

suite "paste":
  const data = @["12345", "あいうえお", "678"]
  test "Half width: 1 line, y = 0":
    const data2 = @["abc"]
    check data.paste(data2) == @["abc45", "あいうえお", "678"]
    check data.paste(data2, x = 1) == @["1abc5", "あいうえお", "678"]
    check data.paste(data2, x = 3) == @["123abc", "あいうえお", "678"]
  test "Half width: 1 line, y != 0":
    const data2 = @["abc"]
    check data.paste(data2, y = 1) == @["12345", "abc うえお", "678"]
    check data.paste(data2, y = 2) == @["12345", "あいうえお", "abc"]
    check data.paste(data2, y = 3) == @["12345", "あいうえお", "678", "abc"]
    check data.paste(data2, x = 1, y = 3) == @["12345", "あいうえお", "678", " abc"]
    check data.paste(data2, x = 1, y = 4) == @["12345", "あいうえお", "678", "", " abc"]
  test "Half width: 2 line, y = 0":
    const data2 = @["abc", "def"]
    check data.paste(data2) == @["abc45", "def うえお", "678"]
    check data.paste(data2, x = 1) == @["1abc5", " defうえお", "678"]
    check data.paste(data2, x = 3) == @["123abc", "あ defえお", "678"]
  test "Half width: 2 line, y != 0":
    const data2 = @["abc", "def"]
    check data.paste(data2, y = 1) == @["12345", "abc うえお", "def"]
    check data.paste(data2, y = 2) == @["12345", "あいうえお", "abc", "def"]
    check data.paste(data2, y = 3) == @["12345", "あいうえお", "678", "abc", "def"]
    check data.paste(data2, x = 1, y = 3) == @["12345", "あいうえお", "678", " abc", " def"]
    check data.paste(data2, x = 1, y = 4) == @["12345", "あいうえお", "678", "", " abc", " def"]
  test "Full width":
    check data.paste(@["abc", "かきく"]) == @["abc45", "かきくえお", "678"]
    check data.paste(@["abc", "かきx"]) == @["abc45", "かきx えお", "678"]
    check data.paste(@["abc", "かきく"], x = 1) == @["1abc5", " かきく お", "678"]
    check data.paste(@["abc", "かきx"], x = 1) == @["1abc5", " かきxえお", "678"]
    check data.paste(@["abc", "かきx"], x = 1, y = 2) == @["12345", "あいうえお", "6abc", " かきx"]
    check data.paste(@["abc", "かきく"], x = 9, y = 1) == @["12345", "あいうえ abc", "678      かきく"]
    check data.paste(@["abc", "かきく"], x = 10, y = 1) == @["12345", "あいうえおabc", "678       かきく"]
    check data.paste(@["abc", "かきく"], x = 11, y = 1) == @["12345", "あいうえお abc", "678        かきく"]
    check data.paste(@["abc", "かきく"], x = 12, y = 1) == @["12345", "あいうえお  abc", "678         かきく"]
    check data.paste(@["abc", "かきく"], x = 12, y = 4) == @["12345", "あいうえお", "678", "", "            abc", "            かきく"]
