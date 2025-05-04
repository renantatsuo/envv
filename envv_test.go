package envv

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	// Test basic Get functionality
	env := Get("TEST_VAR")
	if env.name != "TEST_VAR" {
		t.Errorf("Expected name to be TEST_VAR, got %s", env.name)
	}
}

func TestString(t *testing.T) {
	// Test string parsing
	tests := []struct {
		name         string
		envValue     string
		defaultValue string
		required     bool
		expected     string
		shouldPanic  bool
	}{
		{
			name:     "string",
			envValue: "hello",
			expected: "hello",
		},
		{
			name:         "default string",
			envValue:     "",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:     "required string",
			envValue: "a",
			required: true,
			expected: "a",
		},
		{
			name:        "required string missing",
			envValue:    "",
			required:    true,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv("TEST_STRING")
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("Panic = %v, shouldPanic = %v", r != nil, tt.shouldPanic)
				}
			}()

			if tt.envValue != "" {
				os.Setenv("TEST_STRING", tt.envValue)
			}

			var result string
			if tt.defaultValue != "" {
				result = Get("TEST_STRING").String().Default(tt.defaultValue).Parse()
			} else if tt.required {
				result = Get("TEST_STRING").String().Required().Parse()
			} else {
				result = Get("TEST_STRING").String().Optional().Parse()
			}

			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue int
		required     bool
		expected     int
		shouldPanic  bool
	}{
		{
			name:     "valid int",
			envValue: "42",
			expected: 42,
		},
		{
			name:         "default int",
			envValue:     "",
			defaultValue: 100,
			expected:     100,
		},
		{
			name:        "invalid int",
			envValue:    "not-an-int",
			shouldPanic: true,
		},
		{
			name:        "required int missing",
			envValue:    "",
			required:    true,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv("TEST_INT")
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("Panic = %v, shouldPanic = %v", r != nil, tt.shouldPanic)
				}
			}()

			if tt.envValue != "" {
				os.Setenv("TEST_INT", tt.envValue)
			}

			var result int
			if tt.defaultValue != 0 {
				result = Get("TEST_INT").Int().Default(tt.defaultValue).Parse()
			} else if tt.required {
				result = Get("TEST_INT").Int().Required().Parse()
			} else {
				result = Get("TEST_INT").Int().Optional().Parse()
			}

			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue bool
		required     bool
		expected     bool
		shouldPanic  bool
	}{
		{
			name:     "true value",
			envValue: "true",
			expected: true,
		},
		{
			name:     "false value",
			envValue: "false",
			expected: false,
		},
		{
			name:         "default bool",
			envValue:     "",
			defaultValue: true,
			expected:     true,
		},
		{
			name:        "invalid bool",
			envValue:    "not-a-bool",
			shouldPanic: true,
		},
		{
			name:        "required bool missing",
			envValue:    "",
			required:    true,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv("TEST_BOOL")
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("Panic = %v, shouldPanic = %v", r != nil, tt.shouldPanic)
				}
			}()

			if tt.envValue != "" {
				os.Setenv("TEST_BOOL", tt.envValue)
			}

			var result bool
			if tt.defaultValue != false {
				result = Get("TEST_BOOL").Bool().Default(tt.defaultValue).Parse()
			} else if tt.required {
				result = Get("TEST_BOOL").Bool().Required().Parse()
			} else {
				result = Get("TEST_BOOL").Bool().Optional().Parse()
			}

			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue float64
		required     bool
		expected     float64
		shouldPanic  bool
	}{
		{
			name:     "valid float",
			envValue: "3.14",
			expected: 3.14,
		},
		{
			name:         "default float",
			envValue:     "",
			defaultValue: 2.718,
			expected:     2.718,
		},
		{
			name:        "invalid float",
			envValue:    "not-a-float",
			shouldPanic: true,
		},
		{
			name:        "required float missing",
			envValue:    "",
			required:    true,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv("TEST_FLOAT")
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("Panic = %v, shouldPanic = %v", r != nil, tt.shouldPanic)
				}
			}()

			if tt.envValue != "" {
				os.Setenv("TEST_FLOAT", tt.envValue)
			}

			var result float64
			if tt.defaultValue != 0 {
				result = Get("TEST_FLOAT").Float64().Default(tt.defaultValue).Parse()
			} else if tt.required {
				result = Get("TEST_FLOAT").Float64().Required().Parse()
			} else {
				result = Get("TEST_FLOAT").Float64().Optional().Parse()
			}

			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
