package day05

type stack struct {
	head *node
	size int
}

type node struct {
	next *node
	val  string
}

func (s *stack) pop() string {
	if s.size == 0 {
		return ""
	}
	n := s.head
	s.head = n.next
	s.size--
	return n.val
}

func (s *stack) push(val string) {
	n := &node{
		next: s.head,
		val:  val}
	s.head = n
	s.size++
}

func (s stack) peek() string {
	return s.head.val
}

func (s stack) array() []string {
	h := s.head
	o := []string{}
	for h != nil {
		o = append(o, h.val)
		h = h.next
	}
	return o
}
