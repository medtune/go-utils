package errors

import (
	"reflect"
	"testing"
	"time"
)

func Test_errorsToStrings(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorsToStrings(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorsToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Message   string
		Type      string
		Time      time.Time
		SubErrors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message:   tt.fields.Message,
				Type:      tt.fields.Type,
				Time:      tt.fields.Time,
				SubErrors: tt.fields.SubErrors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Append(t *testing.T) {
	type fields struct {
		Message   string
		Type      string
		Time      time.Time
		SubErrors []error
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message:   tt.fields.Message,
				Type:      tt.fields.Type,
				Time:      tt.fields.Time,
				SubErrors: tt.fields.SubErrors,
			}
			e.Append(tt.args.err)
		})
	}
}

func TestError_Appendf(t *testing.T) {
	type fields struct {
		Message   string
		Type      string
		Time      time.Time
		SubErrors []error
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message:   tt.fields.Message,
				Type:      tt.fields.Type,
				Time:      tt.fields.Time,
				SubErrors: tt.fields.SubErrors,
			}
			e.Appendf(tt.args.format, tt.args.a...)
		})
	}
}

func TestError_IsEmpty(t *testing.T) {
	type fields struct {
		Message   string
		Type      string
		Time      time.Time
		SubErrors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message:   tt.fields.Message,
				Type:      tt.fields.Type,
				Time:      tt.fields.Time,
				SubErrors: tt.fields.SubErrors,
			}
			if got := e.IsEmpty(); got != tt.want {
				t.Errorf("Error.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNil(t *testing.T) {
	tests := []struct {
		name string
		want *Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewList(t *testing.T) {
	type args struct {
		errors []error
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.errors...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewf(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Newf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Newf() = %v, want %v", got, tt.want)
			}
		})
	}
}
