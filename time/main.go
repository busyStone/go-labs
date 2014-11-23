package main

import (
  "fmt"
  "log"
  "regexp"
  "strings"
  "time"
)

func main() {
  var t1 time.Time

  t := time.Now()

  log.Println(t.Unix(), t.UTC().Unix(), t.Local().Unix())

  log.Println(t.Unix())
  log.Println(t.UnixNano())
  log.Println(time.Unix(0, t.UnixNano()).UTC())
  log.Println(time.Unix(t.Unix(), 0).UTC())
  log.Println(t.UTC())
  log.Println(t.UTC().Local())
  log.Println(t)
  log.Println(t.Location())
  log.Println(t.UTC().Location())
  log.Println(t.Format("2006-01-02 15:04:05"))

  log.Println(t1)
  log.Println(t1.Format("2006-01-02 15:04:05"))
  log.Println("default uinx nano ", t1.UnixNano())

  log.Println(t.Format("15小时20分钟"))

  td := time.Duration(9000000000000000000)

  log.Println(td.Hours(), td.Minutes(), td.String())

  log.Println(TranDurationToTime(td))

  log.Println(t.UTC().UnixNano())

  log.Println(time.Now().Date())
}

func TranDurationToTime(d time.Duration) string {
  t := d.String()

  var result string

  // 将浮点 s 转换为 整数
  if strings.Contains(t, ".") {
    result, _ = ExtractDataByRegex(t, `(\.\d+s)`, nil)

    result = strings.TrimSuffix(t, result) + "s"
  } else {
    result = t
  }

  return result
}

func ExtractDataByRegex(html, query string, option map[string]interface{}) (string, error) {
  rx := regexp.MustCompile(query)
  value := rx.FindStringSubmatch(html)

  var v interface{}
  ok := false
  if len(value) == 0 {
    if v, ok = option["or"]; ok {
      var op string
      if op, ok = v.(string); ok {
        return ExtractDataByRegex(html, op, nil)
      }
    }

    return "", fmt.Errorf("正则表达式没有匹配到内容:(%s)", query)
  }

  if len(value) < 2 || strings.TrimSpace(value[1]) == "" {
    return "", fmt.Errorf("正则表达式没有匹配到内容:(%s)", query)
  }

  return strings.TrimSpace(value[1]), nil
}
