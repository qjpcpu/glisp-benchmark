package main

import (
	"encoding/json"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/Shopify/go-lua"
)

func BenchmarkFactorial_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)

	script := `
  function factorial(n)
    if n == 1 then
      return 1
    end
    return n * factorial(n-1)
  end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("factorial")
		l.PushInteger(10)

		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)

		i, _ := l.ToInteger(-1)
		MustEqualInt64(t, 3628800, int64(i))
		l.Pop(1)
	}
}

func BenchmarkHashWrite_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	script := `
local m = {
    ["key1"] = "value1",
    ["key2"] = "value2",
    ["key3"] = "value3",
    ["key4"] = "value4",
    ["key5"] = "value5",
    ["key6"] = "value6",
    ["key7"] = "value7",
    ["key8"] = "value8",
    ["key9"] = "value9",
    ["key10"] = "value10"
}

function set_in_hash(key, value)
    m[key] = value
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("set_in_hash")
		l.PushString("key1")
		l.PushString("new_value")
		err = l.ProtectedCall(2, 0, 0) // 2 arguments, 0 return values
		MustSuccess(t, err)
		// No return value to check, just ensure no error
	}
}

func BenchmarkHashDelete_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	script := `
local m = {
    ["key1"] = "value1",
    ["key2"] = "value2",
    ["key3"] = "value3",
    ["key4"] = "value4",
    ["key5"] = "value5",
    ["key6"] = "value6",
    ["key7"] = "value7",
    ["key8"] = "value8",
    ["key9"] = "value9",
    ["key10"] = "value10"
}

function delete_from_hash(key)
    m[key] = nil
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("delete_from_hash")
		l.PushString("key1")
		err = l.ProtectedCall(1, 0, 0) // 1 argument, 0 return values
		MustSuccess(t, err)
	}
}

func BenchmarkRegexpMatch_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	var cache sync.Map
	l.Register("test", func(l *lua.State) int {
		text := lua.CheckString(l, 1)
		pattern := lua.CheckString(l, 2)

		var re *regexp.Regexp
		val, ok := cache.Load(pattern)
		if ok {
			re = val.(*regexp.Regexp)
		} else {
			re0, err := regexp.Compile(pattern)
			if err != nil {
				l.PushBoolean(false)
				l.PushString(err.Error())
				return 2
			}
			re = re0
			cache.Store(pattern, re0)
		}

		l.PushBoolean(re.MatchString(text))
		return 1
	})
	script := `
  function testPhoneNumber(n)
    return test(n,"^\\d{3}\\d{4}\\d{4}$")
  end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("testPhoneNumber")
		l.PushString("15744882345")

		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)

		i := l.ToBoolean(-1)
		MustTrue(t, i)
		l.Pop(1)
	}
}

func BenchmarkComplexCondition_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	script := `
function complex_condition(n)
  if n >= 0 and n <= 10 then
    return "low"
  elseif n > 10 and n <= 20 then
    return "medium"
  elseif n > 20 and n <= 30 then
    return "high"
  else
    return "unknown"
  end
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("complex_condition")
		l.PushInteger(15)
		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)
		s, ok := l.ToString(-1)
		MustTrue(t, ok)
		MustEqual(t, "medium", s)
		l.Pop(1)
	}
}

func BenchmarkFormatTime_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	l.Register("format", func(l *lua.State) int {
		val := lua.CheckString(l, 1)
		layout := lua.CheckString(l, 2)
		newLayout := lua.CheckString(l, 3)
		t, err := time.Parse(layout, val)
		if err != nil {
			l.PushString(err.Error())
			return 1
		}
		l.PushString(t.Format(newLayout))
		return 1
	})
	script := `
function formatTime(t)
	return format(t, "2006-01-02T15:04:05Z", "2006年01月02日 15时04分05秒")
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("formatTime")
		l.PushString("2006-01-02T15:04:05Z")
		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)
		s, ok := l.ToString(-1)
		MustTrue(t, ok)
		MustEqual(t, "2006年01月02日 15时04分05秒", s)
		l.Pop(1)
	}
}

func BenchmarkHashAccess_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	script := `
local m = {
    ["key1"] = "value1",
    ["key2"] = "value2",
    ["key3"] = "value3",
    ["key4"] = "value4",
    ["key5"] = "value5",
    ["key6"] = "value6",
    ["key7"] = "value7",
    ["key8"] = "value8",
    ["key9"] = "value9",
    ["key10"] = "value10"
}

function get_from_hash(key)
    return m[key]
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)
	for i := 0; i < t.N; i++ {
		l.Global("get_from_hash")
		l.PushString("key5")
		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)
		s, ok := l.ToString(-1)
		MustTrue(t, ok)
		MustEqual(t, "value5", s)
		l.Pop(1)
	}
}

func BenchmarkJSONParseAndModify_lua(t *testing.B) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	l.Register("parse_and_modify", func(l *lua.State) int {
		jsonStr := lua.CheckString(l, 1)
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
			l.PushString(err.Error())
			return 1
		}
		data["name"] = "new_name"
		l.PushString(data["name"].(string))
		return 1
	})

	script := `
function call_parse_and_modify(json_str)
	return parse_and_modify(json_str)
end
`
	err := lua.DoString(l, script)
	MustSuccess(t, err)

	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`
	for i := 0; i < t.N; i++ {
		l.Global("call_parse_and_modify")
		l.PushString(jsonStr)
		err = l.ProtectedCall(1, 1, 0)
		MustSuccess(t, err)
		s, ok := l.ToString(-1)
		MustTrue(t, ok)
		MustEqual(t, "new_name", s)
		l.Pop(1)
	}
}
