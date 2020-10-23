package lostnumber

import "testing"

func TestLostNumber(t *testing.T) {
	type args struct {
		numbers string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput float64
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "Example Input 1",
			args: args{
				numbers: "23242526272830",
			},
			wantOutput: 29,
			wantErr:    false,
		},
		{
			name: "Example Input 2",
			args: args{
				numbers: "101102103104105106107108109111112113",
			},
			wantOutput: 110,
			wantErr:    false,
		},
		{
			name: "Example Input 3",
			args: args{
				numbers: "12346789",
			},
			wantOutput: 5,
			wantErr:    false,
		},
		{
			name: "Example Input 4",
			args: args{
				numbers: "79101112",
			},
			wantOutput: 8,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := LostNumber(tt.args.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("LostNumber() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
