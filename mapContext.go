// mapContext
package context

import (
	goCtx "context"
	"fmt"
	"reflect"
)

type MapContext interface {
	goCtx.Context
	SetValue(key, value interface{})
	GetValue(key interface{}) interface{}
}

type mapContext struct {
	goCtx.Context
	mapValues map[interface{}]interface{}
}

// 获取一个key的value; 如果mapContext.Context是goCtc.Background(),
// 则mapContext的优先级最低, 通过goCtc.WithValue()返回context的Value(key)
// 输出会覆盖mapContext的同名key-value; 反之当前mapContext.Context是最外层
// Context的mapContext会覆盖父级context的Value(key)的结果.
func (m *mapContext) Value(key interface{}) interface{} {
	if val, ok := m.mapValues[key]; ok {
		return val
	}
	return m.Context.Value(key)
}

func WithValue(parent goCtx.Context, key, val interface{}) goCtx.Context {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	if mapCtx, ok := parent.(MapContext); ok {
		mapCtx.SetValue(key, val)
		return parent
	} else {
		m := &mapContext{
			Context:   parent,
			mapValues: make(map[interface{}]interface{}),
		}
		m.mapValues[key] = val
		return m
	}
}

// set map value by key, if value is nil then remote key
func (m *mapContext) SetValue(key, value interface{}) {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	if value == nil {
		delete(m.mapValues, key)
	} else {
		m.mapValues[key] = value
	}
}

// get map value by key, if key is nil then panic
func (m *mapContext) GetValue(key interface{}) interface{} {
	if key == nil {
		panic("nil key")
	}
	return m.mapValues[key]
}

func (m *mapContext) String() string {
	return fmt.Sprintf("%v.MapContext", m.Context)
}
