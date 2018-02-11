## Todo
  * Factorize CURLY_BRACKETS ROUND_BRACKETS
  * IdentifyDelimiter error return
  * syntax_utils_test.go:14
  * File name reference when syntax error happens

  * alternatives to json

  * orden los testeos de las variantes. las variantes Wrapper deben ser testeadas en ultimo
    lugar (dentro de la iteracion de ese componenetes). Porque sino darai un falso positivo

    e.g.:

    ```
    import (
      "fmt"
      "strconv"
    )
    import "math"
    ```

    primero tesear con la variante one-line, y luego con la functional/wrapper

  * TabsStrategy: (config.go:172)
    pattern: \t VS \s\s\s\s        looks the same but are different -> syntax error
    ---> permisive ident: \t == \s\s\s\s
    ---> strict ident:    \t != \s\s\s\s

  * forEach variant.Delimiters (config.go:53)

  * Override config.Identation with variant.Delimiters if exists


## Performance

## Ideas

  * permissive struct's last `,`

  * one-line/mini if-else: a := (b > c) ? true : false

  * config.autocompile = true -> `go build...`

  * block commenting

  * custom error message
