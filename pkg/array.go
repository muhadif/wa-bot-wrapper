package pkg

func GetContentIfExist(str []string, index int) string {
	if str == nil {
		return ""
	}

	if len(str) > index {
		return str[index]
	}
	return ""
}
