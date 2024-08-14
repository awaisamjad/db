package utils

func gap(width uint32) string {
	gap := ""
	for i := 0; i < int(width); i++ {
		gap += " "
	}
	return gap
}