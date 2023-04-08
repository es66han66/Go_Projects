package repo

type DB interface {
	Get(id string) (string, error)
}
type Repository struct {
	db DB
}

func NewRepository(db DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) Get(id string) (string, error) {
	return r.db.Get(id)
}
