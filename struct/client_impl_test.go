package main

import "testing"

func TestClientImple_HttpClient(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		my      ClientImple
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "client test ok",
			my:      ClientImple{},
			args:    args{msg: "string"},
			want:    "}string{",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my := ClientImple{}
			got, err := my.HttpClient(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientImple.HttpClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientImple.HttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientImple_ResponseGen(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		my      ClientImple
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "client test",
			my:      ClientImple{},
			args:    args{msg: "string"},
			want:    ">string<",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my := ClientImple{}
			got, err := my.ResponseGen(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientImple.ResponseGen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientImple.ResponseGen() = %v, want %v", got, tt.want)
			}
		})
	}
}
