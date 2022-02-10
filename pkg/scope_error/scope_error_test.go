package scope_error

import (
	"fmt"
	"testing"
)

func TestScopeError_Is(t *testing.T) {
	type fields struct {
		msg   string
		cause error
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"without cause",
			fields{
				msg: "some error scope",
			},
			args{err: New("some error scope")},
			true,
		},
		{
			"with cause",
			fields{
				msg:   "some error scope",
				cause: fmt.Errorf("oops"),
			},
			args{err: New("some error scope")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ScopeError{
				scope: tt.fields.msg,
				cause: tt.fields.cause,
			}
			if got := s.Is(tt.args.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
