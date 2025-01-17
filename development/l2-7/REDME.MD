# Как использовать


```bash
  "item1;item2;item3;item4;item5" | go run main.go -f "1,2,5" -d=";"
```
результат
>  item1;item2;item5

<hr>

```bash
  "one two three" | go run main.go -f "2,3" -d=" "
```
результат
>  two three