package db

type Database struct {
	Host string
	Port int
}

func (r *Database) Get(id string) (string, error) {
	// Implements DB interface
	return "123", nil
}
