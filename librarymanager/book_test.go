package librarymanager

import "testing"

func TestBook(t *testing.T) {
	type args struct {
		books string
	}
	tests := []struct {
		name           string
		args           args
		wantSortedBook string
		wantErr        bool
	}{
		{
			name: "Example Input",
			args: args{
				books: "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14",
			},
			wantSortedBook: "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSortedBook, err := BooksShorter(tt.args.books)
			if (err != nil) != tt.wantErr {
				t.Errorf("BooksShorter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSortedBook != tt.wantSortedBook {
				t.Errorf("BooksShorter() = %v, want %v", gotSortedBook, tt.wantSortedBook)
			}
		})
	}
}
