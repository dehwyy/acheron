package extractors

import "io"

func BodyToString(body io.ReadCloser) (string, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
