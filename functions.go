package goccavenue

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func escapeStr(str string) string {
	// return url.QueryEscape(str)
	return str
}

// Sets the merchant id
func SetMerchantId(merchantId string) {
	ccAvenueMerchantId = merchantId
}

// Gets the merchant id
func GetMerchantId() (merchantId string) {
	return ccAvenueMerchantId
}

// Sets the encryption key
func SetEncryptionKey(encryptionKey string) {
	ccAvenueEncryptionKey = encryptionKey
}

// Gets the encryption key
func GetEncryptionKey() (encryptionKey string) {
	return ccAvenueEncryptionKey
}

// Creates the payload string in a format acceptable to CCAvenue prior to encryption
func CreateRequest(paymentReq CCAvenueRequest) string {
	var stringsToEncode = []string{
		fmt.Sprintf("merchant_id=%s", escapeStr(paymentReq.MerchantId)),
		fmt.Sprintf("order_id=%s", escapeStr(paymentReq.OrderId)),
		fmt.Sprintf("currency=%s", escapeStr(paymentReq.Currency)),
		fmt.Sprintf("amount=%s", escapeStr(paymentReq.Amount)),
		fmt.Sprintf("redirect_url=%s", escapeStr(paymentReq.RedirectUrl)),
		fmt.Sprintf("cancel_url=%s", escapeStr(paymentReq.CancelUrl)),
		fmt.Sprintf("language=%s", escapeStr(paymentReq.Language)),

		fmt.Sprintf("billing_name=%s", escapeStr(paymentReq.BillingName)),
		fmt.Sprintf("billing_address=%s", escapeStr(paymentReq.BillingAddress)),
		fmt.Sprintf("billing_city=%s", escapeStr(paymentReq.BillingCity)),
		fmt.Sprintf("billing_state=%s", escapeStr(paymentReq.BillingState)),
		fmt.Sprintf("billing_zip=%s", escapeStr(paymentReq.BillingZip)),
		fmt.Sprintf("billing_country=%s", escapeStr(paymentReq.BillingCountry)),
		fmt.Sprintf("billing_tel=%s", escapeStr(paymentReq.BillingTel)),
		fmt.Sprintf("billing_email=%s", escapeStr(paymentReq.BillingEmail)),

		fmt.Sprintf("delivery_name=%s", escapeStr(paymentReq.DeliveryName)),
		fmt.Sprintf("delivery_address=%s", escapeStr(paymentReq.DeliveryAddress)),
		fmt.Sprintf("delivery_city=%s", escapeStr(paymentReq.DeliveryCity)),
		fmt.Sprintf("delivery_state=%s", escapeStr(paymentReq.DeliveryState)),
		fmt.Sprintf("delivery_zip=%s", escapeStr(paymentReq.DeliveryZip)),
		fmt.Sprintf("delivery_country=%s", escapeStr(paymentReq.DeliveryCountry)),
		fmt.Sprintf("delivery_tel=%s", escapeStr(paymentReq.DeliveryTel)),

		fmt.Sprintf("merchant_param1=%s", escapeStr(paymentReq.MerchantParam1)),
		fmt.Sprintf("merchant_param2=%s", escapeStr(paymentReq.MerchantParam2)),
		fmt.Sprintf("merchant_param3=%s", escapeStr(paymentReq.MerchantParam3)),
		fmt.Sprintf("merchant_param4=%s", escapeStr(paymentReq.MerchantParam4)),
		fmt.Sprintf("merchant_param5=%s", escapeStr(paymentReq.MerchantParam5)),

		fmt.Sprintf("integration_type=%s", escapeStr(paymentReq.IntegrationType)),
		fmt.Sprintf("promo_code=%s", escapeStr(paymentReq.PromoCode)),
		fmt.Sprintf("customer_identifier=%s", escapeStr(paymentReq.CustomerIdentifier)),

		fmt.Sprintf("merchant_id=%s", paymentReq.MerchantId),
	}

	return strings.Join(stringsToEncode, "&")
}

