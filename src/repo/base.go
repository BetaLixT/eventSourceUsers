package repo

type Base interface {
  Create(interface{}) (Event, interface{})
  Update(int, interface{}) (Event, interface{})
  Delete(int, interface{}) (Event, interface{})
}
