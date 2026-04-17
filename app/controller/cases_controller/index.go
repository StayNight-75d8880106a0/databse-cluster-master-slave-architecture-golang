package cases_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/cases_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/cases_request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Cases_Controller struct {
	service cases_service_interface.Cases_Service_Interface
}

func NewCasesControllerRegistry(cases_service cases_service_interface.Cases_Service_Interface) *Cases_Controller {
	return &Cases_Controller{
		service: cases_service,
	}
}

// @Summary Membuat entitas kasus baru
// @Description Endpoint ini bertanggung jawab untuk pembuatan kasus baru, termasuk validasi, transformasi input, dan penanganan error. Proses ini krusial dalam workflow investigasi. Asumsi: data valid, judul unik per lokasi dan waktu.
// @Tags Cases
// @Accept application/json
// @Produce json
// @Param request body cases_request.Cases_Request true "Payload kasus dalam format JSON.\n\nField 'status' hanya dapat bernilai salah satu dari: \n- 'Open': Menandakan kasus baru dibuka dan seluruh proses investigasi masih dalam tahap inisiasi. Status ini merepresentasikan fase awal dalam siklus hidup kasus, di mana pengumpulan bukti dan identifikasi pihak terkait menjadi prioritas utama. \n- 'In Progress': Mengindikasikan bahwa kasus sedang dalam proses investigasi aktif. Pada tahap ini, seluruh sumber daya dialokasikan untuk analisis, pemeriksaan saksi, dan pengembangan hipotesis. Status ini menuntut adanya pembaruan berkala dan monitoring ketat terhadap perkembangan kasus. \n- 'Closed': Menunjukkan bahwa seluruh proses investigasi telah selesai, baik karena kasus telah terpecahkan, tidak ditemukan cukup bukti, atau alasan administratif lainnya. Status ini menandai berakhirnya siklus investigasi dan seluruh data kasus diarsipkan untuk keperluan audit dan pembelajaran institusional."
// @Success 201 {object} map[string]interface{} "Kasus berhasil dibuat; response berisi data kasus baru"
// @Failure 404 {object} map[string]interface{} "Permintaan tidak valid; data input salah"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/create [post]
func (c *Cases_Controller) Create(ctx *gin.Context) {

	request := new(cases_request.Cases_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(404, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", request.Incident_Date)

	input := &cases_request.Cases_Dto{
		Case_Title:       &request.Case_Title,
		Case_Description: &request.Case_Description,
		Incident_Date:    dateParse,
		Location:         &request.Location,
		Status:           &request.Status,
		DetectiveIDs:     request.DetectiveIDs,
	}

	cases, errCreate := c.service.Create(input)

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
		"Message": "Success Create Case",
		"Data":    cases,
	})
}

// @Summary Mendapatkan seluruh data kasus
// @Description Endpoint ini mengambil seluruh data kasus dengan dukungan paginasi dan pencarian judul. Cocok untuk monitoring dan analisis kasus. Asumsi: paginasi default 10 item/halaman.
// @Tags Cases
// @Produce json
// @Param page query int false "Nomor halaman, default 1"
// @Param limit query int false "Jumlah item per halaman, default 10"
// @Param search query string false "Pencarian berdasarkan judul kasus"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan daftar kasus beserta metadata paginasi"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases [get]
func (c *Cases_Controller) GetAll(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	search := ctx.Query("search")

	cases, meta, errGet := c.service.GetAll(page, limit, search)

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
		"Message":    "Success Get Cases",
		"Pagination": meta,
		"Data":       cases,
	})

}

// @Summary Mendapatkan kasus berdasarkan ID
// @Description Endpoint ini mengambil data kasus spesifik berdasarkan UUID. Cocok untuk detail view atau proses investigasi lanjutan. Asumsi: ID valid dan eksis.
// @Tags Cases
// @Produce json
// @Param id path string true "UUID kasus yang dicari"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan data kasus"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/{id} [get]
func (c *Cases_Controller) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	cases, errGet := c.service.GetById(id)

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
		"Message": "Success Get Case By Id",
		"Data":    cases,
	})

}

