package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"upload_image/middleware"

	"github.com/gofiber/fiber/v2"
)

// @summary 	  	UPLOAD IMAGE
// @Description	  	UPLOAD IMAGE
// @Tags		  	KPLUS
// @Accept		  	json
// @Produce		  	json
// @Param       	img_src body ImageSrc true  "Complete image path"
// @Success		  	200 {object} Uploaded_Images
// @Router			/upload [post]
func UploadImage(c *fiber.Ctx) error {
	inputImgSrc := ImageSrc{}

	// var upload Uploaded_Images
	if parsErr := c.BodyParser(&inputImgSrc); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	// func ConvertToBase64(b []byte) string {
	// 	return base64.StdEncoding.EncodeToString(b)
	// }

	decodedImage, decodErr := base64.StdEncoding.DecodeString(inputImgSrc.ImgSrc)

	if decodErr != nil {

		return c.JSON(fiber.Map{
			"error": decodErr,
		})
	}
	imgStruct := Uploaded_Images{
		ImgData: decodedImage,
		ImgType: "mimeType",
	}

	if dbErr := middleware.DBConn.Debug().Create(&imgStruct).Error; dbErr != nil {
		return c.JSON(fiber.Map{
			"error": dbErr,
		})
	}

	return c.JSON(fiber.Map{
		"Result": imgStruct,
	})
}

// @summary 	  	FETCH IMAGE DATA
// @Description	  	FETCH IMAGE DATA
// @Tags		  	KPLUS
// @Accept		  	json
// @Produce		  	json
// @Param       	img_id body ImageID true  "Image ID"
// @Success		  	200 {object} Uploaded_Images
// @Router			/fetch [post]
func FetchByte(c *fiber.Ctx) error {
	imgStruct := Uploaded_Images{}

	imgID := ImageID{}

	c.BodyParser(&imgID)

	middleware.DBConn.Debug().Find(&imgStruct, "img_id=?", imgID.ImgId)
	if imgStruct.ImgID == 0 {
		return c.JSON(fiber.Map{
			"result": "Not found",
		})
	}

	return c.JSON(fiber.Map{
		"img_data": imgStruct,
	})
}

// @summary 	  	FETCH IMAGE DATA
// @Description	  	FETCH IMAGE DATA
// @Tags		  	KPLUS
// @Accept		  	json
// @Produce		  	json
// @Param       	img_id path int true  "Image ID"
// @Success		  	200 {object} ImageSrc
// @Router			/fetch_image_data/{img_id} [get]
func GetImageData(c *fiber.Ctx) error {
	imgID := c.Params("img_id")
	imgStruct := Uploaded_Images{}
	middleware.DBConn.Debug().Find(&imgStruct, "img_id=?", imgID)
	if imgStruct.ImgID == 0 {
		return c.JSON(fiber.Map{
			"result": "Not found",
		})
	}

	return c.JSON(fiber.Map{
		"img_data": imgStruct.ImgData,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	KPLUS
// @Accept		  	json
// @Produce		  	json
// @Param       	userInput path string true  "User Input"
// @Success		  	200 {object} UserManagement
// @Router			/get_um/{userInput} [get]
func GetUserManagement(c *fiber.Ctx) error {
	userinput := c.Params("userInput")
	fmt.Println("UserInput:", userinput)

	umModel := []UserManagement{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.usermanagement(?)", userinput).Scan(&umModel).Error; dbErr != nil {
		return c.JSON(fiber.Map{"Error": dbErr})
	}

	return c.Status(http.StatusCreated).JSON(umModel)
}
