package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func rows2Groups(rows pgx.Rows) ([]model.Group, error) {
	var groups []model.Group
	for rows.Next() {
		group := model.Group{}
		err := rows.Scan(&group.Id, &group.GroupNumber, &group.Section)
		if err != nil {
			return nil, fmt.Errorf("convert to group model error: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func rows2Sections(rows pgx.Rows) ([]model.Section, error) {
	var sections []model.Section
	for rows.Next() {
		section := model.Section{}
		err := rows.Scan(&section.Id, &section.Title)
		if err != nil {
			return nil, fmt.Errorf("convert to section model error: %w", err)
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func CreateGroup(pg *db.Postgres, ctx context.Context, group model.Group) (int, error) {
	query := `INSERT INTO groups (group_number, section) VALUES (@number, @section)`
	args := pgx.NamedArgs{
		"number":  group.GroupNumber,
		"section": group.Section,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to insert row in CreateGroup: %w", err)
	}
	query = `select last_value from groups_id_seq`
	rows, err := pg.Db.Query(ctx, query)
	var lastValue int
	rows.Next()
	err = rows.Scan(&lastValue)
	if err != nil {
		return 0, err
	}
	return lastValue, nil
}

func GetGroup(pg *db.Postgres, ctx context.Context, id int) (*model.Group, error) {
	query := `SELECT * FROM groups where id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve attribute in GetGroup: %w", err)
	}

	defer rows.Close()

	var groups []model.Group

	groups, err = rows2Groups(rows)
	if err != nil {
		return nil, err
	}
	if len(groups) == 0 {
		return nil, nil
	}

	return &groups[0], nil
}

func UpdateGroup(pg *db.Postgres, ctx context.Context, group model.Group) error {
	query := `UPDATE groups SET group_number = @number, section = @section WHERE id = @id`
	args := pgx.NamedArgs{
		"id":      group.Id,
		"number":  group.GroupNumber,
		"section": group.Section,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update in UpdateGroup: %w", err)
	}

	return nil
}

func DeleteGroup(pg *db.Postgres, ctx context.Context, id int) error {
	query := `DELETE FROM groups WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove old attribute in DeleteGroup: %w", err)
	}
	return nil
}

func AddGroupMember(pg *db.Postgres, ctx context.Context, person int, group int) error {
	query := `INSERT INTO groups_persons VALUES (@group, @person)`
	args := pgx.NamedArgs{
		"group":  group,
		"person": person,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to add group member: %w", err)
	}
	return nil
}

func RemoveGroupMember(pg *db.Postgres, ctx context.Context, person int, group int) error {
	query := `DELETE FROM groups_persons WHERE person = @person AND group_id = @group`
	args := pgx.NamedArgs{
		"person": person,
		"group":  group,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove group member: %w", err)
	}
	return nil
}

func GetGroupMembers(pg *db.Postgres, ctx context.Context, group int) ([]int, error) {
	query := `SELECT distinct person FROM groups_persons WHERE group_id = @group`
	args := pgx.NamedArgs{
		"group": group,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve group members: %w", err)
	}
	defer rows.Close()

	var members []int
	for rows.Next() {
		var memberId int
		err := rows.Scan(&memberId)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve group members: %w", err)
		}
		members = append(members, memberId)
	}
	return members, nil
}

func GetGroups(pg *db.Postgres, ctx context.Context) ([]model.Group, error) {
	query := `SELECT * from groups`
	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve groups: %w", err)
	}
	defer rows.Close()

	var groups []model.Group

	groups, err = rows2Groups(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve groups: %w", err)
	}
	return groups, nil
}

func CreateSection(pg *db.Postgres, ctx context.Context, section model.Section) (int, error) {
	query := `INSERT INTO sections (title) VALUES (@title)`
	args := pgx.NamedArgs{
		"title": section.Title,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to insert row in CreateSection: %w", err)
	}
	query = `select last_value from sections_id_seq`
	rows, err := pg.Db.Query(ctx, query)
	var lastValue int
	err = rows.Scan(&lastValue)
	if err != nil {
		return 0, err
	}
	return lastValue, nil
}

func GetSection(pg *db.Postgres, ctx context.Context, id int) (*model.Section, error) {
	query := `SELECT * FROM sections WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve section: %w", err)
	}
	defer rows.Close()
	var section model.Section
	err = rows.Scan(&section.Id)
	if err != nil {
		return nil, err
	}
	return &section, nil
}

func UpdateSection(pg *db.Postgres, ctx context.Context, section model.Section) error {
	query := `UPDATE sections SET title = @title WHERE id = @id`
	args := pgx.NamedArgs{
		"id":    section.Id,
		"title": section.Title,
	}
	_, err := pg.Db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to update in UpdateSection: %w", err)
	}
	return nil
}

func DeleteSection(pg *db.Postgres, ctx context.Context, id int) error {
	query := `DELETE FROM sections WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove old section: %w", err)
	}
	return nil
}

func GetAllSections(pg *db.Postgres, ctx context.Context) ([]model.Section, error) {
	query := `SELECT * FROM sections`
	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve all sections: %w", err)
	}
	defer rows.Close()
	var sections []model.Section
	sections, err = rows2Sections(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve all sections: %w", err)
	}
	return sections, nil
}

func GetGroupsFromSections(pg *db.Postgres, ctx context.Context, section int) ([]model.Group, error) {
	query := `SELECT * FROM groups where section = @section`
	args := pgx.NamedArgs{
		"section": section,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve groups: %w", err)
	}
	defer rows.Close()
	var groups []model.Group
	groups, err = rows2Groups(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve groups: %w", err)
	}
	return groups, nil
}
