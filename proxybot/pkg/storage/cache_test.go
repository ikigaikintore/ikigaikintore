package storage

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	newCache := NewCache()
	if newCache == nil {
		t.Error("newCache is nil")
	}
}

func Test_cache_Get(t *testing.T) {
	type fields struct {
		key  string
		data any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		want1  bool
	}{
		{
			name: "set and get data hit",
			fields: fields{
				key: "location",
				data: struct {
					lat  float64
					long float64
				}{
					lat:  21.908790,
					long: 123.987987,
				},
			},
			args: args{key: "location"},
			want: struct {
				lat  float64
				long float64
			}{
				lat:  21.908790,
				long: 123.987987,
			},
			want1: true,
		},
		{
			name: "get data no hit",
			fields: fields{
				key: "location1",
				data: struct {
					lat  float64
					long float64
				}{
					lat:  21.908790,
					long: 123.987987,
				},
			},
			args:  args{key: "location"},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache()
			c.Set(tt.fields.key, tt.fields.data)
			got, got1 := c.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
