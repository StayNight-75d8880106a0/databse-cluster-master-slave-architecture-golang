package detective_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/detective_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/detective_request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Detective_Controller struct {
	service detective_service_interface.Detective_Service_Interface
}

func NewDetectiveControllerRegistry(detective_service detective_service_interface.Detective_Service_Interface) *Detective_Controller {
	return &Detective_Controller{
		service: detective_service,
	}
}

// @Summary Membuat entitas detektif baru
// @Description Endpoint ini melakukan proses pembuatan detektif baru, termasuk validasi data, transformasi input, dan penanganan error. Proses ini penting untuk memastikan integritas data personel investigasi. Asumsi: data valid, badge_number unik. Batasan: hanya admin dapat akses.
// @Tags Detectives
// @Accept application/json
// @Produce json
// @Param request body detective_request.Detective_Request true "Payload detektif dalam format JSON.\n\nField 'investigation_style' hanya dapat bernilai salah satu dari: \n- 'Evidence-Based Investigation': Pendekatan investigasi yang menitikberatkan pada pengumpulan, analisis, dan interpretasi bukti fisik secara sistematis. Metode ini mengedepankan objektivitas dan validitas ilmiah dalam setiap tahapan investigasi. \n- 'Interview-Based Investigation': Strategi investigasi yang berfokus pada teknik wawancara mendalam terhadap saksi, korban, dan tersangka. Pendekatan ini menuntut keterampilan komunikasi interpersonal dan analisis psikologis tingkat tinggi. \n- 'Undercover Investigation': Metode investigasi di mana detektif menyamar untuk memperoleh informasi yang tidak dapat diakses secara terbuka. Pendekatan ini sangat berisiko namun efektif dalam mengungkap kejahatan terorganisir. \n- 'Follow The Money Investigation': Investigasi yang memprioritaskan pelacakan aliran dana untuk mengidentifikasi motif, pelaku, dan jaringan kejahatan. Pendekatan ini sangat relevan dalam kasus korupsi dan kejahatan finansial. \n- 'Report-Based Investigation': Pendekatan yang mengandalkan analisis laporan, dokumen, dan data administratif sebagai sumber utama informasi. Cocok untuk kasus yang membutuhkan audit forensik dan penelusuran administratif secara mendalam."
// @Success 201 {object} map[string]interface{} "Detektif berhasil dibuat; response berisi data detektif baru"
// @Failure 400 {object} map[string]interface{} "Permintaan tidak valid; data input salah atau format tidak sesuai"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/detective/create [post]
func (c *Detective_Controller) Create(ctx *gin.Context) {

	request := new(detective_request.Detective_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &detective_request.Detective_Dto{
		Name:                &request.Name,
		Badge_Number:        &request.Badge_Number,
		Department:          &request.Department,
		Station:             &request.Station,
		Phone:               &request.Phone,
		Investigation_Style: &request.Investigation_Style,
	}

	detective, errCreate := c.service.Create(input)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errCreate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Create Detective",
		"Data":    detective,
	})

}

// @Summary Mendapatkan seluruh data detektif
// @Description Endpoint ini mengambil seluruh data detektif dengan dukungan paginasi. Cocok untuk monitoring personel investigasi. Asumsi: data konsisten, paginasi default 10 item per halaman.
// @Tags Detectives
// @Produce json
// @Param page query int false "Nomor halaman, default 1"
// @Param limit query int false "Jumlah item per halaman, default 10"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan daftar detektif beserta metadata paginasi"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/detective [get]
func (c *Detective_Controller) GetAll(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	detective, meta, errGet := c.service.GetAll(page, limit)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message":    "Success Get Detective",
		"Pagination": meta,
		"Data":       detective,
	})

}

// @Summary Mendapatkan detektif berdasarkan ID
// @Description Endpoint ini mengambil data detektif spesifik berdasarkan UUID. Cocok untuk detail view atau proses assignment kasus. Asumsi: ID valid dan eksis.
// @Tags Detectives
// @Produce json
// @Param id path string true "UUID detektif yang dicari"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan data detektif"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/detective/{id} [get]
func (c *Detective_Controller) GetById(ctx *gin.Context) {

	ID := ctx.Param("id")

	detective, errGet := c.service.GetById(ID)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Data By Id",
		"Data":    detective,
	})
}

// @Summary Memperbarui data detektif
// @Description Endpoint ini memperbarui data detektif berdasarkan ID, mendukung perubahan seluruh atribut utama. Validasi dan error handling dilakukan secara ketat. Asumsi: ID valid, badge_number tetap unik.
// @Tags Detectives
// @Accept application/json
// @Produce json
// @Param id path string true "UUID detektif yang akan diperbarui"
// @Param request body detective_request.Detective_Request true "Payload detektif dalam format JSON.\n\nField 'investigation_style' hanya dapat bernilai salah satu dari: \n- 'Evidence-Based Investigation': Pendekatan investigasi yang menitikberatkan pada pengumpulan, analisis, dan interpretasi bukti fisik secara sistematis. Metode ini mengedepankan objektivitas dan validitas ilmiah dalam setiap tahapan investigasi. \n- 'Interview-Based Investigation': Strategi investigasi yang berfokus pada teknik wawancara mendalam terhadap saksi, korban, dan tersangka. Pendekatan ini menuntut keterampilan komunikasi interpersonal dan analisis psikologis tingkat tinggi. \n- 'Undercover Investigation': Metode investigasi di mana detektif menyamar untuk memperoleh informasi yang tidak dapat diakses secara terbuka. Pendekatan ini sangat berisiko namun efektif dalam mengungkap kejahatan terorganisir. \n- 'Follow The Money Investigation': Investigasi yang memprioritaskan pelacakan aliran dana untuk mengidentifikasi motif, pelaku, dan jaringan kejahatan. Pendekatan ini sangat relevan dalam kasus korupsi dan kejahatan finansial. \n- 'Report-Based Investigation': Pendekatan yang mengandalkan analisis laporan, dokumen, dan data administratif sebagai sumber utama informasi. Cocok untuk kasus yang membutuhkan audit forensik dan penelusuran administratif secara mendalam."
// @Success 200 {object} map[string]interface{} "Berhasil memperbarui data detektif"
// @Failure 400 {object} map[string]interface{} "Permintaan tidak valid; data input salah"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/detective/update/{id} [put]
func (c *Detective_Controller) Update(ctx *gin.Context) {

	ID := ctx.Param("id")

	request := new(detective_request.Detective_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &detective_request.Detective_Dto{
		Name:                &request.Name,
		Badge_Number:        &request.Badge_Number,
		Department:          &request.Department,
		Station:             &request.Station,
		Phone:               &request.Phone,
		Investigation_Style: &request.Investigation_Style,
	}

	detective, errUpdate := c.service.Update(ID, input)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errUpdate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errUpdate.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Update Detective",
		"Data":    detective,
	})

}

// @Summary Menghapus data detektif
// @Description Endpoint ini menghapus data detektif berdasarkan ID. Proses ini tidak dapat dibatalkan dan berdampak pada assignment kasus terkait. Asumsi: ID valid dan eksis.
// @Tags Detectives
// @Produce json
// @Param id path string true "UUID detektif yang akan dihapus"
// @Success 200 {object} map[string]interface{} "Berhasil menghapus detektif"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/detective/delete/{id} [delete]
func (c *Detective_Controller) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Delete Detective",
	})
}
