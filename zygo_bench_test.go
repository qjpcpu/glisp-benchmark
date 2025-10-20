package main

import (
	"fmt"
	"testing"

	"github.com/glycerine/zygomys/v9/zygo"
)

func BenchmarkFactorial_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	_, err := env.EvalString(`(defn factorial[n]
(cond (== 1 n) n (* n (factorial (- n 1))))
)`)
	MustSuccess(t, err)
	v, _ := env.FindObject("factorial")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		res, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpInt{Val: 10}})
		MustSuccess(t, err)
		MustEqualInt64(t, 3628800, res.(*zygo.SexpInt).Val)
	}
}

func BenchmarkRegexpMatch_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	env.ImportRegex()
	_, err := env.EvalString(`
(def re (regexpCompile "^\\d{3}\\d{4}\\d{4}$"))
(defn testPhoneNumber [n]
(regexpMatch re n))`)
	MustSuccess(t, err)
	v, _ := env.FindObject("testPhoneNumber")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		res, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpStr{S: "15744882345"}})
		MustSuccess(t, err)
		num, ok := res.(*zygo.SexpBool)
		MustTrue(t, ok)
		MustTrue(t, num.Val)
	}
}

func BenchmarkComplexCondition_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	_, err := env.EvalString(`
(defn complex_condition [n]
  (cond
    (and (>= n 0) (<= n 10)) "low"
    (and (> n 10) (<= n 20)) "medium"
    (and (> n 20) (<= n 30)) "high"
    "unknown"))
`)
	MustSuccess(t, err)
	v, _ := env.FindObject("complex_condition")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		res, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpInt{Val: 15}})
		MustSuccess(t, err)
		s, ok := res.(*zygo.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "medium", s.S)
	}
}

func BenchmarkHashWrite_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	_, err := env.EvalString(`
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
(defn set_in_hash [key value] (hset m key value))
`)
	MustSuccess(t, err)
	v, _ := env.FindObject("set_in_hash")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		_, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpStr{S: "key1"}, &zygo.SexpStr{S: "new_value"}})
		MustSuccess(t, err)
	}
}

func BenchmarkHashDelete_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	_, err := env.EvalString(`
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
(defn delete_from_hash [key] (hdel m key))
`)
	MustSuccess(t, err)
	v, _ := env.FindObject("delete_from_hash")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		_, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpStr{S: "key1"}})
		MustSuccess(t, err)
	}
}

func BenchmarkHashAccess_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	_, err := env.EvalString(`
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
(defn get_from_hash [key] (hget m key))
`)
	MustSuccess(t, err)
	v, _ := env.FindObject("get_from_hash")
	fn := v.(*zygo.SexpFunction)
	for i := 0; i < t.N; i++ {
		res, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpStr{S: "key5"}})
		MustSuccess(t, err)
		MustEqual(t, "value5", res.(*zygo.SexpStr).S)
	}
}

func BenchmarkJSONParseAndModify_zygo(t *testing.B) {
	env := zygo.NewZlisp()
	env.AddFunction("parseJSON",
		func(env *zygo.Zlisp, name string, args []zygo.Sexp) (zygo.Sexp, error) {
			if len(args) != 1 {
				return zygo.SexpNull, fmt.Errorf("需要1个参数")
			}

			jsonStr, ok := args[0].(*zygo.SexpStr)
			if !ok {
				return zygo.SexpNull, fmt.Errorf("参数必须是字符串")
			}
			return zygo.JsonToSexp([]byte(jsonStr.S), env)
		})
	_, err := env.EvalString(`
(defn parse_and_modify [json_str]
    (def data (parseJSON json_str))
    (hset data "name" "new_name")
    (hget data "name"))
`)
	MustSuccess(t, err)
	v, _ := env.FindObject("parse_and_modify")
	fn := v.(*zygo.SexpFunction)
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`
	for i := 0; i < t.N; i++ {
		res, err := env.Apply(fn, []zygo.Sexp{&zygo.SexpStr{S: jsonStr}})
		MustSuccess(t, err)
		s, ok := res.(*zygo.SexpStr)
		MustTrue(t, ok)
		MustEqual(t, "new_name", s.S)
	}
}
