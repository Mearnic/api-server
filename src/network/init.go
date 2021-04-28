package network

import (
	"api-server/src/common"
	"api-server/src/reflects"
	"fmt"
	"github.com/labstack/echo/v4"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/sys/unix"
	"net/http"
	"os"
	"reflect"
)

const subPackage = "src/api"

func getFuncSet() []*ast.FuncDecl {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, subPackage, nil, 0)
	if err != nil {
		os.Exit(1)
	}
	var funcs []*ast.FuncDecl
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcs = append(funcs, fn)
				}
			}
		}
	}
	return funcs
}

func InitApi(e *echo.Echo) *echo.Echo {
	var funs = getFuncSet()
	fmt.Println("准备初始化监听请求...")
	for _, funcItem := range funs {
		name := funcItem.Name.String()
		var r common.RequestDefine
		s := reflect.ValueOf(&r).Elem()
		s.Set(reflects.LoadRequest(name)[0])
		if r.RequestType == common.GET {
			e = get(e, r)
		} else if r.RequestType == common.POST {
			e = post(e, r)
		}
		fmt.Printf("%d--- 请求 %v%s\n", unix.Getpid(), r.RequestType, " "+r.RequestUrl+" 完成监听!")
	}
	fmt.Println("监听请求初始化完毕...")
	return e
}

func post(e *echo.Echo, r common.RequestDefine) *echo.Echo {
	e.POST(r.RequestUrl, func(c echo.Context) error {
		return c.String(http.StatusOK, r.Data)
	})
	return e
}

func get(e *echo.Echo, r common.RequestDefine) *echo.Echo {
	e.GET(r.RequestUrl, func(c echo.Context) error {
		return c.String(http.StatusOK, r.Data)
	})
	return e
}
