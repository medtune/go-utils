package errors

import (
	"fmt"
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
		{
			"empty error list",
			args{[]error{}},
			[]string{},
		},
		{
			"non empty error list",
			args{[]error{fmt.Errorf("test error 1"), fmt.Errorf("test error 2")}},
			[]string{"test error 1", "test error 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorsToStrings(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("------- %s", got)
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
		{
			"empty error",
			fields{
				"",
				"",
				time.Now(),
				[]error{},
			},
			"",
		},
		{
			"simple error",
			fields{
				"unexpected max value in this case",
				"max-value-exceeded",
				time.Now(),
				[]error{},
			},
			"unexpected max value in this case",
		},
		{
			"composed error",
			fields{
				"unexpected values for all sub types:",
				"initilization error",
				time.Now(),
				[]error{fmt.Errorf("field1: 10"), fmt.Errorf("field2: 20"), fmt.Errorf("field3: 30")},
			},
			"unexpected values for all sub types:\n\tfield1: 10\n\tfield2: 20\n\tfield3: 30\n",
		},
		{
			"composed error without message",
			fields{
				"",
				"initilization error",
				time.Now(),
				[]error{fmt.Errorf("field1: 10"), fmt.Errorf("field2: 20"), fmt.Errorf("field3: 30")},
			},
			"\n\tfield1: 10\n\tfield2: 20\n\tfield3: 30\n",
		},
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
		{
			"error is empty",
			fields{},
			true,
		},
		{
			"error not empty",
			fields{SubErrors: []error{fmt.Errorf("easy error")}},
			false,
		},
		{
			"error not empty",
			fields{Message: "i am not empty"},
			false,
		},
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

func TestNew(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			"new error",
			args{"simple error"},
			&Error{Message: "simple error"},
		},
		{
			"new empty error",
			args{""},
			&Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.err); !reflect.DeepEqual(got.Error(), tt.want.Error()) {
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
		{
			"no errors case",
			args{[]error{}},
			&Error{},
		},
		{
			"some errors",
			args{[]error{fmt.Errorf("error 1"), fmt.Errorf("error 2")}},
			&Error{SubErrors: []error{fmt.Errorf("error 1"), fmt.Errorf("error 2")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.errors...); !reflect.DeepEqual(got.Error(), tt.want.Error()) {

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
		{
			"normal case",
			args{format: "error proc: %v", a: []interface{}{5}},
			&Error{Message: "error proc: 5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Newf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got.Error(), tt.want.Error()) {
				t.Errorf("Newf() = %v, want %v", got, tt.want)
			}
		})
	}
}
