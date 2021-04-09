package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type DataPegawai struct {
	ID               string `orm:"column(id);pk"`
	NamaLengkap      string
	Nip              string
	TempatLahir      string
	TanggalLahir     string
	TanggalBergabung string
	Status           string
	Atasan           string
	Catatan          string
}

func init() {
	orm.RegisterModel(new(DataPegawai))
}

func IsDataPegawaiExist(ID int) bool {
	o := orm.NewOrm()

	exist := o.QueryTable("data_pegawai").Filter("id", ID).Exist()

	return exist
}

func ReadDataAtasan() ([]DataPegawai, error) {
	var v []DataPegawai
	o := orm.NewOrm()

	if _, err := o.QueryTable("data_pegawai").Filter("status", 1).All(&v); err != nil {
		return nil, err
	}

	if len(v) == 0 {
		return nil, errors.New("empty data")
	}

	return v, nil
}

func ReadDataPegawai() ([]DataPegawai, error) {
	var v []DataPegawai
	o := orm.NewOrm()

	if _, err := o.QueryTable("data_pegawai").All(&v); err != nil {
		return nil, err
	}

	if len(v) == 0 {
		return nil, errors.New("empty data")
	}

	return v, nil
}

func ReadDataPegawaiDetail(ID int) (DataPegawai, error) {
	o := orm.NewOrm()

	var v DataPegawai
	if err := o.QueryTable("data_pegawai").Filter("id", ID).One(&v); err != nil {
		return DataPegawai{}, err
	}

	return v, nil
}

func CreateDataPegawai(data DataPegawai) (int, error) {
	o := orm.NewOrm()

	id, err := o.Insert(&data)

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func UpdateDataPegawai(data_pegawai orm.Params) error {
	o := orm.NewOrm()

	if _, err := o.QueryTable("data_pegawai").Filter("id", data_pegawai["id"]).Update(data_pegawai); err != nil {
		return err
	}
	return nil
}

func DeleteDataPegawai(dataPegawaiID int) error {
	o := orm.NewOrm()
	if _, err := o.QueryTable("data_pegawai").Filter("id", dataPegawaiID).Delete(); err != nil {
		return err
	}
	return nil
}
