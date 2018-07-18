package page

type Repository struct{}

func (r *Repository) NewRepository() *Repository {
	return &Repository{}
}

// TODO: domain/page.go PageInterfaceを実装する。
