package log

type Store struct {
    sequence int
}

func (s *Store) Next() int {
    s.sequence++
    return s.sequence
}

func (s *Store) Current() int {
    return s.sequence
}

