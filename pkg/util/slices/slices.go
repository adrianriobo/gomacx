package slices

func Contains[X any, Y comparable](source []*X, target Y, compare func(x *X, y Y) bool) bool {
	for _, s := range source {
		if compare(s, target) {
			return true
		}
	}
	return false
}

func GetElementByMatch[X any, Y comparable](source []*X, target Y, match func(x *X, y Y) bool) *X {
	for _, s := range source {
		if match(s, target) {
			return s
		}
	}
	return nil
}
