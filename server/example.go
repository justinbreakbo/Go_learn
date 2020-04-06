package main

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// TLesson represents public.t_lesson
type TLesson struct {
	LessonID         int            // lesson_id
	OrganizationID   sql.NullInt64  // organization_id
	Name             sql.NullString // name
	Type             sql.NullInt64  // type
	LessonTime       sql.NullString // lesson_time
	Deadline         *time.Time     // deadline
	Address          sql.NullString // address
	IsTransfer       bool           // is_transfer
	Contact          []byte         // contact
	QrCode           sql.NullString // qr_code
	Note             sql.NullString // note
	Poster           sql.NullString // poster
	Introduce        sql.NullString // introduce
	TeacherIntroduce sql.NullString // teacher_introduce
	Prize            sql.NullString // prize
	State            int            // state
	Addi             []byte         // addi
	Creator          sql.NullInt64  // creator
	Remark           sql.NullString // remark
	CreateTime       *time.Time     // create_time
	DomainID         sql.NullInt64  // domain_id
}

// Create inserts the TLesson to the database.
func (r *TLesson) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_lesson (organization_id, name, type, lesson_time, deadline, address, is_transfer, contact, qr_code, note, poster, introduce, teacher_introduce, prize, state, addi, creator, remark, create_time, domain_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING lesson_id`,
		&r.OrganizationID, &r.Name, &r.Type, &r.LessonTime, &r.Deadline, &r.Address, &r.IsTransfer, &r.Contact, &r.QrCode, &r.Note, &r.Poster, &r.Introduce, &r.TeacherIntroduce, &r.Prize, &r.State, &r.Addi, &r.Creator, &r.Remark, &r.CreateTime, &r.DomainID).Scan(&r.LessonID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_lesson")
	}
	return nil
}

// GetTLessonByPk select the TLesson from the database.
func GetTLessonByPk(db Queryer, pk0 int) (*TLesson, error) {
	var r TLesson
	err := db.QueryRow(
		`SELECT lesson_id, organization_id, name, type, lesson_time, deadline, address, is_transfer, contact, qr_code, note, poster, introduce, teacher_introduce, prize, state, addi, creator, remark, create_time, domain_id FROM t_lesson WHERE lesson_id = $1`,
		pk0).Scan(&r.LessonID, &r.OrganizationID, &r.Name, &r.Type, &r.LessonTime, &r.Deadline, &r.Address, &r.IsTransfer, &r.Contact, &r.QrCode, &r.Note, &r.Poster, &r.Introduce, &r.TeacherIntroduce, &r.Prize, &r.State, &r.Addi, &r.Creator, &r.Remark, &r.CreateTime, &r.DomainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_lesson")
	}
	return &r, nil
}

// TMatch represents public.t_match
type TMatch struct {
	MatchID        int            // match_id
	OrganizationID sql.NullInt64  // organization_id
	Name           sql.NullString // name
	Type           sql.NullInt64  // type
	TimeStart      *time.Time     // time_start
	TimeEnd        *time.Time     // time_end
	Deadline       *time.Time     // deadline
	Address        sql.NullString // address
	Contact        []byte         // contact
	QrCode         sql.NullString // qr_code
	Note           sql.NullString // note
	Poster         sql.NullString // poster
	State          int            // state
	Addi           []byte         // addi
	Creator        sql.NullInt64  // creator
	CreateTime     *time.Time     // create_time
	DomainID       sql.NullInt64  // domain_id
	Remark         sql.NullString // remark
}

// Create inserts the TMatch to the database.
func (r *TMatch) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_match (organization_id, name, type, time_start, time_end, deadline, address, contact, qr_code, note, poster, state, addi, creator, create_time, domain_id, remark) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING match_id`,
		&r.OrganizationID, &r.Name, &r.Type, &r.TimeStart, &r.TimeEnd, &r.Deadline, &r.Address, &r.Contact, &r.QrCode, &r.Note, &r.Poster, &r.State, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark).Scan(&r.MatchID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_match")
	}
	return nil
}

// GetTMatchByPk select the TMatch from the database.
func GetTMatchByPk(db Queryer, pk0 int) (*TMatch, error) {
	var r TMatch
	err := db.QueryRow(
		`SELECT match_id, organization_id, name, type, time_start, time_end, deadline, address, contact, qr_code, note, poster, state, addi, creator, create_time, domain_id, remark FROM t_match WHERE match_id = $1`,
		pk0).Scan(&r.MatchID, &r.OrganizationID, &r.Name, &r.Type, &r.TimeStart, &r.TimeEnd, &r.Deadline, &r.Address, &r.Contact, &r.QrCode, &r.Note, &r.Poster, &r.State, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_match")
	}
	return &r, nil
}

