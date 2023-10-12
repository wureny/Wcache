package lru

import (
	"container/list"
	"reflect"
	"testing"
)

// TODO 写完测试
func TestCache_Add(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	type args struct {
		key   string
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				maxBytes:  0,
				nbytes:    0,
				ll:        nil,
				cache:     nil,
				OnEvicted: nil,
			},
			args: args{
				key:   "",
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue Value
		wantOk    bool
	}{
		// TODO: Add test cases.
		{
			name: "1element",
			fields: fields{
				maxBytes:  0,
				nbytes:    0,
				ll:        nil,
				cache:     map[string]*list.Element{"1": nil},
				OnEvicted: nil,
			},
			args: args{
				key: "1",
			},
			wantValue: nil,
			wantOk:    false,
		},
		{
			name: "0element",
			fields: fields{
				maxBytes:  0,
				nbytes:    0,
				ll:        nil,
				cache:     map[string]*list.Element{},
				OnEvicted: nil,
			},
			args: args{
				key: "1",
			},
			wantValue: nil,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			gotValue, gotOk := c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestCache_Len(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "2elements",
			fields: fields{
				maxBytes:  0,
				nbytes:    0,
				ll:        l,
				cache:     nil,
				OnEvicted: nil,
			},
			want: 2,
		},
		{
			name: "0elements",
			fields: fields{
				maxBytes:  0,
				nbytes:    0,
				ll:        list.New(),
				cache:     nil,
				OnEvicted: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.RemoveOldest()
		})
	}
}

func TestNewCache(t *testing.T) {
	type args struct {
		maxBytes  int64
		onEvicted func(string, Value)
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(tt.args.maxBytes, tt.args.onEvicted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
