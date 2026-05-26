# Zydis Go Bindings

基于 [Zydis](https://github.com/zyantific/zydis) 的 Go 语言绑定，**纯 Go 实现，无 CGO 依赖**。

## 特性

- **🚀 无 CGO**: 纯 Go 实现，使用 `golang.org/x/sys/windows` 动态加载 DLL
- **自动生成**: 从 Zydis 头文件自动生成所有绑定代码
- **DLL 嵌入**: `zydis.dll` 嵌入到 Go 二进制文件中
- **完整覆盖**: 包含所有 Zydis 解码、格式化函数

## 使用方法

```go
package main

import (
    "fmt"
    "github.com/ddkwork/zydis"
)

func main() {
    z := &zydis.Zydis{}
    
    // 解码指令
    data := []byte{0x48, 0x89, 0xE5} // mov rbp, rsp
    var insn zydis.ZydisDecodedInstruction
    if z.DecodeInstructionBuffer(data, 64, &insn) == zydis.ZYAN_STATUS_SUCCESS {
        fmt.Printf("Decoded: %s\n", insn.Mnemonic)
    }
}
```

## 构建

```bash
# 生成绑定
go test -run TestGenerate -v

# 构建 DLL (需要 EWDK)
build.cmd
```

## 目录结构

```
zydis/
├── Zydis.go           # 自动生成的绑定代码
├── dll.go             # DLL 加载和嵌入
├── zydis.dll          # Zydis 动态库
├── build.cmd          # DLL 构建脚本
├── generate_test.go   # 绑定生成配置
└── clone/             # Zydis 源码
```

## 版本

- Zydis: 4.x
- Go 1.26+
- 平台: Windows x64

## 许可证

MIT License
