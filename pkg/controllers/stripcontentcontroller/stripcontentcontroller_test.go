package stripcontentcontroller

import "testing"

func TestStripInputOfAngledQuotes(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{
			name:      "Empty input returns error",
			input:     "",
			want:      "",
			wantError: true,
		},
		{
			name:      "Replace angled quotes",
			input:     `“Hello ‘world’” and ’quotes’`,
			want:      `"Hello 'world'" and 'quotes'`,
			wantError: false,
		},
		{
			name:      "No angled quotes returns same input",
			input:     "No angled quotes here",
			want:      "No angled quotes here",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			contentController, _ := NewStripContentController()

			got, err := contentController.StripInputOfAngledQuotes(tt.input)
			if (err != nil) != tt.wantError {
				t.Fatalf("expected error: %v, got: %v", tt.wantError, err)
			}
			if got != tt.want {
				t.Errorf("expected output: %q, got: %q", tt.want, got)
			}
		})
	}
}

func TestStripInputOfEmoji(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{
			name:      "Empty input returns error",
			input:     "",
			want:      "",
			wantError: true,
		},
		{
			name:      "Input with emojis",
			input:     "Hello 🌍🚀world!",
			want:      "Hello world!",
			wantError: false,
		},
		{
			name:      "Input without emojis",
			input:     "Just plain text.",
			want:      "Just plain text.",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cc, _ := NewStripContentController()

			got, err := cc.StripInputOfEmoji(tt.input)
			if (err != nil) != tt.wantError {
				t.Fatalf("expected error: %v, got: %v", tt.wantError, err)
			}
			if got != tt.want {
				t.Errorf("expected output: %q, got: %q", tt.want, got)
			}
		})
	}
}

func TestStripInputOfEmDash(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{
			name:      "Empty input returns error",
			input:     "",
			want:      "",
			wantError: true,
		},
		{
			name:      "StripEmDash false returns input unchanged",
			input:     "Example - test",
			want:      "Example - test",
			wantError: false,
		},
		{
			name:      "StripEmDash true replaces em-dash and normalizes spaces",
			input:     "Example—test",
			want:      "Example - test",
			wantError: false,
		},
		{
			name:      "StripEmDash true normalizes spaces around hyphen",
			input:     "Example  —   test",
			want:      "Example - test",
			wantError: false,
		},
		{
			name:      "Multiple em-dashes replaced and spaces normalized",
			input:     "Test—string—with—multiple—em-dashes",
			want:      "Test - string - with - multiple - em-dashes",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cc, _ := NewStripContentController()

			got, err := cc.StripInputOfEmDash(tt.input)
			if (err != nil) != tt.wantError {
				t.Fatalf("expected error: %v, got: %v", tt.wantError, err)
			}
			if got != tt.want {
				t.Errorf("expected output %q, got %q", tt.want, got)
			}
		})
	}
}
