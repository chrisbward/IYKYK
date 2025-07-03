package contentcontroller

import (
	"errors"
	"strings"
	"testing"

	"github.com/chrisbward/IYKYK/pkg/controllers/stripcontentcontroller"
	"github.com/chrisbward/IYKYK/pkg/entities"
	mockedinterfaces "github.com/chrisbward/IYKYK/pkg/entities/_mocks"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCleanContentAutomatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// var mockedStripperController entities.IStripContentController
	mockedStripperController := mockedinterfaces.NewMockIStripContentController(ctrl)

	tests := []struct {
		name            string
		options         ContentControllerOptions
		useMockStripper bool
		input           string
		mockSetup       func()
		wantOutput      string
		wantErr         bool
	}{
		{
			name:            "empty input returns error",
			options:         ContentControllerOptions{},
			useMockStripper: true,
			input:           "",
			mockSetup: func() {
				// no calls expected because it should return early
			},
			wantErr: true,
		},
		{
			name: "clean string as per readme documentation",
			options: ContentControllerOptions{
				StripEmDash:       true,
				StripEmoji:        true,
				StripAngledQuotes: true,
			},
			useMockStripper: false,
			input:           "Here is some content - I wish to be “cleaned”. It’s very useful 🚀 for certain purposes ",
			mockSetup: func() {
				// nothing here as we use real stripper for demo
			},
			wantOutput: `Here is some content - I wish to be "cleaned". It's very useful for certain purposes`,
			wantErr:    false,
		},

		{
			name: "strip emoji enabled calls DoCleanWithStripFunctions",
			options: ContentControllerOptions{
				StripEmoji: true,
			},
			useMockStripper: true,
			input:           "some emoji input",
			mockSetup: func() {
				mockedStripperController.EXPECT().
					StripInputOfEmoji("some emoji input").
					Return("cleaned emoji output", nil)
			},
			wantOutput: "cleaned emoji output",
			wantErr:    false,
		},
		{
			name: "strip emoji and em dash enabled",
			options: ContentControllerOptions{
				StripEmoji:  true,
				StripEmDash: true,
			},
			useMockStripper: true,
			input:           "some input with emoji 🌍🚀 and em dash —",
			mockSetup: func() {
				mockedStripperController.EXPECT().
					StripInputOfEmoji("some input with emoji 🌍🚀 and em dash —").
					Return("some input with emoji and em dash —", nil).
					Times(1)
				mockedStripperController.EXPECT().
					StripInputOfEmDash("some input with emoji and em dash —").
					Return("some input with emoji and em dash", nil).
					Times(1)
			},
			wantOutput: "some input with emoji and em dash",
			wantErr:    false,
		},
		{
			name: "strip emoji, angled quotes and em dash enabled",
			options: ContentControllerOptions{
				StripEmoji:        true,
				StripEmDash:       true,
				StripAngledQuotes: true,
			},
			useMockStripper: true,
			input:           "some input with emoji 🌍🚀 and em dash — and some angled “quotes”",
			mockSetup: func() {
				mockedStripperController.EXPECT().
					StripInputOfEmoji("some input with emoji 🌍🚀 and em dash — and some angled “quotes”").
					Return("some input with emoji and em dash — and some angled “quotes”", nil).
					Times(1)
				mockedStripperController.EXPECT().
					StripInputOfEmDash("some input with emoji and em dash — and some angled “quotes”").
					Return("some input with emoji and em dash - and some angled “quotes”", nil).
					Times(1)
				mockedStripperController.EXPECT().
					StripInputOfAngledQuotes("some input with emoji and em dash - and some angled “quotes”").
					Return(`some input with emoji and em dash - and some angled "quotes"`, nil).
					Times(1)
			},
			wantOutput: `some input with emoji and em dash - and some angled "quotes"`,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup()
			var stripContentController entities.IStripContentController

			if !tt.useMockStripper {
				stripContentController, _ = stripcontentcontroller.NewStripContentController()
			} else {
				stripContentController = mockedStripperController
			}

			cc := &ContentController{
				Options:                tt.options,
				StripContentController: stripContentController,
			}

			got, err := cc.CleanContentAutomatic(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got error: %v", tt.wantErr, err)
			}
			if got != tt.wantOutput {
				t.Errorf("expected output %q, got %q", tt.wantOutput, got)
			}
		})
	}
}
func TestDoCleanWithStripFunctions(t *testing.T) {
	cc := &ContentController{}

	stripContentController, _ := stripcontentcontroller.NewStripContentController()

	tests := []struct {
		name       string
		input      string
		funcs      []entities.StripFunction
		want       string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "Empty input string returns error",
			input:      "",
			funcs:      []entities.StripFunction{func(s string) (string, error) { return s, nil }},
			wantErr:    true,
			errMessage: "input is empty",
		},
		{
			name:       "No strip functions returns error",
			input:      "Hello 🌍, world!",
			funcs:      nil,
			wantErr:    true,
			errMessage: "must provide strip functions",
		},
		{
			name:  "All strip functions succeed",
			input: "Hello 🌍, world!",
			funcs: []entities.StripFunction{
				stripContentController.StripInputOfEmoji,
				func(s string) (string, error) {
					return strings.ReplaceAll(s, "!", ""), nil
				},
			},
			want:    "Hello, world",
			wantErr: false,
		},
		{
			name:  "Fails on second function and short-circuits",
			input: "start",
			funcs: []entities.StripFunction{
				func(s string) (string, error) {
					return strings.ReplaceAll(s, "⚠️", ""), nil
				},
				func(s string) (string, error) {
					return "", errors.New("strip failed")
				},
				func(s string) (string, error) {
					return strings.ReplaceAll(s, "error", ""), nil
				},
			},
			wantErr:    true,
			errMessage: "strip failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cc.DoCleanWithStripFunctions(tt.input, tt.funcs...)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMessage)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
