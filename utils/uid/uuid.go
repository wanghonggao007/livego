package uid

import (
	"encoding/base64"

	"github.com/wanghonggao007/go.uuid"
)

func NewId() string {
	id, err := uuid.NewV4()
	if err != nil {
		return err.Error()
	}
	b64 := base64.URLEncoding.EncodeToString(id.Bytes()[:12])
	return b64
}
