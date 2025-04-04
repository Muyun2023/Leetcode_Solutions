func (codec *Codec) Encode(strs []string) string {
	defer codec.b.Reset()

	for _, word := range strs {
		codec.b.WriteString(strconv.Itoa(len(word)))
		codec.b.WriteRune('|')
		codec.b.WriteString(word)
	}

	return codec.b.String()
}

func (codec *Codec) Decode(strs string) []string {
	var words []string

	for i := 0; i < len(strs);  {
		lenStart := i
		lenEnd := i

		for strs[lenEnd] != '|' {
			lenEnd++
		}

		var l int
		dec := 1

		for j := lenEnd - 1; j >= lenStart; j-- {
			l += (int(strs[j]) - 48) * dec
			dec *= 10
		}

		start := lenEnd + 1
		end := start + l

		words = append(words, string(strs[start:end]))

		i = end
	}

	return words
}
