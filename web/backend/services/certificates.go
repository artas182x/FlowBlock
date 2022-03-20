package services

import (
	"github.com/artas182x/FlowBlock/backend/models"
	fabricCaUtil "github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/lib/attrmgr"
)

func LoginToUser(login models.Login) (*models.User, error) {
	cert, err := fabricCaUtil.BytesToX509Cert([]byte(login.Certificate))
	if err != nil {
		return nil, err
	}

	roles := []string{}

	attributes, err := attrmgr.New().GetAttributesFromCert(cert)

	if err != nil {
		return nil, err
	}

	if attributes.Contains("ReadOthersData") {
		hfType, _, _ := attributes.Value("ReadOthersData")
		if hfType == "1" {
			roles = append(roles, "manage others data")
		}
	}

	if attributes.Contains("RequestTokenRole") {
		hfType, _, _ := attributes.Value("RequestTokenRole")
		if hfType == "1" {
			roles = append(roles, "computation")
		}
	}

	if attributes.Contains("hf.Type") {
		hfType, _, _ := attributes.Value("hf.Type")
		if hfType == "admin" {
			roles = append(roles, "admin")
		}
	}

	// TODO CERTS
	return &models.User{
		UserName: cert.Subject.CommonName,
		Login:    login,
		Roles:    roles,
	}, nil
}
