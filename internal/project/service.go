package project

type ProjectHandle struct {
	store *Project
}

func NewProjectHandle(store *Project) *ProjectHandle {
	return &ProjectHandle{
		store: store,
	}
}
