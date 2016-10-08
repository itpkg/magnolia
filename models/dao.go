package models

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/itpkg/magnolia/models/site"
)

//T translate
func T(locale, code string, args ...interface{}) string {

	o := orm.NewOrm()
	m := site.Locale{}
	err := o.Raw("SELECT message FROM locales WHERE code = ? AND lang = ? LIMIT 1", code, locale).QueryRow(&m)
	if err == nil {
		return m.Message
	}
	if err != orm.ErrNoRows {
		beego.Error(err)
	}
	return i18n.Tr(locale, code, args...)
}

//Set set key => val
func Set(k string, v interface{}, f bool) error {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	var val []byte
	if f {
		val, err = Encrypt(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		val = buf.Bytes()
	}

	o := orm.NewOrm()
	m := site.Setting{}
	updated := time.Now()
	err = o.Raw("SELECT id FROM settings WHERE key = ? LIMIT 1", k).QueryRow(&m)
	if err == orm.ErrNoRows {
		_, err = o.Raw("INSERT INTO settings(key, val, flag, updated_at) VALUES(? , ?, ?)", k, val, f, updated).Exec()
	} else {
		_, err = o.Raw("UPDATE settings SET val = ?, flag = ? WHERE id = ?", val, f, m.ID).Exec()
	}
	return err
}

//Get get val by key
func Get(k string, v interface{}) error {
	o := orm.NewOrm()

	m := site.Setting{}
	err := o.Raw("SELECT key, val FROM settings WHERE key = ? LIMIT 1", k).QueryRow(&m)
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = Decrypt(m.Val); err != nil {
			return err
		}
	}

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(m.Val)
	return dec.Decode(v)
}

// ----------------------------------------------------------------------------

func init() {
	orm.Debug = beego.AppConfig.String("runmode") != "prod"
}
