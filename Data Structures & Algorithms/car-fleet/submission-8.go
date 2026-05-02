func countFleet(time []float64) int {
    fleet := 0
    maxTime := 0.0

    for _, t := range time {
        if t > maxTime {
            fleet++
            maxTime = t
        }
    }

    return fleet
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	indices := make([]int,n)
	for i := range indices{
		indices[i]=i
	}
	sort.Slice(indices, func(i,j int)bool{
		return position[indices[i]] > position[indices[j]]
	})

    times := make([]float64, n)
    for i, idx := range indices {
        times[i] = float64(target-position[idx]) / float64(speed[idx])
    }

	return countFleet(times)
}
