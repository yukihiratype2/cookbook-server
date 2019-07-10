package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	// "github.com/yukihiratype2/cookbook-server/internal/obs"
)

type uploadAssetResp struct {
	PresignedURL string `json:"presigned_url"`
	SHA2         string `json:"SHA2"`
}

func (ph *postHandler) UploadAsset(c echo.Context) (err error) {
	asset := new(apim.PostAssetParams)
	if err = c.Bind(asset); err != nil {
		return
	}
	fmt.Println(asset)
	str, err := ph.obs.GenPresign(asset.SHA2)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusCreated, &uploadAssetResp{str, asset.SHA2})
	return
}