// @Summary Mendapatkan jumlah kasus
// @Description Endpoint ini mengambil jumlah total kasus yang terdaftar dalam sistem. Cocok untuk statistik dan monitoring. Asumsi: data konsisten.
// @Tags Cases
// @Produce json
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan jumlah kasus"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/count [get]
func (c *Cases_Controller) GetCount(ctx *gin.Context) {

	CasesCount, errGet := c.service.GetCount()

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
		"Message": "Success Count Cases",
		"Data":    CasesCount,
	})

}

// @Summary Memperbarui data kasus
// @Description Endpoint ini memperbarui data kasus berdasarkan ID, mendukung perubahan seluruh atribut utama. Validasi dan error handling dilakukan secara ketat. Asumsi: ID valid, judul tetap unik.
// @Tags Cases
// @Accept application/json
// @Produce json
// @Param id path string true "UUID kasus yang akan diperbarui"
// @Param request body cases_request.Cases_Request true "Payload kasus dalam format JSON.\n\nField 'status' hanya dapat bernilai salah satu dari: \n- 'Open': Menandakan kasus baru dibuka dan seluruh proses investigasi masih dalam tahap inisiasi. Status ini merepresentasikan fase awal dalam siklus hidup kasus, di mana pengumpulan bukti dan identifikasi pihak terkait menjadi prioritas utama. \n- 'In Progress': Mengindikasikan bahwa kasus sedang dalam proses investigasi aktif. Pada tahap ini, seluruh sumber daya dialokasikan untuk analisis, pemeriksaan saksi, dan pengembangan hipotesis. Status ini menuntut adanya pembaruan berkala dan monitoring ketat terhadap perkembangan kasus. \n- 'Closed': Menunjukkan bahwa seluruh proses investigasi telah selesai, baik karena kasus telah terpecahkan, tidak ditemukan cukup bukti, atau alasan administratif lainnya. Status ini menandai berakhirnya siklus investigasi dan seluruh data kasus diarsipkan untuk keperluan audit dan pembelajaran institusional."
// @Success 200 {object} map[string]interface{} "Berhasil memperbarui data kasus"
// @Failure 404 {object} map[string]interface{} "Permintaan tidak valid; data input salah"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/update/{id} [put]
func (c *Cases_Controller) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	request := new(cases_request.Cases_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(404, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", request.Incident_Date)

	input := &cases_request.Cases_Dto{
		Case_Title:       &request.Case_Title,
		Case_Description: &request.Case_Description,
		Incident_Date:    dateParse,
		Location:         &request.Location,
		Status:           &request.Status,
		DetectiveIDs:     request.DetectiveIDs,
	}

	cases, errUpdate := c.service.Update(id, input)

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
		"Message": "Success Update Case",
		"Data":    cases,
	})

}

// @Summary Menghapus data kasus
// @Description Endpoint ini menghapus data kasus berdasarkan ID. Proses ini tidak dapat dibatalkan dan berdampak pada data terkait. Asumsi: ID valid dan eksis.
// @Tags Cases
// @Produce json
// @Param id path string true "UUID kasus yang akan dihapus"
// @Success 200 {object} map[string]interface{} "Berhasil menghapus kasus"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/delete/{id} [delete]
func (c *Cases_Controller) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	errDelete := c.service.Delete(id)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Succes Delete Case",
	})

}

// @Summary Mendapatkan seluruh data kasus yang terakhir diperbarui
// @Description Endpoint ini mengambil seluruh data kasus dengan dukungan paginasi dan pencarian judul. Cocok untuk monitoring dan analisis kasus.
// @Tags Cases
// @Produce json
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan daftar kasus terbaru dengan maksimal 5 data
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/cases/latest [get]
func (c *Cases_Controller) GetCasesLatestUpdate(ctx *gin.Context) {

	cases, errGet := c.service.GetCasesLatestUpdate()

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
		"Message": "Success Get Latest Cases Update",
		"Data":    cases,
	})

}
