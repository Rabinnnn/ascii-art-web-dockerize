package asciiart

import "testing"

func TestArt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"valid", "Hello", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Art(tt.s, BannerArt("banners/standard.txt"))); got.String() != tt.want {
				t.Errorf("Art() = %v, want %v", got, tt.want)
			}
		})
	}
}
