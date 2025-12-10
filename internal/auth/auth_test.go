package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKeyWithComplexKey tests with a complex API key containing special characters
func TestGetAPIKeyWithComplexKey(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey abc123!@#$%^&*()_+-=[]{}|;:,.<>?/`~"},
	}
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "abc123!@#$%^&*()_+-=[]{}|;:,.<>?/`~"
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

// TestGetAPIKeyWithMultipleSpaces tests with multiple spaces between ApiKey and the actual key
func TestGetAPIKeyWithMultipleSpaces(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey    my-api-key"},
	}
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// 当有多个空格时，strings.Split会返回空字符串作为第二个元素
	// 所以实际返回值应该是空字符串
	expected := ""
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
