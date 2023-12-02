package solutions

func Day_six(arr []byte, markerLen int) int {
	symbol := [26]int{}
	anchor, i := 0, 0
	for ; i-anchor != markerLen && anchor < len(arr)-markerLen-1; i++ {
		lastSymbolPos := symbol[arr[i]-97]
		if lastSymbolPos != 0 && lastSymbolPos >= anchor {
			anchor = lastSymbolPos + 1
		}
		symbol[arr[i]-97] = i
	}
	return i
}