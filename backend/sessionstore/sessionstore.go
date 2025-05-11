// simple session store, store key value in memory

package sessionstore

type SessionStore struct {
	data map[string]interface{}
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		data: make(map[string]interface{}),
	}
}
func (s *SessionStore) Set(key string, value interface{}) {
	s.data[key] = value
}
func (s *SessionStore) Get(key string) (interface{}, bool) {
	value, exists := s.data[key]
	return value, exists
}
func (s *SessionStore) Delete(key string) {
	delete(s.data, key)
}
func (s *SessionStore) Clear() {
	s.data = make(map[string]interface{})
}
func (s *SessionStore) GetAll() map[string]interface{} {
	return s.data
}
