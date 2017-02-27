package musixmatch

import (
	"bytes"
	"time"
)

// FlexBool is a custom bool type that can convert 1 and 0 to true and false
type FlexBool bool

// UnmarshalJSON implements the Unmarshaler interface for FlexBool
func (fx *FlexBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("1")) {
		*fx = true
	} else {
		*fx = false
	}
	return nil
}

// MarshalJSON implements the Marshaler interface for FlexBool
func (fx FlexBool) MarshalJSON() ([]byte, error) {
	if fx {
		return []byte("1"), nil
	}

	return []byte("0"), nil
}

// FlexTime is a custom type that holds musixmatch time representation, which can be either time or empty string
type FlexTime struct {
	time.Time
}

// UnmarshalJSON implements the Unmarshaler interface for FlexTime
func (ft *FlexTime) UnmarshalJSON(data []byte) error {
	ft.Time.UnmarshalJSON(data)
	if string(data) == "\"\"" {
		ft.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	if err != nil {
		return err
	}
	ft.Time = t
	return nil
}

func String(v string) *string {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Int(v int) *int {
	return &v
}