// TOptions represents public.t_options
type TOptions struct {
	OptionID    int            // option_id
	LessonID    sql.NullInt64  // lesson_id
	ClassNumber int            // class_number
	Price       int            // price
	Discount    []byte         // discount
	Addi        []byte         // addi
	Creator     sql.NullInt64  // creator
	CreateTime  *time.Time     // create_time
	DomainID    sql.NullInt64  // domain_id
	Remark      sql.NullString // remark
	MemberLimit sql.NullInt64  // member_limit
}

// Create inserts the TOptions to the database.
func (r *TOptions) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_options (lesson_id, class_number, price, discount, addi, creator, create_time, domain_id, remark, member_limit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING option_id`,
		&r.LessonID, &r.ClassNumber, &r.Price, &r.Discount, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark, &r.MemberLimit).Scan(&r.OptionID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_options")
	}
	return nil
}

// GetTOptionsByPk select the TOptions from the database.
func GetTOptionsByPk(db Queryer, pk0 int) (*TOptions, error) {
	var r TOptions
	err := db.QueryRow(
		`SELECT option_id, lesson_id, class_number, price, discount, addi, creator, create_time, domain_id, remark, member_limit FROM t_options WHERE option_id = $1`,
		pk0).Scan(&r.OptionID, &r.LessonID, &r.ClassNumber, &r.Price, &r.Discount, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark, &r.MemberLimit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_options")
	}
	return &r, nil
}

// TOrganization represents public.t_organization
type TOrganization struct {
	OrganizationID int            // organization_id
	Type           int            // type
	Name           sql.NullString // name
	Password       string         // password
	Phone          string         // phone
	Introduce      sql.NullString // introduce
	Contact        []byte         // contact
	About          sql.NullString // about
	Logo           sql.NullString // logo
	QrCode         sql.NullString // qr_code
	News           []byte         // news
	Authority      string         // authority
	Token          sql.NullString // token
	Addi           []byte         // addi
	Creator        sql.NullInt64  // creator
	CreateTime     *time.Time     // create_time
	Remark         sql.NullString // remark
	DomainID       sql.NullInt64  // domain_id
}

// Create inserts the TOrganization to the database.
func (r *TOrganization) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_organization (type, name, password, phone, introduce, contact, about, logo, qr_code, news, authority, token, addi, creator, create_time, remark, domain_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING organization_id`,
		&r.Type, &r.Name, &r.Password, &r.Phone, &r.Introduce, &r.Contact, &r.About, &r.Logo, &r.QrCode, &r.News, &r.Authority, &r.Token, &r.Addi, &r.Creator, &r.CreateTime, &r.Remark, &r.DomainID).Scan(&r.OrganizationID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_organization")
	}
	return nil
}

// GetTOrganizationByPk select the TOrganization from the database.
func GetTOrganizationByPk(db Queryer, pk0 int) (*TOrganization, error) {
	var r TOrganization
	err := db.QueryRow(
		`SELECT organization_id, type, name, password, phone, introduce, contact, about, logo, qr_code, news, authority, token, addi, creator, create_time, remark, domain_id FROM t_organization WHERE organization_id = $1`,
		pk0).Scan(&r.OrganizationID, &r.Type, &r.Name, &r.Password, &r.Phone, &r.Introduce, &r.Contact, &r.About, &r.Logo, &r.QrCode, &r.News, &r.Authority, &r.Token, &r.Addi, &r.Creator, &r.CreateTime, &r.Remark, &r.DomainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_organization")
	}
	return &r, nil
}

// TProject represents public.t_project
type TProject struct {
	ProjectID        int            // project_id
	MatchID          sql.NullInt64  // match_id
	Type             sql.NullInt64  // type
	Name             sql.NullString // name
	Price            sql.NullInt64  // price
	NumberLimit      sql.NullInt64  // number_limit
	TimeStart        *time.Time     // time_start
	TimeEnd          *time.Time     // time_end
	Poster           sql.NullString // poster
	EarlyPrice       sql.NullInt64  // early_price
	EarlyDeadline    *time.Time     // early_deadline
	TeamNumber       sql.NullInt64  // team_number
	RepeatTeamNumber sql.NullInt64  // repeat_team_number
	TeamNumberLow    sql.NullInt64  // team_number_low
	TeamNumberHigh   sql.NullInt64  // team_number_high
	Addi             []byte         // addi
	Creator          sql.NullInt64  // creator
	CreateTime       *time.Time     // create_time
	DomainID         sql.NullInt64  // domain_id
	Remark           sql.NullString // remark
}

