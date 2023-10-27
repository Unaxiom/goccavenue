package goccavenue

// Stores the necessary parameters to initiate a payment request via CCAvenue
type CCAvenueRequest struct {
	MerchantId  string `json:"merchant_id"`
	OrderId     string `json:"order_id"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	RedirectUrl string `json:"redirect_url"`
	CancelUrl   string `json:"cancel_url"`
	Language    string `json:"language"`

	BillingName    string `json:"billing_name"`
	BillingAddress string `json:"billing_address"`
	BillingCity    string `json:"billing_city"`
	BillingState   string `json:"billing_state"`
	BillingZip     string `json:"billing_zip"`
	BillingCountry string `json:"billing_country"`
	BillingTel     string `json:"billing_tel"`
	BillingEmail   string `json:"billing_email"`

	DeliveryName    string `json:"delivery_name"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryCity    string `json:"delivery_city"`
	DeliveryState   string `json:"delivery_state"`
	DeliveryZip     string `json:"delivery_zip"`
	DeliveryCountry string `json:"delivery_country"`
	DeliveryTel     string `json:"delivery_tel"`

	MerchantParam1 string `json:"merchant_param1"`
	MerchantParam2 string `json:"merchant_param2"`
	MerchantParam3 string `json:"merchant_param3"`
	MerchantParam4 string `json:"merchant_param4"`
	MerchantParam5 string `json:"merchant_param5"`

	IntegrationType    string `json:"integration_type"`
	PromoCode          string `json:"promo_code"`
	CustomerIdentifier string `json:"customer_identifier"`
}

// Stores the Merchant ID
var ccAvenueMerchantId = ""

// Stores the Encryption Key as given by CCAvenue
var ccAvenueEncryptionKey = ""

// Stores the block size for encrypting the payload
const ccAvenueBlockSize = 16

// Stores the initialization vector for encrypting the payload
var ccAvenueInitializationVector = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
