package productController

import (
	"encoding/json"
	"net/http"

	"github.com/gagassurya19/go-rest-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	// mengambil struct product
	var products []models.Product

	// mengambil data dari database berdasarkan struct product
	models.DB.Find(&products)

	// return statusOK(200) dan return data product
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	// mengambil struct product
	var product models.Product

	// tangkap ID dari parameter
	id := c.Param("id")

	// cek apakah data dengan id tersebut ada
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// return statusOK(200) dan return data
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	// mengambil struct product
	var product models.Product

	// ambil data dari body
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// push data product dari body ke database
	models.DB.Create(&product)

	// return statusOK(200) dan return data
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	// mengambil struct product
	var product models.Product

	// ambil id dari param
	id := c.Param("id")

	// ambil data dari body berdasar pada struct product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// update data berdasar id dan cek apakah ada perubahan
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate produk"})
		return
	}

	// return statusOK(200) dan return message berhasil di update
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di update"})
}

func Delete(c *gin.Context) {
	// mengambil struct product
	var product models.Product

	// input := map[string]string{"id": "0"} // mengatasi "json: cannot unmarshal number into Go value of type string"
	
	// struc untuk input 
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// id, _ := strconv.ParseInt(input["id"], 10, 64) // mengatasi "json: cannot unmarshal number into Go value of type string"
	id, _ := input.Id.Int64() // ambil ID dari body dan konversi input ke int64

	// delete data berdasar id dan cek apakah ada perubahan row
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data"})
		return
	}

	// return statusOK(200) dan return message data berhasil dihapus
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
