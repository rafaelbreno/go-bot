package helpers

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FindInSliceStr(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func RemoveElementStr(arr []string, item string) []string {
	for pos, v := range arr {
		if v == item {
			return append(arr[0:pos], arr[pos+1:]...)
		}
	}
	return arr
}
