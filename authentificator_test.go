package main

import "testing"

func Test_writeStringToFile(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeStringToFile(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("writeStringToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
