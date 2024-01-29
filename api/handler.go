package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-sms-verify/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeOut = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.ErrorResponse(c, err)
			return
		}

		app.SuccessResponse(c, http.StatusAccepted, "OTP sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		dataJson, _ := json.Marshal(newData)
		fmt.Println("dataJson : ", string(dataJson))

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err : ", err)
		if err != nil {
			app.ErrorResponse(c, err)
			return
		}

		app.SuccessResponse(c, http.StatusAccepted, "OTP verified successfully")
	}
}
