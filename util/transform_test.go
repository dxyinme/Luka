package util

import (
	"log"
	"math"
	"reflect"
	"testing"
)

func TestTransformList(t *testing.T) {

}

func Test_transform(t *testing.T) {
	type args struct {
		param    string
		typeName string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1 - test - string",
			args: args{
				param: "1",
				typeName: "string",
			},
			want: "1",
			wantErr: false,
		},
		{
			name: "47 - test - int",
			args: args{
				param: "47",
				typeName: "int",
			},
			want: 47,
			wantErr: false,
		},
		{
			name: "3.571E+00 - test - float32",
			args: args{
				param: "3.571E+00",
				typeName: "float32",
			},
			want: 3.571,
			wantErr: false,
		},
		{
			name: "3.571E+00 - test - float64",
			args: args{
				param: "3.571E+00",
				typeName: "float64",
			},
			want: 3.571,
			wantErr: false,
		},
	}
	eps32 := 1e-6
	eps64 := 1e-10
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transform(tt.args.param, tt.args.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.typeName == "float32" {
				if math.Abs(got.(float64) - tt.want.(float64)) < eps32 {
					return
				}else{
					t.Errorf("transform() got = %v, want %v", got, tt.want)
				}
			}
			if tt.args.typeName == "float64" {
				if math.Abs(got.(float64) - tt.want.(float64)) < eps64 {
					return
				}else{
					t.Errorf("transform() got = %v, want %v", got, tt.want)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransIntoReflect(t *testing.T) {

}

func TestTransIntoString(t *testing.T) {

}

func TestTransformList1(t *testing.T) {

}

func Test_transReflectIntoString(t *testing.T) {
	o := 3.571
	resString,typeString := transReflectIntoString(reflect.ValueOf(o))
	log.Printf("string : %s type : %s",resString,typeString)

}

