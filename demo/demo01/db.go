package hello

type db interface {
	Get() (string, error)
}

type Demo struct {
	Db db
}

type MyDB struct {
}

func (db *MyDB) Get() (string, error) {
	return "Called real db", nil
}
