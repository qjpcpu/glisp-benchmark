package main

import (
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/dop251/goja"
)

func BenchmarkFactorial_goja(t *testing.B) {
	const SCRIPT = `
function factorial(n) {
    return n === 1 ? n : n * factorial(--n);
}
`

	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	fac, ok := goja.AssertFunction(vm.Get("factorial"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := fac(goja.Undefined(), vm.ToValue(10))
		MustSuccess(t, err)
		MustEqualInt64(t, 3628800, res.ToInteger())
	}
}

func BenchmarkRegexpMatch_goja(t *testing.B) {
	const SCRIPT = `
		function testPhoneNumber(phone) {
			return test(phone, "^\\d{3}\\d{4}\\d{4}$");
		}
`

	vm := goja.New()
	var cache sync.Map
	vm.Set("test", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		pattern := call.Argument(1).String()
		var re *regexp.Regexp
		val, ok := cache.Load(pattern)
		if ok {
			re = val.(*regexp.Regexp)
		} else {
			re0, err := regexp.Compile(pattern)
			if err != nil {
				panic(err)
			}
			re = re0
			cache.Store(pattern, re0)
		}
		return vm.ToValue(re.MatchString(text))
	})
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	fac, ok := goja.AssertFunction(vm.Get("testPhoneNumber"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := fac(goja.Undefined(), vm.ToValue("15744882345"))
		MustSuccess(t, err)
		MustTrue(t, res.ToBoolean())
	}
}

func BenchmarkComplexCondition_goja(t *testing.B) {
	const SCRIPT = `
function complex_condition(n) {
  if (n >= 0 && n <= 10) {
    return "low";
  } else if (n > 10 && n <= 20) {
    return "medium";
  } else if (n > 20 && n <= 30) {
    return "high";
  } else {
    return "unknown";
  }
}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("complex_condition"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := f(goja.Undefined(), vm.ToValue(15))
		MustSuccess(t, err)
		MustEqual(t, "medium", res.String())
	}
}

func BenchmarkFormatTime_goja(t *testing.B) {
	const SCRIPT = `
function formatTime(t) {
	return format(t, "2006-01-02T15:04:05Z", "2006年01月02日 15时04分05秒");
}
`
	vm := goja.New()
	vm.Set("format", func(call goja.FunctionCall) goja.Value {
		val := call.Argument(0).String()
		layout := call.Argument(1).String()
		newLayout := call.Argument(2).String()
		t, err := time.Parse(layout, val)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(t.Format(newLayout))
	})
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("formatTime"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := f(goja.Undefined(), vm.ToValue("2006-01-02T15:04:05Z"))
		MustSuccess(t, err)
		MustEqual(t, "2006年01月02日 15时04分05秒", res.String())
	}
}

func BenchmarkHashAccess_goja(t *testing.B) {
	const SCRIPT = `
const m = {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
    "key4": "value4",
    "key5": "value5",
    "key6": "value6",
    "key7": "value7",
    "key8": "value8",
    "key9": "value9",
    "key10": "value10"
};

function get_from_hash(key) {
    return m[key];
}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("get_from_hash"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		res, err := f(goja.Undefined(), vm.ToValue("key5"))
		MustSuccess(t, err)
		MustEqual(t, "value5", res.String())
	}
}

func BenchmarkHashWrite_goja(t *testing.B) {
	const SCRIPT = `
const m = {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
    "key4": "value4",
    "key5": "value5",
    "key6": "value6",
    "key7": "value7",
    "key8": "value8",
    "key9": "value9",
    "key10": "value10"
};

function set_in_hash(key, value) {
    m[key] = value;
}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("set_in_hash"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		_, err := f(goja.Undefined(), vm.ToValue("key1"), vm.ToValue("new_value"))
		MustSuccess(t, err)
	}
}

func BenchmarkHashDelete_goja(t *testing.B) {
	const SCRIPT = `
const m = {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
    "key4": "value4",
    "key5": "value5",
    "key6": "value6",
    "key7": "value7",
    "key8": "value8",
    "key9": "value9",
    "key10": "value10"
};

function delete_from_hash(key) {
    delete m[key];
}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("delete_from_hash"))
	MustTrue(t, ok)
	for i := 0; i < t.N; i++ {
		_, err := f(goja.Undefined(), vm.ToValue("key1"))
		MustSuccess(t, err)
	}
}

func BenchmarkJSONParseAndModify_goja(t *testing.B) {
	const SCRIPT = `
function parse_and_modify(json_str) {
    let data = JSON.parse(json_str);
    data.name = "new_name";
	return data.name;
}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	MustSuccess(t, err)
	f, ok := goja.AssertFunction(vm.Get("parse_and_modify"))
	MustTrue(t, ok)
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`
	for i := 0; i < t.N; i++ {
		res, err := f(goja.Undefined(), vm.ToValue(jsonStr))
		MustSuccess(t, err)
		MustEqual(t, "new_name", res.String())
	}
}
