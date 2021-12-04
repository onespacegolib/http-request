package requests

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"testing"
)

func TestCallGet(t *testing.T) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImU3Yzg1MTFkNTU5MDkwZTRmZDEwMWVmNDU4YjBiNjRjMzM4NzRkMmU2MDZmOGI5MDllMTRkNDVmYzY2OGMxNjk0YTY5ZmU5YTZhNGQ2NGQ4In0.eyJhdWQiOiI1MTQiLCJqdGkiOiJlN2M4NTExZDU1OTA5MGU0ZmQxMDFlZjQ1OGIwYjY0YzMzODc0ZDJlNjA2ZjhiOTA5ZTE0ZDQ1ZmM2NjhjMTY5NGE2OWZlOWE2YTRkNjRkOCIsImlhdCI6MTYyNDg1MjQ4NywibmJmIjoxNjI0ODUyNDg3LCJleHAiOjE2MjUxMTE2ODcsInN1YiI6IjI1MjU2OTE0NjI0Iiwic2NvcGVzIjpbXX0.O2ZlJauovIsIHxdYxSF5j1jDSwt3sKQ42XzlvoicErtWjkEVSpl4Wy7kj4Urs12TWlJIHhSUelp6ABAXIDQIkWP9MLuwY7N4eA5EQztvFEyHWK8YiegN62wwuyHweMTUPBmbu6kRye7U5qQ2U8_iCugLCECuNCCXi4IBeoIR2QCjWLPF08UfOSC3S7aqOkBNZIRn2xy_MRZnqwpHDiy-EYrCavzN40p0wOEn27YFT8N-FknUFfsf3vP2x64SzKJFzNpy0SnQQpPI3p4k68PYN5hdYwr-b9lw5k9bCUN7rAaVvHzV9t0oATHjM8BTquAjqFpKz-KfVdxmMcTbEU9gSD7LhtSrP7JvkbgVZWDstsvjBU7iNy5iQ4ZjeA0eBcz9JDHpXuOgGhif0QTZW3rN-cn7ZeoNobUmQ4KLBnvHXxnWGfW1FKAhlOButCIen4nmPGrGzUfDPFzti3xstdL43HFrHpJ230umLhzHI98VSzDbfsigtWLhHBdLqJSF9DoTLf4eOWb2_tW2Iuz2FqyzlYWVTkS_2w_12iHJ46wn-tmPLKY1Ih4dSaDNIqodO22yZepkeWSvi5mjOF5xfdLmrZIAWlaBfA_ROHFzabPXo0ROYmc9YgOciJuKZR3cDyFmjviidtwVCD1ICWrOtyqllJDXhseDUKO6r6nOtbvawfo"
	param := Params{
		URL:  "https://testoneid.inet.co.th/api/account",
		BODY: nil,
		HEADERS: map[string]string{
			echo.HeaderContentType:   "application/json",
			echo.HeaderAuthorization: "Bearer " + token,
		},
		TIMEOUT: 5,
	}
	var res Response
	if err := Call().Get(param, &res).Error(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
