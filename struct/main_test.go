package main

import (
	"fmt"
	"testing"
)

func TestTestRepository_GetResponse(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "client impl test",
			fields: fields{client: ClientImple{}},
			args: args{
				msg: "aaa",
			},
			want: ">}aaa{<",
		},
		{
			name:   "Test impl test",
			fields: fields{client: TestImple{}},
			args: args{
				msg: "aaa",
			},
			want: "|=aaa=|",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := TestRepository{
				client: tt.fields.client,
			}
			if got := rep.GetResponse(tt.args.msg); got != tt.want {
				t.Errorf("TestRepository.GetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestImple struct {
}

func (my TestImple) HttpClient(msg string) (string, error) {
	fmt.Printf("TestImple ClientImple is : %s , %s", msg, my)
	return fmt.Sprintf("=%v=", msg), nil
}

func (my TestImple) ResponseGen(msg string) (string, error) {
	fmt.Printf("TestImple ResponseGen is : %s , %s", msg, my)
	return fmt.Sprintf("|%v|", msg), nil
}
