package common

import "wazetiti/internal/structures"

func RemoveDuplicatesAlerts(Slice []structures.CsvData) []structures.CsvData {
	allKeys := make(map[string]bool)
	rows := []structures.CsvData{}
	for _, item := range Slice {
		if _, value := allKeys[item.UUID]; !value {
			allKeys[item.UUID] = true
			rows = append(rows, item)
		}
	}
	return rows
}