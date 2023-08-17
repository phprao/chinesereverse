# chinesereverse
golang版本的简体繁体互转

## 基本使用方式
```go
fmt.Println(chinesereverse.SimplifiedToTraditional("连续"))
fmt.Println(chinesereverse.TraditionalToSimplified("連續"))
```

## 增加自定义简繁对照表
```go
WithExtraDictFile("dict2.txt")

fmt.Println(SimplifiedToTraditional("中"))
fmt.Println(TraditionalToSimplified("种"))
```