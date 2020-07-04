package internal

func removeFromSlice(removeString string, list []string) []string {
	indexItem := 0

	for index, item := range list {
		if item == removeString {
			indexItem = index
		}
	}

	return append(list[:indexItem], list[indexItem+1:]...)
}
