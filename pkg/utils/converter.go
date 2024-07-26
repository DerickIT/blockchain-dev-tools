package utils

import (
    "fmt"
    "strconv"
    "strings"
)

func ToHex(input string) (string, error) {
    n, err := strconv.ParseInt(input, 10, 64)
    if err != nil {
        return "", fmt.Errorf("invalid input: %v", err)
    }
    return fmt.Sprintf("0x%X", n), nil
}

// FromHex 将十六进制字符串转换为十进制整数
func FromHex(input string) (int64, error) {
  // 移除 "0x" 前缀（如果存在）
  input = strings.TrimPrefix(input, "0x")
  
  // 使用 strconv.ParseInt 进行转换，基数为16
  result, err := strconv.ParseInt(input, 16, 64)
  if err != nil {
      return 0, fmt.Errorf("invalid hex input: %v", err)
  }
  return result, nil
}