// Create inserts the TProject to the database.
func (r *TProject) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_project (match_id, type, name, price, number_limit, time_start, time_end, poster, early_price, early_deadline, team_number, repeat_team_number, team_number_low, team_number_high, addi, creator, create_time, domain_id, remark) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) RETURNING project_id`,
		&r.MatchID, &r.Type, &r.Name, &r.Price, &r.NumberLimit, &r.TimeStart, &r.TimeEnd, &r.Poster, &r.EarlyPrice, &r.EarlyDeadline, &r.TeamNumber, &r.RepeatTeamNumber, &r.TeamNumberLow, &r.TeamNumberHigh, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark).Scan(&r.ProjectID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_project")
	}
	return nil
}

// GetTProjectByPk select the TProject from the database.
func GetTProjectByPk(db Queryer, pk0 int) (*TProject, error) {
	var r TProject
	err := db.QueryRow(
		`SELECT project_id, match_id, type, name, price, number_limit, time_start, time_end, poster, early_price, early_deadline, team_number, repeat_team_number, team_number_low, team_number_high, addi, creator, create_time, domain_id, remark FROM t_project WHERE project_id = $1`,
		pk0).Scan(&r.ProjectID, &r.MatchID, &r.Type, &r.Name, &r.Price, &r.NumberLimit, &r.TimeStart, &r.TimeEnd, &r.Poster, &r.EarlyPrice, &r.EarlyDeadline, &r.TeamNumber, &r.RepeatTeamNumber, &r.TeamNumberLow, &r.TeamNumberHigh, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_project")
	}
	return &r, nil
}

// TStuLesson represents public.t_stu_lesson
type TStuLesson struct {
	ID         int            // id
	OptionID   sql.NullInt64  // option_id
	StudentID  sql.NullInt64  // student_id
	PayState   sql.NullInt64  // pay_state
	PayTime    *time.Time     // pay_time
	PayPrice   sql.NullInt64  // pay_price
	Addi       []byte         // addi
	Creator    sql.NullInt64  // creator
	CreateTime *time.Time     // create_time
	DomainID   sql.NullInt64  // domain_id
	Remark     sql.NullString // remark
}

// Create inserts the TStuLesson to the database.
func (r *TStuLesson) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_stu_lesson (option_id, student_id, pay_state, pay_time, pay_price, addi, creator, create_time, domain_id, remark) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		&r.OptionID, &r.StudentID, &r.PayState, &r.PayTime, &r.PayPrice, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_stu_lesson")
	}
	return nil
}

// GetTStuLessonByPk select the TStuLesson from the database.
func GetTStuLessonByPk(db Queryer, pk0 int) (*TStuLesson, error) {
	var r TStuLesson
	err := db.QueryRow(
		`SELECT id, option_id, student_id, pay_state, pay_time, pay_price, addi, creator, create_time, domain_id, remark FROM t_stu_lesson WHERE id = $1`,
		pk0).Scan(&r.ID, &r.OptionID, &r.StudentID, &r.PayState, &r.PayTime, &r.PayPrice, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_stu_lesson")
	}
	return &r, nil
}

// TUserMatch represents public.t_user_match
type TUserMatch struct {
	ID               int            // id
	ProjectID        sql.NullInt64  // project_id
	UserID           sql.NullInt64  // user_id
	PayState         sql.NullInt64  // pay_state
	PayTime          *time.Time     // pay_time
	PayPrice         sql.NullInt64  // pay_price
	UserName         sql.NullString // user_name
	TeamName         sql.NullString // team_name
	TeamActualNumber sql.NullInt64  // team_actual_number
	Addi             []byte         // addi
	Creator          sql.NullInt64  // creator
	CreateTime       *time.Time     // create_time
	DomainID         sql.NullInt64  // domain_id
	Remark           sql.NullString // remark
}

// Create inserts the TUserMatch to the database.
func (r *TUserMatch) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO t_user_match (project_id, user_id, pay_state, pay_time, pay_price, user_name, team_name, team_actual_number, addi, creator, create_time, domain_id, remark) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		&r.ProjectID, &r.UserID, &r.PayState, &r.PayTime, &r.PayPrice, &r.UserName, &r.TeamName, &r.TeamActualNumber, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert t_user_match")
	}
	return nil
}

// GetTUserMatchByPk select the TUserMatch from the database.
func GetTUserMatchByPk(db Queryer, pk0 int) (*TUserMatch, error) {
	var r TUserMatch
	err := db.QueryRow(
		`SELECT id, project_id, user_id, pay_state, pay_time, pay_price, user_name, team_name, team_actual_number, addi, creator, create_time, domain_id, remark FROM t_user_match WHERE id = $1`,
		pk0).Scan(&r.ID, &r.ProjectID, &r.UserID, &r.PayState, &r.PayTime, &r.PayPrice, &r.UserName, &r.TeamName, &r.TeamActualNumber, &r.Addi, &r.Creator, &r.CreateTime, &r.DomainID, &r.Remark)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select t_user_match")
	}
	return &r, nil
}

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}
