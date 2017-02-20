package musixmatch

import "bytes"

type FlexBool bool

func (fx *FlexBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("1")) {
		*fx = true
	} else {
		*fx = false
	}
	return nil
}

func (fx FlexBool) MarshalJSON() ([]byte, error) {
	if fx {
		return []byte("1"), nil
	}

	return []byte("0"), nil
}

func String(v string) *string {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}
