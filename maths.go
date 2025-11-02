package main

func XYCoordMap(x, y float64, level *Level) (float64, float64) {
	return XCoordMap(x, level), YCoordMap(y, level)
}

func XCoordMap(value float64, level *Level) float64 {
	return coordMap(value, 0, level.Width, 0, screenWidth)
}

func YCoordMap(value float64, level *Level) float64 {
	return coordMap(value, 0, level.Height, 50, 500)
}

func coordMap(value, levelStart, levelEnd, screenStart, screenEnd float64) float64 {
	//	v1 := ((value-start1)/(stop1-start1))*(stop2-start2) + start2
	v1 := coordMapResize(value, levelStart, levelEnd, screenStart, screenEnd)
	v1 += screenStart
	return v1
}

func coordMapResize(value, levelStart, levelEnd, screenStart, screenEnd float64) float64 {
	//	v1 := ((value-start1)/(stop1-start1))*(stop2-start2) + start2
	v1 := value / (levelEnd - levelStart)
	v2 := v1 * (screenEnd - screenStart)
	return v2
}
