package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TOrganization struct {
	Id         int       `orm:"column(organization_id);pk"`
	Type       int       `orm:"column(type)"`
	Name       string    `orm:"column(name);null"`
	Password   string    `orm:"column(password)"`
	Phone      string    `orm:"column(phone)"`
	Introduce  string    `orm:"column(introduce);null"`
	Contact    string    `orm:"column(contact);null"`
	About      string    `orm:"column(about);null"`
	Logo       string    `orm:"column(logo);null"`
	QrCode     string    `orm:"column(qr_code);null"`
	News       string    `orm:"column(news);null"`
	Authority  string    `orm:"column(authority)"`
	Token      string    `orm:"column(token);null"`
	Addi       string    `orm:"column(addi);null"`
	Creator    int       `orm:"column(creator);null"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp without time zone);null"`
	Remark     string    `orm:"column(remark);null"`
	DomainId   int       `orm:"column(domain_id);null"`
}

func (t *TOrganization) TableName() string {
	return "t_organization"
}

func init() {
	orm.RegisterModel(new(TOrganization))
}

// AddTOrganization insert a new TOrganization into database and returns
// last inserted Id on success.
func AddTOrganization(m *TOrganization) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTOrganizationById retrieves TOrganization by Id. Returns error if
// Id doesn't exist
func GetTOrganizationById(id int) (v *TOrganization, err error) {
	o := orm.NewOrm()
	v = &TOrganization{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTOrganization retrieves all TOrganization matches certain condition. Returns empty list if
// no records exist
func GetAllTOrganization(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TOrganization))
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

	var l []TOrganization
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

// UpdateTOrganization updates TOrganization by Id and returns error if
// the record to be updated doesn't exist
func UpdateTOrganizationById(m *TOrganization) (err error) {
	o := orm.NewOrm()
	v := TOrganization{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTOrganization deletes TOrganization by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTOrganization(id int) (err error) {
	o := orm.NewOrm()
	v := TOrganization{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TOrganization{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
