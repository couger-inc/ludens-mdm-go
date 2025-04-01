package userconsole

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func headers() (string, error) {
	service, serviceExists := os.LookupEnv("SERVICE");
	serviceKid, serviceKidExists := os.LookupEnv("SERVICE_KID")
	serviceKey, serviceKeyExists := os.LookupEnv("SERVICE_KEY")
	issuer := fmt.Sprintf("https://couger.co.jp/service/%v", service)
	if (!serviceExists) {
		return "", errors.New("SERVICE is not set")
	}
	if (!serviceKidExists) {
		return "", errors.New("SERVICE_KID is not set")
	}
  	if (!serviceKeyExists) {
		return "", errors.New("SERVICE_KEY is not set")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"iss": issuer,
	})
	t.Header["kid"] = serviceKid
	s, err := t.SignedString([]byte(serviceKey))
	return s, err
}


func GetUsers(email string) (*GetUsersResponse, error) {
	identityApiBasePath := os.Getenv("IDENTITY_API_BASE_PATH")
	usersUrl := fmt.Sprintf("%v/users?q=%v", identityApiBasePath, url.QueryEscape(email));
	req, err := http.NewRequest("GET", usersUrl, nil)
	if err != nil {
		return nil, err
	}
	if auth, err := headers(); err == nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", auth))
		log.Printf("Calling %v with params: [%v]", usersUrl, email)
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		responseData, err := io.ReadAll(response.Body)
		log.Printf("StatusCode: %v, Body: %v", response.StatusCode, string(responseData))
		if err != nil {
			return nil, err
		}
		if (response.StatusCode != 200) {
			return nil, fmt.Errorf("unable to query luden's user console. StatusCode: %v, Message: %v", response.StatusCode, string(responseData))
		}
		var usersResponse GetUsersResponse
		err = json.Unmarshal(responseData, &usersResponse)
		if (err != nil) {
			return nil, err
		}
		log.Printf("GOT RESPONSE: Users: %v, TotalCount: %v", usersResponse.Users, usersResponse.TotalCount)
		return &usersResponse, nil
	} else {
		return nil, err
	}
	
}
