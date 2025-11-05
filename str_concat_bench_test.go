package main

import (
	"bytes"
	"testing"

	"github.com/dop251/goja"
	"github.com/glycerine/zygomys/v9/zygo"
	"github.com/qjpcpu/glisp"
	ext "github.com/qjpcpu/glisp/extensions"
	"github.com/yuin/gopher-lua"
)

func BenchmarkStringConcat_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	vm.SourceStream(bytes.NewBufferString(`(defn string-concat [a b c] (concat a b c))`))
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("concat", glisp.MakeArgs(glisp.SexpStr("hello"), glisp.SexpStr(" "), glisp.SexpStr("world")))
		MustSuccess(t, err)
		s, ok := v.(glisp.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "hello world", string(s))
	}
}

func BenchmarkStringConcat_goja(t *testing.B) {
	vm := goja.New()
	_, err := vm.RunString(`
	function stringConcat(a, b, c) {
		return a + b + c;
	}
	`)
	MustSuccess(t, err)
	stringConcat, ok := goja.AssertFunction(vm.Get("stringConcat"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := stringConcat(goja.Undefined(), vm.ToValue("hello"), vm.ToValue(" "), vm.ToValue("world"))
		MustSuccess(t, err)
		MustEqual(t, "hello world", res.String())
	}
}

func BenchmarkStringConcat_lua(t *testing.B) {
	L := lua.NewState()
	defer L.Close()
	err := L.DoString(`
	function stringConcat(a, b, c)
		return a .. b .. c
	end
	`)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("stringConcat"),
			NRet:    1,
			Protect: true,
		}, lua.LString("hello"), lua.LString(" "), lua.LString("world"))
		MustSuccess(t, err)
		ret := L.Get(-1)
		L.Pop(1)
		MustEqual(t, "hello world", ret.String())
	}
}

func BenchmarkStringConcat_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	for i := 0; i < t.N; i++ {
		res, err := env.EvalString(`(concat "hello" " " "world")`)
		MustSuccess(t, err)
		s, ok := res.(*zygo.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "hello world", s.S)
	}
}
