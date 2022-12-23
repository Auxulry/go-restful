package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type typeMessageColorized struct {
	name, request, color string
}

func TestMessageColorize(t *testing.T) {
	tests := []typeMessageColorized{
		{
			name:    "MessageColorized('Green', Green)",
			request: "Green",
			color:   Green,
		},
		{
			name:    "MessageColorized('Orange', Orange)",
			request: "Orange",
			color:   Orange,
		},
		{
			name:    "MessageColorized('Blue', Blue)",
			request: "Blue",
			color:   Blue,
		},
		{
			name:    "MessageColorized('Red', Red)",
			request: "Red",
			color:   Red,
		},
		{
			name:    "MessageColorized('DefaultColor', DefaultColor)",
			request: "DefaultColor",
			color:   DefaultColor,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := fmt.Sprintf("%v%v%v", test.color, test.request, DefaultColor)
			result := MessageColorized(test.color, test.request)

			assert.Equal(t, expected, result)
		})
	}
}
