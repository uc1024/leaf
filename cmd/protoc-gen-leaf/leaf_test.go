package main

import "testing"

func Test_parsePathToCodesV1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    []uint16
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test path",
			path:    "/request/102/response/103",
			want:    []uint16{102, 103},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parsePathToCodesV1(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parsePathToCodesV1() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parsePathToCodesV1() succeeded unexpectedly")
			}
			t.Log(got)
		})
	}
}
