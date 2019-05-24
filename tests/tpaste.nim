import unittest

include rect/paste

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
