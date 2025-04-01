package mdm_firebase_credentials

import (
	"context"
	"os"

	"cloud.google.com/go/auth"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)


func getCredentialJson() string {
	// GOOGLE_APPLICATION_CREDENTIALS_DATA 環境変数が指定されていればその値を利用する (ECS で起動する場合).
	if googleCreds, credsExist := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS_DATA"); credsExist {
		return googleCreds
	}

	// GOOGLE_APPLICATION_CREDENTIALS_SSM_PARAMETER 環境変数が指定されていれば、パラメターの値を取得する (Lambda で起動する場合).
	if ssmParameter, credsExist := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS_SSM_PARAMETER"); credsExist {
		return ssmParameter
	}
	return ""
};

func getCredential() *auth.Credentials{
	return auth.NewCredentials(&auth.CredentialsOptions{JSON: []byte(getCredentialJson())})
}

func GetApp(ctx context.Context) (*firebase.App, error) {
	_, isEmulatorHost := os.LookupEnv("FIREBASE_AUTH_EMULATOR_HOST"); if isEmulatorHost {
		config := &firebase.Config{}
		// Firebase Local Emulator Suite を利用.
    	// https://firebase.google.com/docs/emulator-suite/connect_auth
		config.ProjectID = os.Getenv("FIREBASE_EMULATOR_PROJECT_ID")
		return firebase.NewApp(ctx, config, nil)
	}
  	// 本物の Firebase を利用.
	return firebase.NewApp(ctx, nil, option.WithAuthCredentials(getCredential()))
};
