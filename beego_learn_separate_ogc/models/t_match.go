package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TMatch struct {
	Id             int            `orm:"column(match_id);pk"`
	OrganizationId *TOrganization `orm:"column(organization_id);rel(fk)"`
	Name           string         `orm:"column(name);null"`
	Type           int            `orm:"column(type);null"`
	TimeStart      time.Time      `orm:"column(time_start);type(timestamp without time zone);null"`
	TimeEnd        time.Time      `orm:"column(time_end);type(timestamp without time zone);null"`
	Deadline       time.Time      `orm:"column(deadline);type(timestamp without time zone);null"`
	Address        string         `orm:"column(address);null"`
	Contact        string         `orm:"column(contact);null"`
	QrCode         string         `orm:"column(qr_code);null"`
	Note           string         `orm:"column(note);null"`
	Poster         string         `orm:"column(poster);null"`
	State          int            `orm:"column(state)"`
	Addi           string         `orm:"column(addi);null"`
	Creator        int            `orm:"column(creator);null"`
	CreateTime     time.Time      `orm:"column(create_time);type(timestamp without time zone);null"`
	DomainId       int            `orm:"column(domain_id);null"`
	Remark         string         `orm:"column(remark);null"`
}

func (t *TMatch) TableName() string {
	return "t_match"
}

func init() {
	orm.RegisterModel(new(TMatch))
}

// AddTMatch insert a new TMatch into database and returns
// last inserted Id on success.
func AddTMatch(m *TMatch) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTMatchById retrieves TMatch by Id. Returns error if
// Id doesn't exist
func GetTMatchById(id int) (v *TMatch, err error) {
	o := orm.NewOrm()
	v = &TMatch{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTMatch retrieves all TMatch matches certain condition. Returns empty list if
// no records exist
func GetAllTMatch(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TMatch))
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

	var l []TMatch
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

// UpdateTMatch updates TMatch by Id and returns error if
// the record to be updated doesn't exist
func UpdateTMatchById(m *TMatch) (err error) {
	o := orm.NewOrm()
	v := TMatch{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTMatch deletes TMatch by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTMatch(id int) (err error) {
	o := orm.NewOrm()
	v := TMatch{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TMatch{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
