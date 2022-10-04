package v1

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/note_project/pkg/logger"
	"github.com/note_project/pkg/structures"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
)

// @Summary Register User
// @Description This API for registration new user
// @Tags Users
// @Accept json
// @Produse json
// @Param user body structures.UserStruct true "user body"
// @Success 200 {string} Success
// @Router /v1/register/ [post]
func (h handlerV1) RegisterUser(c *gin.Context) {
	var body1 structures.UserStruct

	err := c.ShouldBindJSON(&body1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
	}

	fmt.Println("---------------------------------------------")

	body1.Email = strings.TrimSpace(body1.Email)
	body1.Email = strings.ToLower(body1.Email)


	UserStatus, err := h.userStorage.CheckField(structures.UserCheckRequest{
		Field: "username",
		Value: body1.UserName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		h.log.Error("failed while calling chechfield function with username")
		return
	}
	fmt.Println("---------------------------------------------")

	if !UserStatus {
		EmailStatus, err := h.userStorage.CheckField(structures.UserCheckRequest{
			Field: "email",
			Value: body1.UserName,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			h.log.Error("failed while calling chechfield function with email")
			return
		}

		if EmailStatus {
			c.JSON(http.StatusConflict, gin.H{
				"error": "user_name already in use",
			})
			h.log.Error("user already exists")
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user_name already in use",
		})
		h.log.Error("user_name already exists")
		return
	}
	fmt.Println("---------------------------------------------")

	// verPass-d...
	err = verifyPassword(body1.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("password verify error")
		return
	}
	fmt.Println("-------------------------------------------2--")
	// hashing the password
	hashedPassqord, err := bcrypt.GenerateFromPassword([]byte(body1.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while hashing password", logger.Error(err))
		return
	}
	fmt.Println("--------------------------------------------1-")
	body1.Password = string(hashedPassqord)

	min := 99999
	max := 1000000

	rand.Seed(time.Now().UnixNano())
	Code := rand.Intn(max-min) + min

	strCode := strconv.Itoa(Code)

	fmt.Println("---------------------------------------------0")
	SendEmail(body1.Email, strCode)
	fmt.Println("---------------------------------------------0")
	body1.EmailCode = strCode

	setBodyRedis, err := json.Marshal(body1)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed swhile marashling body1", logger.Error(err))
		return
	}

	err = h.redisStorage.Set(body1.Email, string(setBodyRedis))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while setting to redis", logger.Error(err))
		return
	}

}

// @Summary Verify User
// @Description This api for sending email code to user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body structures.EmailVer true "user body"
// @Success 200 {string} Success
// @Router /v1/users/verify_user/ [post]
func (h handlerV1) VerifyUser(c *gin.Context) {
	var mailData structures.EmailVer

	err := c.ShouldBindJSON(&mailData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	mailData.Email = strings.TrimSpace(mailData.Email)
	mailData.Email = strings.ToLower(mailData.Email)

	redisBody, err := redis.String(h.redisStorage.Get(mailData.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		h.log.Error("failed while getting from redis", logger.Error(err))
		return
	}

	var body structures.UserStruct

	err = json.Unmarshal([]byte(redisBody), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to unmarshal body", logger.Error(err))
		return
	}

	if mailData.Email == body.Email {
		createdUser, err := h.userStorage.CreateUser(&body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while creating user", logger.Error(err))
			return
		}
		c.JSON(http.StatusOK, createdUser)
	}

}

func SendEmail(email, code string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "postgrespostgresovnt@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)
	// id,err := uuid.NewUUID()
	// if err != nil {
	//   fmt.Println(err)
	// }
	// Set E-Mail subject
	m.SetHeader("code:", "Verification code")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", code)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "postgrespostgresovnt@gmail.com", "qmxlgijkvuuoacrh")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 32
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
