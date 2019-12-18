package store

// BaseStore contains the required methods of
// any created session store
type BaseStore interface {
	Destroy(sid string) error                         // delete a session from the store
	Get(sid string) (interface{}, error)              // get a session from the store
	Set(sid string, val map[string]interface{}) error // set values for a session in the store
}

// ExtendedStore contains additional methods
// that a session store can have to extend
// functionality as needed
type ExtendedStore interface {
	BaseStore
	All() ([]string, error) // get all sessions in the store
	Clear() error           // clear/remove all sessions in the store
	Length() (int, error)   // how many sessions in the store
	Touch(sid string) error // what to do when a session is "touched", used in stores where a session is auto deleted
}
