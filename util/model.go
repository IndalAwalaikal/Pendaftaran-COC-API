package util

import (
	"github.com/syrlramadhan/pendaftaran-coc/dto"
	"github.com/syrlramadhan/pendaftaran-coc/model"
)

func ToPendaftarResponse(pendaftar model.Pendaftar) dto.PendaftarResponse {
	return dto.PendaftarResponse{
		Id:          pendaftar.Id,
		NamaLengkap: pendaftar.NamaLengkap,
		Email:       pendaftar.Email,
		NoTelp:      pendaftar.NoTelp,
		BuktiFollow: pendaftar.BuktiFollow,
		AsalSekolah: pendaftar.AsalSekolah,
		PunyaLaptop: pendaftar.PunyaLaptop,
	}
}

func ToPendaftarListResponse(pendaftar []model.Pendaftar) []dto.PendaftarResponse {
	var pendaftarResponse []dto.PendaftarResponse
	for _, pendaftars := range pendaftar {
		pendaftarResponse = append(pendaftarResponse, ToPendaftarResponse(pendaftars))
	}

	return pendaftarResponse
}
