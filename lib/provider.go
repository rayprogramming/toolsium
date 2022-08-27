package lib

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
)

type MfaProvider struct {
	creds aws.Credentials
}

func (p *MfaProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	if p.creds != (aws.Credentials{}) {
		return p.creds, nil
	}
	p.creds = aws.Credentials{
		AccessKeyID:     viper.GetString("session.Credentials.AccessKeyId"),
		SecretAccessKey: viper.GetString("session.Credentials.SecretAccessKey"),
		SessionToken:    viper.GetString("session.Credentials.SessionToken"),
		Expires:         viper.GetTime("session.Credentials.Expiration"),
		CanExpire:       true,
	}
	return p.creds, nil
}
