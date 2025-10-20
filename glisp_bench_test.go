package main

import (
	"bytes"
	"testing"

	"github.com/qjpcpu/glisp"
	ext "github.com/qjpcpu/glisp/extensions"
)

func BenchmarkFactorial_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	vm.SourceStream(bytes.NewBufferString(`(defn factorial[n]
(cond (= 1 n) n (* n (factorial (- n 1))))
)`))
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("factorial", glisp.MakeArgs(glisp.NewSexpInt(10)))
		MustSuccess(t, err)
		num, ok := v.(glisp.SexpInt)
		MustTrue(t, ok)
		MustEqualInt64(t, 3628800, num.ToInt64())
	}
}

func BenchmarkRegexpMatch_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportRegex(vm)
	err := vm.SourceStream(bytes.NewBufferString(`(defn testPhoneNumber[n]
(regexp/match "^\\d{3}\\d{4}\\d{4}$" n)
)`))
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("testPhoneNumber", glisp.MakeArgs(glisp.SexpStr("15744882345")))
		MustSuccess(t, err)
		num, ok := v.(glisp.SexpBool)
		MustTrue(t, ok)
		MustTrue(t, bool(num))
	}
}

func BenchmarkComplexCondition_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	vm.SourceStream(bytes.NewBufferString(`
(defn complex-condition [n]
  (cond
    (and (>= n 0) (<= n 10)) "low"
    (and (> n 10) (<= n 20)) "medium"
    (and (> n 20) (<= n 30)) "high"
    "unknown"))
`))
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("complex-condition", glisp.MakeArgs(glisp.NewSexpInt(15)))
		MustSuccess(t, err)
		s, ok := v.(glisp.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "medium", string(s))
	}
}

func BenchmarkFormatTime_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportTime(vm)
	err := vm.SourceStream(bytes.NewBufferString(`(defn formatTime [t]
  (time/format (time/parse t "2006-01-02T15:04:05Z") "2006年01月02日 15时04分05秒")
)`))
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("formatTime", glisp.MakeArgs(glisp.SexpStr("2006-01-02T15:04:05Z")))
		MustSuccess(t, err)
		s, ok := v.(glisp.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "2006年01月02日 15时04分05秒", string(s))
	}
}

func BenchmarkHashWrite_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	err := vm.SourceStream(bytes.NewBufferString(`
(def m (hash
    "key1" "value1"
    "key2" "value2"
    "key3" "value3"
    "key4" "value4"
    "key5" "value5"
    "key6" "value6"
    "key7" "value7"
    "key8" "value8"
    "key9" "value9"
    "key10" "value10"))
(defn set-in-hash [key value] (hset! m key value))
`))
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		// Overwrite an existing key
		_, err := vm.ApplyByName("set-in-hash", glisp.MakeArgs(glisp.SexpStr("key1"), glisp.SexpStr("new_value")))
		MustSuccess(t, err)
	}
}

func BenchmarkHashDelete_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	err := vm.SourceStream(bytes.NewBufferString(`
(def m (hash
    "key1" "value1"
    "key2" "value2"
    "key3" "value3"
    "key4" "value4"
    "key5" "value5"
    "key6" "value6"
    "key7" "value7"
    "key8" "value8"
    "key9" "value9"
    "key10" "value10"))
(defn delete-from-hash [key] (hdel! m key))
`))
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		_, err := vm.ApplyByName("delete-from-hash", glisp.MakeArgs(glisp.SexpStr("key1")))
		MustSuccess(t, err)
	}
}

func BenchmarkHashAccess_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportCoreUtils(vm)
	err := vm.SourceStream(bytes.NewBufferString(`
(def m (hash
    "key1" "value1"
    "key2" "value2"
    "key3" "value3"
    "key4" "value4"
    "key5" "value5"
    "key6" "value6"
    "key7" "value7"
    "key8" "value8"
    "key9" "value9"
    "key10" "value10"))
(defn get-from-hash [key] (hget m key))
`))
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("get-from-hash", glisp.MakeArgs(glisp.SexpStr("key5")))
		MustSuccess(t, err)
		MustEqual(t, "value5", string(v.(glisp.SexpStr)))
	}
}

func BenchmarkJSONParseAndModify_glisp(t *testing.B) {
	vm := glisp.New()
	ext.ImportJSON(vm)
	vm.SourceStream(bytes.NewBufferString(`
(defn parse_and_modify [json_str]
    (def data (json/parse json_str))
    (hset! data "name" "new_name")
    (hget data "name"))
`))
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`
	for i := 0; i < t.N; i++ {
		v, err := vm.ApplyByName("parse_and_modify", glisp.MakeArgs(glisp.SexpStr(jsonStr)))
		MustSuccess(t, err)
		num, ok := v.(glisp.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "new_name", string(num))
	}
}
