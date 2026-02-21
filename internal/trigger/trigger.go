package trigger

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Rule defines a trigger condition for notifications.
type Rule struct {
	Type  string `json:"type"`  // contains, not_contains, regex, not_regex
	Value string `json:"value"` // text or regex pattern
}

// ParseShorthand parses "type:value" shorthand into a JSON rule string.
// e.g. "contains:out of stock" → {"type":"contains","value":"out of stock"}
func ParseShorthand(input string) (string, error) {
	idx := strings.Index(input, ":")
	if idx < 0 {
		return "", fmt.Errorf("invalid trigger rule: expected 'type:value' (e.g. 'contains:some text')")
	}

	typ := input[:idx]
	val := input[idx+1:]

	switch typ {
	case "contains", "not_contains", "regex", "not_regex":
	default:
		return "", fmt.Errorf("unknown trigger type %q (valid: contains, not_contains, regex, not_regex)", typ)
	}

	if val == "" {
		return "", fmt.Errorf("trigger value cannot be empty")
	}

	// Validate regex if applicable
	if typ == "regex" || typ == "not_regex" {
		if _, err := regexp.Compile(val); err != nil {
			return "", fmt.Errorf("invalid regex %q: %w", val, err)
		}
	}

	r := Rule{Type: typ, Value: val}
	b, _ := json.Marshal(r)
	return string(b), nil
}

// Evaluate checks whether the trigger condition is met for the given content.
// Returns true if the notification should fire.
func Evaluate(ruleJSON string, content string) (bool, error) {
	if ruleJSON == "" {
		return true, nil
	}

	var r Rule
	if err := json.Unmarshal([]byte(ruleJSON), &r); err != nil {
		return true, fmt.Errorf("invalid trigger rule JSON: %w", err)
	}

	switch r.Type {
	case "contains":
		return strings.Contains(content, r.Value), nil
	case "not_contains":
		return !strings.Contains(content, r.Value), nil
	case "regex":
		re, err := regexp.Compile(r.Value)
		if err != nil {
			return true, fmt.Errorf("invalid regex: %w", err)
		}
		return re.MatchString(content), nil
	case "not_regex":
		re, err := regexp.Compile(r.Value)
		if err != nil {
			return true, fmt.Errorf("invalid regex: %w", err)
		}
		return !re.MatchString(content), nil
	default:
		return true, fmt.Errorf("unknown trigger type: %s", r.Type)
	}
}

// Describe returns a human-readable description of the trigger rule.
func Describe(ruleJSON string) string {
	if ruleJSON == "" {
		return ""
	}
	var r Rule
	if err := json.Unmarshal([]byte(ruleJSON), &r); err != nil {
		return ruleJSON
	}
	switch r.Type {
	case "contains":
		return fmt.Sprintf("trigger if contains %q", r.Value)
	case "not_contains":
		return fmt.Sprintf("trigger if missing %q", r.Value)
	case "regex":
		return fmt.Sprintf("trigger if matches /%s/", r.Value)
	case "not_regex":
		return fmt.Sprintf("trigger if not matches /%s/", r.Value)
	default:
		return ruleJSON
	}
}
