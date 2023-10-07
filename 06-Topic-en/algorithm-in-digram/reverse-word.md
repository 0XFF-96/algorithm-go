# reverse word 


```

func reverseWords(s string) string {
    strList := strings.Split(s, " ")
    res := make([]string, 0, len(strList))

    for i := len(strList) - 1; i >= 0; i-- {
        str := strings.TrimSpace(strList[i])

        if len(str) > 0 {
            res = append(res, strList[i])
        }
    }

    return strings.Join(res, " ")
}

```
