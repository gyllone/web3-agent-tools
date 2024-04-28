package main

/*
#cgo CFLAGS: -I../../dependencies
#include <key.h>
*/
import "C"
import (
	"coinmarketcap/utils"
	"encoding/json"
	"net/http"
	"net/url"
)

//export query_info
func query_info() C.Result_Info {
	u, err := url.Parse(InfoUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Info(C.CString(errStr))
	}

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_Info(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_Info(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Info(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody InfoResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
		return C.err_Info(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Info(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.Info{
		plan: C.Plan{
			credit_limit_monthly:           C.Int(respData.Plan.CreditLimitMonthly),
			credit_limit_monthly_reset:     C.CString(respData.Plan.CreditLimitMonthlyReset),
			credit_limit_monthly_reset_UTC: C.CString(respData.Plan.CreditLimitMonthlyResetTimestamp.String()),
			rate_limit_minute:              C.Int(respData.Plan.RateLimitMinute),
		},
		usage: C.Usage{
			current_minute: C.CurrentMinute{
				requests_made: C.Int(respData.Usage.CurrentMinute.RequestsMade),
				requests_left: C.Int(respData.Usage.CurrentMinute.RequestsLeft),
			},
			current_day: C.CurrentDay{
				credits_used: C.Int(respData.Usage.CurrentDay.CreditsUsed),
			},
			current_month: C.CurrentMonth{
				credits_used: C.Int(respData.Usage.CurrentMonth.CreditsUsed),
				credits_left: C.Int(respData.Usage.CurrentMonth.CreditsLeft),
			},
		},
	}

	return C.ok_Info(data)
}

//export query_info_release
func query_info_release(result C.Result_Info) {
	C.release_Result_Info(result)
}

func main() {}
