package vm

func newMem() *memory {
	return &memory{
		state: map[Address]*AccountState{},
	}
}

type memory struct {
	state map[Address]*AccountState
}

func (s *memory) Get(addr Address) *AccountState {
	return s.state[addr]
}

func (s *memory) Put(addr Address, account *AccountState) {
	s.state[addr] = account
}
