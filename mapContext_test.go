// mapContext_text
package context

import (
	goCtx "context"
	"testing"
)

func TestMapContext(t *testing.T) {
	ctx := WithValue(goCtx.Background(), "key1", "value1")
	mapCtx, ok := ctx.(MapContext)
	if !ok {
		t.Errorf("Assert Type Fail: %s,\n ctx.Value(%v) => %s", ctx,
			"key1", ctx.Value("key1"))
		return
	} else {
		t.Logf("Assert Type Succ: %s,\n ctx.Value(%v) => %s", mapCtx,
			"key1", mapCtx.GetValue("key1"))
	}

	mapCtx.SetValue("key2", "value2")
	if mapCtx.GetValue("key2") == "value2" {
		t.Logf("Succ,\n mapCtx.SetValue(%v) => %s",
			"key2", mapCtx.GetValue("key2"))
	} else {
		t.Errorf("Fail,\n mapCtx.SetValue(%v) => %s",
			"key2", mapCtx.GetValue("key2"))
		return
	}

	curCtx := goCtx.WithValue(mapCtx, "key1", "unknow")
	if curCtx.Value("key1") == "unknow" {
		t.Logf("Succ, %v\n curCtx.WithValue(%v, %v) => %s\n", curCtx,
			"key1", "unknow", curCtx.Value("key1"))
	} else {
		t.Errorf("Fail, %v\n curCtx.WithValue(%v, %v) => %s", curCtx,
			"key1", "unknow", curCtx.Value("key1"))
		return
	}

	if mapCtx.GetValue("key1") == "value1" {
		t.Logf("Succ,\n mapCtx.SetValue(%v) => %s",
			"key1", mapCtx.GetValue("key1"))
	} else {
		t.Errorf("Fail,\n mapCtx.SetValue(%v) => %s",
			"key1", mapCtx.GetValue("key1"))
		return
	}
}
