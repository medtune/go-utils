package vector

import (
	"reflect"
	"testing"
)

func TestVector_Scalar(t *testing.T) {
	type fields struct {
		Dim  int
		Data []float64
	}
	type args struct {
		k float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []float64
	}{
		{
			name: "zero scalar",
			fields: fields{
				Dim:  2,
				Data: []float64{3., 4.},
			},
			args: args{
				k: 0.,
			},
			want: []float64{0., 0.},
		},
		{
			name: "one scalar",
			fields: fields{
				Dim:  2,
				Data: []float64{3., 4.},
			},
			args: args{
				k: 1.,
			},
			want: []float64{3., 4.},
		},
		{
			name: "negative scalar",
			fields: fields{
				Dim:  2,
				Data: []float64{3., 4.},
			},
			args: args{
				k: -3.,
			},
			want: []float64{-9., -12.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				Dim:  tt.fields.Dim,
				Data: tt.fields.Data,
			}
			v.Scalar(tt.args.k)
			if equal := reflect.DeepEqual(v.Data, tt.want); !equal {
				t.Errorf("Vector.Scalar(%v) expected = %v got = %v", tt.args.k, tt.want, v.Data)
			}
		})
	}
}

func TestVector_Add(t *testing.T) {
	type fields struct {
		Dim  int
		Data []float64
	}
	type args struct {
		vector *Vector
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []float64
	}{
		{
			name: "zero vector",
			fields: fields{
				Dim:  2,
				Data: []float64{3., 4.},
			},
			args: args{
				vector: &Vector{Dim: 2, Data: []float64{0., 0.}},
			},
			want: []float64{0., 0.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				Dim:  tt.fields.Dim,
				Data: tt.fields.Data,
			}
			if err := v.Add(tt.args.vector); (err != nil) != tt.wantErr {
				t.Errorf("Vector.Add(%v) error = %v, wantErr %v", tt.args.vector, err, tt.wantErr)
			}
			if equal := reflect.DeepEqual(v.Data, tt.want); !equal {
				t.Errorf("Vector.Add(%v) expected = %v got = %v", tt.args.vector, tt.want, v.Data)
			}
		})
	}
}

func TestVector_Sub(t *testing.T) {
	type fields struct {
		Dim  int
		Data []float64
	}
	type args struct {
		vector *Vector
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				Dim:  tt.fields.Dim,
				Data: tt.fields.Data,
			}
			if err := v.Sub(tt.args.vector); (err != nil) != tt.wantErr {
				t.Errorf("Vector.Sub(%v) error = %v, wantErr %v", tt.args.vector, err, tt.wantErr)
			}
		})
	}
}

func TestVector_Mul(t *testing.T) {
	type fields struct {
		Dim  int
		Data []float64
	}
	type args struct {
		vector *Vector
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				Dim:  tt.fields.Dim,
				Data: tt.fields.Data,
			}
			if err := v.Mul(tt.args.vector); (err != nil) != tt.wantErr {
				t.Errorf("Vector.Mul(%v) error = %v, wantErr %v", tt.args.vector, err, tt.wantErr)
			}
		})
	}
}
