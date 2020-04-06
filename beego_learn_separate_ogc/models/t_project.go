package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TProject struct {
	Id               int       `orm:"column(project_id);pk"`
	MatchId          *TMatch   `orm:"column(match_id);rel(fk)"`
	Type             int       `orm:"column(type);null"`
	Name             string    `orm:"column(name);null"`
	Price            int       `orm:"column(price);null"`
	NumberLimit      int       `orm:"column(number_limit);null"`
	TimeStart        time.Time `orm:"column(time_start);type(timestamp without time zone);null"`
	TimeEnd          time.Time `orm:"column(time_end);type(timestamp without time zone);null"`
	Poster           string    `orm:"column(poster);null"`
	EarlyPrice       int       `orm:"column(early_price);null"`
	EarlyDeadline    time.Time `orm:"column(early_deadline);type(timestamp without time zone);null"`
	TeamNumber       int       `orm:"column(team_number);null"`
	RepeatTeamNumber int       `orm:"column(repeat_team_number);null"`
	TeamNumberLow    int       `orm:"column(team_number_low);null"`
	TeamNumberHigh   int       `orm:"column(team_number_high);null"`
	Addi             string    `orm:"column(addi);null"`
	Creator          int       `orm:"column(creator);null"`
	CreateTime       time.Time `orm:"column(create_time);type(timestamp without time zone);null"`
	DomainId         int       `orm:"column(domain_id);null"`
	Remark           string    `orm:"column(remark);null"`
}

func (t *TProject) TableName() string {
	return "t_project"
}

func init() {
	orm.RegisterModel(new(TProject))
}

// AddTProject insert a new TProject into database and returns
// last inserted Id on success.
func AddTProject(m *TProject) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTProjectById retrieves TProject by Id. Returns error if
// Id doesn't exist
func GetTProjectById(id int) (v *TProject, err error) {
	o := orm.NewOrm()
	v = &TProject{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTProject retrieves all TProject matches certain condition. Returns empty list if
// no records exist
func GetAllTProject(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TProject))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TProject
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTProject updates TProject by Id and returns error if
// the record to be updated doesn't exist
func UpdateTProjectById(m *TProject) (err error) {
	o := orm.NewOrm()
	v := TProject{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTProject deletes TProject by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTProject(id int) (err error) {
	o := orm.NewOrm()
	v := TProject{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TProject{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
