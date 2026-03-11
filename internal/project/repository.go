package project

import "github.com/jmoiron/sqlx"

type Project struct {
	db *sqlx.DB
}

func NewProject(db *sqlx.DB) *Project {
	return &Project{
		db: db,
	}
}

func (p *Project) ProjectFindByID(id string) (*ModelProject, error) {
	var project ModelProject
	query := `SELECT * FROM project WHERE id = $1`
	if err := p.db.Get(&project, query, id); err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *Project) ProjectFindByName(name string) (*ModelProject, error) {
	var project ModelProject
	query := `SELECT * FROM project WHERE name = $1`
	if err := p.db.Get(&project, query, name); err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *Project) ProjectGetAll() ([]ModelProject, error) {
	var projects []ModelProject
	query := `SELECT * FROM project OFFSET 0 LIMIT 50`
	if err := p.db.Select(&projects, query); err != nil {
		return nil, err
	}
	return projects, nil
}

func (p *Project) ProjectCreate(project ModelCreateProject) (*ModelCreateProject, error) {
	var projects ModelCreateProject
	query := `INSERT INTO project (description, price, skill , 	user_id) VALUES ($1, $2, $3, $4)
	RETURNING id, description, price, skill, user_id`

	err := p.db.QueryRow(query, project.Description, project.Price,
		project.Skill, project.UserID).Scan(&projects.ID, &projects.Description, &projects.Price,
		&projects.Skill, &projects.UserID)

	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (p *Project) ProjectUpdate(project ModelUpdateProject) (*ModelUpdateProject, error) {
	var projects ModelUpdateProject
	query := `UPDATE project
		SET description = $1,
		    price = $2,
		    skill = $3
		WHERE id = $4
		RETURNING id, description, price, skill`

	err := p.db.QueryRow(query, project.Description, project.Price, project.Skill, project.ID).Scan(
		&projects.ID, &projects.Description, &projects.Price, &projects.Skill)
	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (p *Project) ProjectDelete(id string) error {
	query := `DELETE FROM project WHERE id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
