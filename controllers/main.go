package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"monolit/models"
	"monolit/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Home() {
	c.TplName = "index.html"
}

func (c *MainController) Add() {
	c.TplName = "add.html"
}

func (c *MainController) Edit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.Data["id"] = id
	if !models.IsDataPegawaiExist(id) {
		c.Data["status"] = "Data Pegawai not found"
		c.TplName = "404.html"
	} else {
		data, _ := models.ReadDataPegawaiDetail(id)

		c.Data["nama_lengkap"] = data.NamaLengkap
		c.Data["nip"] = data.Nip
		c.Data["tempat_lahir"] = data.TempatLahir
		c.Data["tanggal_lahir"] = data.TanggalLahir
		c.Data["tanggal_bergabung"] = data.TanggalBergabung
		c.Data["status"] = data.Status
		c.Data["id_atasan"] = data.Atasan
		c.Data["catatan"] = data.Catatan

		fmt.Println(data.TanggalBergabung)

		c.TplName = "edit.html"
	}

}

func (c *MainController) ReadDataPegawai() {
	var resp utils.ResponseSchema
	resp.Status.Code = 400
	resp.Data = nil

	id, _ := c.GetInt("id")

	if id == 0 {
		if dataPegawai, err := models.ReadDataPegawai(); err != nil {
			resp.Status.Message = err.Error()
		} else {
			resp.Status.Code = 200
			resp.Status.Message = "success"
			resp.Data = dataPegawai
		}
	} else {
		if !models.IsDataPegawaiExist(id) {
			resp.Status.Message = "Data Pegawai not found"
		} else {
			data, _ := models.ReadDataPegawaiDetail(id)
			resp.Status.Code = 200
			resp.Status.Message = "success"
			resp.Data = data
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "GET, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Status.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *MainController) ReadDataAtasan() {
	var resp utils.ResponseSchema
	resp.Status.Code = 400
	resp.Data = nil

	if dataPegawai, err := models.ReadDataAtasan(); err != nil {
		resp.Status.Message = "Data Pegawai not found"
	} else {
		resp.Status.Code = 200
		resp.Status.Message = "success"
		resp.Data = dataPegawai
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "GET, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Status.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

var t string = time.Now().Format("2006-01-02 15:04:05")

func (c *MainController) CreateDataPegawai() {
	var resp utils.ResponseSchema
	resp.Status.Code = 400
	resp.Data = nil

	NamaLengkap := c.GetString("nama_lengkap")
	Nip := c.GetString("nip")
	TempatLahir := c.GetString("tempat_lahir")
	TanggalLahir := c.GetString("tanggal_lahir")
	TanggalBergabung := c.GetString("tanggal_bergabung")
	Status := c.GetString("status")
	Atasan := c.GetString("id_atasan")
	Catatan := c.GetString("catatan")

	valid := validation.Validation{}
	valid.Required(NamaLengkap, "Nama Lengkap")
	valid.Required(Nip, "Nip")
	valid.Required(TempatLahir, "Tempat Lahir")
	valid.Required(TanggalLahir, "Tanggal Lahir")
	valid.Required(TanggalBergabung, "Tanggal Bergabung")
	valid.Required(Status, "Status")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Status.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		var dataPegawai models.DataPegawai
		dataPegawai.NamaLengkap = NamaLengkap
		dataPegawai.Nip = Nip
		dataPegawai.TempatLahir = TempatLahir
		dataPegawai.TanggalLahir = TanggalLahir
		dataPegawai.TanggalBergabung = TanggalBergabung
		dataPegawai.Status = Status
		dataPegawai.Atasan = Atasan
		dataPegawai.Catatan = Catatan

		if _, err := models.CreateDataPegawai(dataPegawai); err != nil {
			resp.Status.Message = err.Error()
		} else {
			resp.Status.Code = 200
			resp.Status.Message = "success"
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Status.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *MainController) UpdateDataPegawai() {
	var resp utils.ResponseSchema
	resp.Status.Code = 400
	resp.Data = nil

	id, _ := c.GetInt("id")
	NamaLengkap := c.GetString("nama_lengkap")
	Nip := c.GetString("nip")
	TempatLahir := c.GetString("tempat_lahir")
	TanggalLahir := c.GetString("tanggal_lahir")
	TanggalBergabung := c.GetString("tanggal_bergabung")
	Status := c.GetString("status")
	Atasan := c.GetString("id_atasan")
	Catatan := c.GetString("catatan")

	valid := validation.Validation{}
	valid.Required(NamaLengkap, "Nama Lengkap")
	valid.Required(id, "ID")
	valid.Required(Nip, "Nip")
	valid.Required(TempatLahir, "Tempat Lahir")
	valid.Required(TanggalLahir, "Tanggal Lahir")
	valid.Required(TanggalBergabung, "Tanggal Bergabung")
	valid.Required(Status, "Status")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Status.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		if !models.IsDataPegawaiExist(id) {
			resp.Status.Message = "data not found"
		} else {
			dataKaryawan := orm.Params{
				"id":                id,
				"nama_lengkap":      NamaLengkap,
				"tempat_lahir":      TempatLahir,
				"tanggal_lahir":     TanggalLahir,
				"Status":            Status,
				"Atasan":            Atasan,
				"tanggal_bergabung": TanggalBergabung,
				"catatan":           Catatan,
			}

			if err := models.UpdateDataPegawai(dataKaryawan); err != nil {
				resp.Status.Message = err.Error()
			} else {
				resp.Status.Code = 200
				resp.Status.Message = "success"
			}
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "PUT")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Status.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *MainController) DeleteDataPegawai() {
	var resp utils.ResponseSchema
	resp.Status.Code = 400
	resp.Data = nil

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	valid := validation.Validation{}
	valid.Required(id, "id")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Status.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		if !models.IsDataPegawaiExist(id) {
			resp.Status.Message = "Data not found"
		} else {
			if err := models.DeleteDataPegawai(id); err != nil {
				resp.Status.Message = err.Error()
			} else {
				resp.Status.Code = 200
				resp.Status.Message = "success"
			}
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "DELETE")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Status.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}
