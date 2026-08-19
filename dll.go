package zydis

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

type Zydis struct{}

//go:embed zydis.dll
var dllBytes []byte

var (
	dll       *windows.LazyDLL
	procCache = make(map[string]*windows.LazyProc)
)

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dir = filepath.Join(dir, "zydis_dll_cache")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}
	sha := sha256.Sum256(dllBytes)
	dllName := fmt.Sprintf("%s-%s.dll", "zydis", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, dllBytes, 0o644); err != nil {
			panic(err)
		}
	}
	// 用绝对路径加载，避免依赖进程级的 SetDllDirectory 全局状态。
	// SetDllDirectory 是进程级单值，多个包同时调用会互相覆盖，
	// 导致后续 LoadLibrary 找不到本包的 DLL 而 panic。
	// 绝对路径让 LoadLibraryW 直接定位文件，不依赖搜索路径。
	dll = windows.NewLazyDLL(filePath)
}

func getProc(name string) *windows.LazyProc {
	if p, ok := procCache[name]; ok {
		return p
	}
	p := dll.NewProc(name)
	procCache[name] = p
	return p
}
