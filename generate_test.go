package zydis

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:     "clone/zydis/include",
		OutputDir:      ".",
		PackageName:    "zydis",
		RecurseHeaders: true,
		HeaderOrder:    []string{"Zydis/Zydis.h"},
		BindDll:        true,
		DllName:        "zydis.dll",
		Predefined: `
#define ZYDIS_STATIC_BUILD
#define ZYCORE_STATIC_BUILD
#define ZYAN_REQUIRES_LIBC
#define ZYAN_FALSE 0
#define ZYAN_TRUE 1
#define ZYAN_NO_LIBC
`,
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "Zydis") || strings.HasPrefix(name, "Zyan")
		},
		ExtraIncludeDirs: []string{
			"clone/zydis/dependencies/zycore/include",
		},
	}})
}
