package goccavenue_test

import (
	"testing"

	"github.com/Unaxiom/goccavenue"
	"github.com/stretchr/testify/require"
)

func TestCCAvenueEncryption(t *testing.T) {
	assert := require.New(t)

	const merchantId = "9999999"
	const encryptionKey = "ABCDEFGHIJ0123456789abcdefghij01"
	goccavenue.SetMerchantId(merchantId)
	goccavenue.SetEncryptionKey(encryptionKey)

	assert.Equal(goccavenue.GetMerchantId(), merchantId)
	assert.Equal(goccavenue.GetEncryptionKey(), encryptionKey)

	req := goccavenue.CCAvenueRequest{
		MerchantId:  goccavenue.GetMerchantId(),
		OrderId:     "1234",
		Currency:    "INR",
		Amount:      "120.00",
		RedirectUrl: "http://localhost:33333/invoices/ccav/payments/redirect",
		CancelUrl:   "http://localhost:33333/invoices/ccav/payments/cancel",
		Language:    "EN",

		BillingName:    "First Last",
		BillingAddress: "No Man Land",
		BillingCity:    "Hyderabad",
		BillingState:   "Telangana",
		BillingZip:     "500001",
		BillingCountry: "India",
		BillingTel:     "9999999999",
		BillingEmail:   "test@domain.com",
	}

	plainText := goccavenue.CreatePayload(req)
	encryptedText, err := goccavenue.EncryptPayload(req)
	assert.Nil(err)
	// -------------------------------------------------------------------------------------------------------------------------
	// Matching it against a known encrypted output
	assert.Equal(encryptedText, `f83543470d23118042a0f4c5197baa647ee3ee81e5483a1ba24ffc6956de70aa0c2827478cccdfa207598cdbdc7d0c1947d9207f879f1ee2e048b9eefade91141194f7796eb447ab6ee973c5bda3c58712a8924056647e90b7decd1865ade8ae862a2b85d1e6fcb5a5f300830e648c49bfc1b4a6c742d8cc1895b6a68461a486f61899d0103a863d9c5c69750c3eb045a788e070f7c19afd5e873e7f437eb233364e8601b6657dcc600510c16090adffdbf1f25c4c29aa18d6d5fb3e80412aea583ffe71c82fb0eff21c3d7403b1696beae77fa4762ab4fad2006d99dd90bc1067561e63f6cd795cd17f088380225a7281a85435be15c40b6e010fe2d356d44f20e0c7102dbd48549c5a0fd1cb4d304989583649d8373b8884b5812267b3e21820063719e68f395da7374da328d20e806fe4d031202f7f47a8249ef35e5ec75373bd12daf402945d20ff2760393e69ce26d0dc4cfa7386e19fa857dd0b97f041250a94429d1ffbcc41a5299fc6178dbc2808a9a5e036ab9e117177ae019092163ce022d735e019254e4eb54784f52d42904d148628cc08b46b6def552db9e7f08b9bad7fe0a8930c3d798aa1e63e708cbbf6a9f1fcd564826a6e190757d3825232e3bf5bc1ac9cbe2c4d3548c2c075695613e4d4652d437e98dcf8576280c26eacafa44b10aa38df8a4db4e22f0be50f360a4cfea2cbaf17ab0b8bd29fd623a51a206b8760df40573cf646c5279912a1d41b1225abf162cbd1de8ad325a26ab94dab819794f4a64110cca38ade55bc49c5b7f71195787ed890f6e25cc9622c47b1282715403700fc6633068b895c27e24e5d5d6d0a460d202cf37b2015af392c3997a81369b0f767fb56925c7255449fd3c38215c0cf50173df9b889d183cffd2aa79f64be968f52b86515b788546fffb79a0fe805e7134a35999fbc5bdec460`)
	// -------------------------------------------------------------------------------------------------------------------------

	decryptedText, err := goccavenue.DecryptPayload(encryptedText)
	assert.Nil(err)
	assert.Equal(decryptedText, plainText)
}
