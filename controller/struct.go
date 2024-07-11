package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example: "123456789"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty" example: "Devi Wulandari"`
	Npm          int				 `bson:"npm,omitempty" json:"npm,omitempty" example: "714220050"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example: "08123456789"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty" example: "Informatika"`
	Alamat		 string				 `bson:"alamat,omitempty" json:"alamat,omitempty" example: "Sarijadi"`
	Email		 string				 `bson:"email,omitempty" json:"email,omitempty" example: "devi443@gmail.com"`
}

type Matkul struct {
	Nama_matkul   string   `bson:"namamatkul,omitempty" json:"namamatkul,omitempty" example: "Kewirausahaan"`
	Jadwal        string   `bson:"jadwal,omitempty" json:"jadwal,omitempty" example: "Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu"`
	Sks      	  int      `bson:"sks,omitempty" json:"sks,omitempty" example: "2"`
	Dosen         string   `bson:"dosen,omitempty" json:"dosen,omitempty" example: "Roni Habibie"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example: "123456789"`
	Npm			 int                `bson:"npm,omitempty" json:"npm,omitempty" example:"714220050"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"swaggertype:"string" example: "2024-09-01T00:00:00Z" format:"date-time"`
	Matkul       Matkul             `bson:"matkul,omitempty" json:"matkul,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Checkin		 string				`bson:"checkin,omitempty" json:"checkin,omitempty" example: "Hadir"`
}

type ReqMahasiswa struct {
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty" example: "Devi Wulandari"`
	Npm          int				 `bson:"npm,omitempty" json:"npm,omitempty" example: "714220050"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example: "08123456789"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty" example: "Informatika"`
	Alamat		 string				 `bson:"alamat,omitempty" json:"alamat,omitempty" example: "Sarijadi"`
	Email		 string				 `bson:"email,omitempty" json:"email,omitempty" example: "devi443@gmail.com"`
}

type ReqPresensi struct {
	Npm			 int                `bson:"npm,omitempty" json:"npm,omitempty" example:"714220050"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"swaggertype:"string" example: "2024-09-01T00:00:00Z" format:"date-time"`
	Matkul       Matkul             `bson:"matkul,omitempty" json:"matkul,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Checkin		 string				`bson:"checkin,omitempty" json:"checkin,omitempty" example: "Hadir"`
}