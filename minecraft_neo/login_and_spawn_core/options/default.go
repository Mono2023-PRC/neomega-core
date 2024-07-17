package options

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"log"
	rand2 "math/rand"
	"os"

	"github.com/OmineDev/neomega-core/minecraft/protocol"
	"github.com/OmineDev/neomega-core/minecraft/protocol/login"
	"github.com/google/uuid"
)

func NewDefaultOptions(
	address, chainData string,
	growthLevel int,
	PrivateKey *ecdsa.PrivateKey,

) *Options {
	var err error
	opt := &Options{
		Salt:       make([]byte, 16),
		PrivateKey: PrivateKey,
		ErrorLog:   log.New(os.Stderr, "", log.LstdFlags),
	}
	_, _ = rand.Read(opt.Salt)
	opt.ClientData = defaultClientData(address, growthLevel)
	opt.Request = login.Encode(chainData, opt.ClientData, PrivateKey)
	opt.IdentityData, _, _, err = login.Parse(opt.Request)
	if err != nil {
		panic(err)
	}
	opt.ClientData.ThirdPartyName = opt.IdentityData.DisplayName
	if opt.IdentityData.DisplayName == "" {
		panic("invalid identity data: display name")
	}
	if opt.IdentityData.Identity == "" {
		panic("invalid identity data: identity in uuid")
	}
	return opt
}

//go:embed skin_resource_patch.json
var skinResourcePatch []byte

//go:embed skin_geometry.json
var skinGeometry []byte

// defaultClientData edits the ClientData passed to have defaults set to all fields that were left unchanged.
func defaultClientData(address string, growthLevel int) login.ClientData {
	d := login.ClientData{}
	d.ServerAddress = address
	d.DeviceOS = protocol.DeviceAndroid
	d.GameVersion = protocol.CurrentVersion
	d.GrowthLevel = growthLevel
	d.ClientRandomID = rand2.Int63()
	d.DeviceID = uuid.New().String()
	d.LanguageCode = "zh_CN"
	d.AnimatedImageData = make([]login.SkinAnimation, 0)
	d.PersonaPieces = make([]login.PersonaPiece, 0)
	d.PieceTintColours = make([]login.PersonaPieceTintColour, 0)
	d.SelfSignedID = uuid.New().String()
	d.SkinID = uuid.New().String()
	d.SkinData = base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{0, 0, 0, 255}, 32*64))
	d.SkinImageHeight = 32
	d.SkinImageWidth = 64
	d.SkinResourcePatch = base64.StdEncoding.EncodeToString(skinResourcePatch)
	d.SkinGeometry = base64.StdEncoding.EncodeToString(skinGeometry)
	return d
}
