// copy from https://go101.org/article/memory-leaking.html

package memory_leak

func WrongSliceElement() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	// do something with s ...

	return s[1:3:3]
}

func GoodSliceElement() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	s[0], s[len(s)-1] = nil, nil
	return s[1:3:3]
}
