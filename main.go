package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Nerzal/gocloak/v11"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	client := gocloak.NewClient("http://127.0.0.1:9080/")
	ctx := context.Background()
	token, err := client.Login(ctx, "smb-client", "98691423-a407-48a6-92da-a032e8c2910c", "smbRecruitmentBe", "dongocanh96", "ngocanh8")
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	var jwtToken *jwt.Token
	jwtToken, _, err = client.DecodeAccessToken(ctx, token.AccessToken, "smbRecruitmentBe")
	if err != nil {
		fmt.Println(err.Error())
	}

	keyData, err := ioutil.ReadFile("./token.key")
	if err != nil {
		fmt.Println(err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		fmt.Println(err)
	}

	parts := strings.Split(jwtToken.Raw, ".")
	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Verify success!")
}
