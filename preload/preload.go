package preload

import "fmt"

func Cast(preload string) (map[string]string, error) {

	if preload == "" {
		return make(map[string]string), nil
	}
	if len(preload) < 2 {
		return nil, fmt.Errorf("invalid preload format %s", preload)
	}
	if !(string(preload[0]) == "{" && string(preload[len(preload)-1]) == "}") {
		return nil, fmt.Errorf("invalid preload format %s", preload)
	}

	var scope int
	var values, key string
	m := make(map[string]string)

	for _, char := range preload[1 : len(preload)-1] {

		if string(char) == " " || string(char) == "." || string(char) == "," {
			continue
		}

		if string(char) == "{" {
			scope++

			if scope == 1 {
				key = values
				values = ""
			}
		}

		if string(char) == "}" {
			scope--

			if scope == 0 {
				m[key] = values + string(char)
				values = ""
				continue
			}
		}

		values += string(char)
	}

	if scope != 0 {
		return nil, fmt.Errorf("invalid preload format %s", preload)
	}

	return m, nil
}
