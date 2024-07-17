package fbauth

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/OmineDev/neomega-core/i18n"
)

type AccessWrapper struct {
	ServerCode     string
	ServerPassword string
	Token          string
	Client         *Client
	Username       string
	Password       string
	writeBackToken bool
}

func NewAccessWrapper(Client *Client, ServerCode, ServerPassword, Token, username, password string, writeBackToken bool) *AccessWrapper {
	return &AccessWrapper{
		Client:         Client,
		ServerCode:     ServerCode,
		ServerPassword: ServerPassword,
		Token:          Token,
		Username:       username,
		Password:       password,
		writeBackToken: writeBackToken,
	}
}

func (aw *AccessWrapper) GetAccess(ctx context.Context, publicKey []byte) (map[string]any, error) {
	pubKeyData := base64.StdEncoding.EncodeToString(publicKey)
	authResp, err := aw.Client.Auth(ctx, aw.ServerCode, aw.ServerPassword, pubKeyData, aw.Token, aw.Username, aw.Password)
	if err != nil {
		return nil, err
	}
	token, _ := authResp["token"].(string)
	if len(token) != 0 && aw.writeBackToken {
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(i18n.T(i18n.S_cannot_find_user_home_dir_save_token_in_current_dir))
			homedir = "."
		}
		fbconfigdir := filepath.Join(homedir, ".config", "fastbuilder")
		os.MkdirAll(fbconfigdir, 0755)
		ptoken := filepath.Join(fbconfigdir, "fbtoken")
		// 0600: -rw-------
		token_file, err := os.OpenFile(ptoken, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil, err
		}
		_, err = token_file.WriteString(token)
		if err != nil {
			return nil, err
		}
		token_file.Close()
	}
	return authResp, nil
}