// Creates the response payload from the already decrypted string
func CreateResponseFromDecryptedText(decryptedString string) (response CCAvenueResponse, err error) {
	var resp = CCAvenueResponse{}
	// Create a map
	var dict = make(map[string]string)
	listOfValues := strings.Split(decryptedString, "&")
	for _, combined := range listOfValues {
		split := strings.Split(combined, "=")
		if len(split) != 2 {
			continue
		}
		key := split[0]
		value := split[1]
		dict[key] = value
	}

	// Assign the values here
	resp.OrderId = dict["order_id"]
	resp.TrackingId = dict["tracking_id"]
	resp.BankRefNo = dict["bank_ref_no"]
	resp.OrderStatus = strToOrderStatus(dict["order_status"])
	resp.FailureMessage = dict["failure_message"]
	resp.PaymentMode = dict["payment_mode"]
	resp.CardName = dict["card_name"]
	resp.StatusCode = dict["status_code"]
	resp.StatusMessage = dict["status_message"]
	resp.Currency = dict["currency"]
	resp.Amount = dict["amount"]
	resp.BillingName = dict["billing_name"]
	resp.BillingAddress = dict["billing_address"]
	resp.BillingCity = dict["billing_city"]
	resp.BillingState = dict["billing_state"]
	resp.BillingZip = dict["billing_zip"]
	resp.BillingCountry = dict["billing_country"]
	resp.BillingTel = dict["billing_tel"]
	resp.BillingEmail = dict["billing_email"]
	resp.DeliveryName = dict["delivery_name"]
	resp.DeliveryAddress = dict["delivery_address"]
	resp.DeliveryCity = dict["delivery_city"]
	resp.DeliveryState = dict["delivery_state"]
	resp.DeliveryZip = dict["delivery_zip"]
	resp.DeliveryCountry = dict["delivery_country"]
	resp.DeliveryTel = dict["delivery_tel"]
	resp.MerchantParam1 = dict["merchant_param1"]
	resp.MerchantParam2 = dict["merchant_param2"]
	resp.MerchantParam3 = dict["merchant_param3"]
	resp.MerchantParam4 = dict["merchant_param4"]
	resp.MerchantParam5 = dict["merchant_param5"]
	resp.Vault = dict["vault"]
	resp.OfferType = dict["offer_type"]
	resp.OfferCode = dict["offer_code"]
	resp.DiscountValue = dict["discount_value"]
	resp.MerchantAmount = dict["mer_amount"]
	resp.ECIValue = dict["eci_value"]
	resp.Retry = dict["retry"]
	resp.ResponseCode = dict["response_code"]
	resp.BillingNotes = dict["billing_notes"]
	resp.TransactionDate = dict["trans_date"]
	resp.BinCountry = dict["bin_country"]

	return resp, nil
}

// Creates the response payload from the given encrypted string from CCAvenue
func CreateResponseFromEncryptedText(encryptedResponse string) (response CCAvenueResponse, err error) {
	if strings.HasPrefix(encryptedResponse, "encResp=") {
		encryptedResponse = strings.Split(encryptedResponse, "encResp=")[1]
	}
	decryptedString, err := DecryptPayload(encryptedResponse)
	if err != nil {
		return CCAvenueResponse{}, err
	}
	return CreateResponseFromDecryptedText(decryptedString)
}

// Pads the cipher text so that the length is a multiple of the block size (as per CCAvenue's requirement)
func padPayload(ciphertext string) (paddedString string) {
	padding := (ccAvenueBlockSize - len(ciphertext)%ccAvenueBlockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return ciphertext + string(padtext)
}

// Encrypts plain text string into cipher text string that is recognized by CCAvenue
func EncryptPayload(paymentReq CCAvenueRequest) (string, error) {
	hash := md5.Sum([]byte(GetEncryptionKey()))
	key := hash[:]
	plainText := padPayload(CreateRequest(paymentReq))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, ccAvenueInitializationVector)
	mode.CryptBlocks(cipherText, []byte(plainText))

	return hex.EncodeToString(cipherText), err
}

// Decrypts cipher text string into plain text string
func DecryptPayload(cipherText string) (decryptedString string, err error) {
	hash := md5.Sum([]byte(GetEncryptionKey()))
	key := hash[:]
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddedBytes := make([]byte, len(cipherTextDecoded))

	mode := cipher.NewCBCDecrypter(block, ccAvenueInitializationVector)
	mode.CryptBlocks(paddedBytes, cipherTextDecoded)

	paddingLen := paddedBytes[len(paddedBytes)-1]
	if paddingLen < ccAvenueBlockSize {
		// Message has been padded, so trim the last n bytes
		paddedBytes = paddedBytes[:len(paddedBytes)-int(paddingLen)]
	}

	plainText := string(paddedBytes)

	return plainText, nil
}
