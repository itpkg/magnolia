package site

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/itpkg/magnolia/models"
)

//Set set key => val
func Set(k string, v interface{}, f bool) error {
	o := orm.NewOrm()
	m := Setting{Key: k}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	if f {
		m.Val, err = models.TextEncryptor.Encode(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		m.Val = buf.Bytes()
	}
	m.Flag = f
	m.UpdatedAt = time.Now()

	err = o.Raw("select `key`, val from settings where `key` = ?", "Key").QueryRow(&m)
	if err == orm.ErrNoRows {
		_, err = o.Raw("insert into settings(`key`, val, flag, updated_at) values(? , ?, ?)", m.)
	} else {
		_, err = o.Update(&m, "Val", "UpdatedAt", "Flag")
	}
	return err
}

//Get get val by key
func Get(k string, v interface{}) error {
	o := orm.NewOrm()

	m := Setting{Key: k}
	err := o.Read(&m, "Key")
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = models.TextEncryptor.Decode(m.Val); err != nil {
			return err
		}
	}

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(m.Val)
	return dec.Decode(v)
}
