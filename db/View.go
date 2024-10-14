package db

type View struct {
	ID        int    `db:"id" json:"id" backup:"-"`
	ProjectID int    `db:"project_id" json:"project_id" backup:"-"`
	Title     string `db:"title" json:"title"`
	Position  int    `db:"position" json:"position"`
}

func (view *View) Validate() error {
	if view.Title == "" {
		return &ValidationError{"title can not be empty"}
	}
	return nil
}
