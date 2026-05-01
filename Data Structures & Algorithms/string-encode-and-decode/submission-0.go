type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _,v := range strs {
		sb.WriteString(strconv.Itoa(len(v)))
		sb.WriteByte('#')
		sb.WriteString(v)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}

	i := 0
	for i < len(encoded) {
		j := i

		for encoded[j] != '#'{
			j++
		}
		
		length,_ := strconv.Atoi(encoded[i:j])
		j++

		word := encoded[j:j+length]
		res = append(res, word)
		i = j + length
	}
	return res;
}
