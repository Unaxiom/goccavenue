# Installation
```bash
go get -u github.com/Unaxiom/goccavenue
```

# Usage
```golang
import (
    "github.com/Unaxiom/goccavenue"
)

func main() {
    const merchantId = "9999999"
	const encryptionKey = "ABCDEFGHIJ0123456789abcdefghij01"
	goccavenue.SetMerchantId(merchantId)
	goccavenue.SetEncryptionKey(encryptionKey)

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

    encryptedText, err := goccavenue.EncryptPayload(req)
    if err != nil {
        panic(err)
    }

    decryptedText, err := goccavenue.DecryptPayload(encryptedText)
    if err != nil {
        panic(err)
    }
}
```