package dto

type PendaftarResponse struct {
	Id          string `json:"id"`
	NamaLengkap string `json:"nama-lengkap"`
	Email       string `json:"email"`
	NoTelp      string `json:"no-telp"`
	BuktiFollow string `json:"bukti-follow"`
}
