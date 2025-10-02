[![Go Reference](https://pkg.go.dev/badge/github.com/ugent-library/vo.svg)](https://pkg.go.dev/github.com/ugent-library/vo)

# vo

Simple and typesafe Go validation library

```go
v := vo.New(
    vo.NotEmpty("titles", rec.Titles),
)

for i, title := range rec.Titles {
    v.In("titles").Index(i).Add(
        vo.ISO639_2("lang", title.Lang),
        vo.NotBlank("val", title.Val),
    )
}

return v.Validate().ToError()
```