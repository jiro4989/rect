import unittest

include rect/crop

suite "cropLine":
  test "Default width":
    check "abcde".cropLine == "a"
    check "".cropLine == ""
  test "Width = 0":
    check "abcde".cropLine(width = 0) == ""
    check "".cropLine(width = 0) == ""
  test "Range over":
    check "abcde".cropLine(width = 7) == "abcde"
  test "Half width":
    check "abcde".cropLine(width = 1) == "a"
    check "abcde".cropLine(x = 1, width = 1) == "b"
  test "Full width, but result is empty":
    check "あいうえお".cropLine(width = 1) == " "
    check "あいうえお".cropLine(x = 1, width = 1) == " "
    check "あいうえお".cropLine(x = 2, width = 1) == " "
  test "Full width":
    check "あiうえお".cropLine(x = 2, width = 1) == "i"
    check "あいうえお".cropLine(width = 2) == "あ"
    check "あいうえお".cropLine(width = 3) == "あ "
    check "あいうえお".cropLine(width = 4) == "あい"
    check "あいうえお".cropLine(x = 1, width = 2) == "  "
    check "あいうえお".cropLine(x = 1, width = 3) == " い"
    check "あいうえお".cropLine(x = 1, width = 4) == " い "
    check "あいうえお".cropLine(x = 1, width = 5) == " いう"
    check "あiうえお".cropLine(width = 3) == "あi"
  test "Illegal parameter":
    check "abcde".cropLine(x = -1) == ""
    check "abcde".cropLine(x = -1, width = 2) == "a"
    check "abcde".cropLine(x = -2, width = 2) == ""
    check "abcde".cropLine(x = -2, width = 3) == "a"
    check "abcde".cropLine(x = -2, width = 4) == "ab"
    check "abcde".cropLine(x = 1, width = 0) == ""
    check "abcde".cropLine(x = 1, width = -1) == ""


suite "crop":
  const data = @["12345", "あいうえお", "678"]
  test "Default position":
    check data.crop == @["1"]
    check data.crop(width = 4) == @["1234"]
    check data.crop(width = 4, height = 2) == @["1234", "あい"]
    check data.crop(width = 3, height = 2) == @["123", "あ "]
  test "x = 1":
    check data.crop(x = 1, width = 4) == @["2345"]
    check data.crop(x = 1, width = 4, height = 2) == @["2345", " い "]
    check data.crop(x = 1, width = 3, height = 2) == @["234", " い"]
  test "y = 1":
    check data.crop(y = 1, width = 4) == @["あい"]
    check data.crop(y = 1, width = 4, height = 2) == @["あい", "678"]
    check data.crop(y = 1, width = 3, height = 2) == @["あ ", "678"]
  var empty: seq[string]
  test "Range over":
    check data.crop(y = 2, width = 4) == @["678"]
    check data.crop(y = 3, width = 4) == empty
    check data.crop(y = 255, width = 4) == empty
  test "Illegal parameter":
    check data.crop(x = -1) == empty
    check data.crop(y = -1) == empty
    check data.crop(width = 0) == empty
    check data.crop(width = -1) == empty
    check data.crop(height = 0) == empty
    check data.crop(height = -1) == empty
