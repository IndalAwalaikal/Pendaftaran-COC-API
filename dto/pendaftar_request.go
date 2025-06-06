package dto

type PendaftarRequest struct {
	NamaLengkap string `json:"nama-lengkap"`
	Email       string `json:"email"`
	NoTelp      string `json:"no-telp"`
	BuktiFollow string `json:"bukti-follow"`
	AsalSekolah string `json:"asal-sekolah"`
	PunyaLaptop string `json:"punya-laptop"`
}
