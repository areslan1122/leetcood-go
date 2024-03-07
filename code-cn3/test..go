package main


func dataEditor(data map[int64]string) map[]string {

	tmpData := map[string]int64{}
	for key, value := range data {
		if key > tmpData[value] {
			tmpData[value] = key
		}
	}

}