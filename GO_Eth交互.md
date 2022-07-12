# Go Eth 智能合約交互

### ABI 生成GO文件

``` 
abigen --abi xx.abi --pkg packageName --type structName --out xx.go
```

- `abi` 文件在 `remix` 部署時可以得到
- `pkg` 指定輸出文件的包名，也就是 package 名稱
- `type` 指定合約結構體名稱
- `out` 指定輸出go文件名稱

