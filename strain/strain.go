package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(predicate func(int) bool) (keeped Ints) {
	if len(ints) == 0 {
		return ints
	}
	keeped = make(Ints, 0)

	for _, i := range ints {
		if predicate(i) {
			keeped = append(keeped, i)
		}
	}

	return
}

func (ints Ints) Discard(predicate func(int) bool) (keeped Ints) {
	if len(ints) == 0 {
		return ints
	}
	keeped = make(Ints, 0)

	for _, i := range ints {
		if !predicate(i) {
			keeped = append(keeped, i)
		}
	}

	return
}

func (lists Lists) Keep(predicate func([]int) bool) (keeped Lists) {
	if len(lists) == 0 {
		return lists
	}

	keeped = make(Lists, 0)

	for _, i := range lists {
		if predicate(i) {
			keeped = append(keeped, i)
		}
	}

	return
}

func (strings Strings) Keep(predicate func(string) bool) (keeped Strings) {
	if len(strings) == 0 {
		return strings
	}

	keeped = make(Strings, 0)

	for _, i := range strings {
		if predicate(i) {
			keeped = append(keeped, i)
		}
	}

	return
}
