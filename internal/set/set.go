package set

type MySet struct {
    data map[interface{}]bool
}


func NewMySet() *MySet {
    return &MySet{data: make(map[interface{}]bool)}
}

func (s *MySet) Add(value interface{}) {
    s.data[value] = true
}

func (s *MySet) Remove(value interface{}) {
    delete(s.data, value)
}

func (s *MySet) Contains(value interface{}) bool {
    _, exists := s.data[value]
    return exists
}

func (s *MySet) Size() int {
    return len(s.data)
}