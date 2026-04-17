package suspect_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/suspect_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/suspects_request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Suspect_Controller struct {
	service suspect_service_interface.Suspect_Service_Interface
}

func NewSuspectControllerRegistry(suspect_service suspect_service_interface.Suspect_Service_Interface) *Suspect_Controller {
	return &Suspect_Controller{
		service: suspect_service,
	}
}

// @Summary Membuat entitas tersangka baru
// @Description Endpoint ini bertanggung jawab untuk pembuatan tersangka baru pada kasus tertentu, termasuk validasi, transformasi input, dan penanganan error. Proses ini penting untuk integritas data investigasi. Asumsi: ID kasus valid, NIK unik.
// @Tags Suspects
// @Accept application/json
// @Produce json
// @Param id_case path string true "UUID kasus, wajib diisi"
// @Param request body suspects_request.Suspects_Request true "Payload tersangka dalam format JSON.\n\nField 'gender' hanya dapat bernilai: \n- 'Male': Menunjukkan jenis kelamin laki-laki. Dalam perspektif kriminologi dan psikologi forensik, identifikasi gender ini penting untuk analisis pola perilaku, kecenderungan kriminal, serta pendekatan interogasi dan perlindungan hukum yang relevan. Gender laki-laki sering dikaitkan dengan karakteristik fisik dan psikososial tertentu yang dapat mempengaruhi jalannya investigasi.\n- 'Female': Menunjukkan jenis kelamin perempuan. Gender ini memiliki implikasi signifikan dalam proses investigasi, terutama terkait perlindungan saksi, pendekatan psikologis, dan sensitivitas terhadap isu gender-based violence. Penanganan tersangka atau saksi perempuan menuntut penerapan prinsip non-diskriminasi dan perlakuan khusus sesuai standar hak asasi manusia.\n\nField 'status' hanya dapat bernilai: \n- 'Arrested': Status ini menandakan bahwa tersangka telah resmi ditahan oleh aparat penegak hukum. Dalam sistem peradilan pidana, status ini menuntut pemenuhan hak-hak hukum tersangka, termasuk akses bantuan hukum dan perlindungan dari perlakuan sewenang-wenang. Penahanan merupakan fase kritis yang mempengaruhi kelanjutan proses investigasi dan peradilan.\n- 'Released': Menunjukkan bahwa tersangka telah dibebaskan, baik karena tidak cukup bukti, alasan hukum, atau pertimbangan kemanusiaan. Status ini penting dalam evaluasi efektivitas investigasi dan menjadi indikator penerapan asas praduga tak bersalah.\n- 'Wanted': Status ini berarti tersangka sedang dalam pencarian aktif oleh aparat penegak hukum. Penetapan status 'wanted' memerlukan koordinasi lintas lembaga, publikasi identitas, dan penggunaan sumber daya intelijen untuk penangkapan.\n- 'Under Investigation': Menandakan bahwa individu masih dalam proses penyelidikan dan status hukumnya belum final. Status ini menuntut kehati-hatian dalam publikasi data dan perlakuan terhadap individu, guna menjaga asas praduga tak bersalah dan integritas proses hukum.\n- 'Eyewitness': Status ini diberikan kepada individu yang berperan sebagai saksi mata dalam kasus. Saksi mata memiliki nilai probatif tinggi dalam pembuktian perkara, sehingga perlindungan dan penanganan profesional sangat diperlukan untuk menjaga keamanan dan keabsahan kesaksian."
// @Failure 400 {object} map[string]interface{} "Permintaan tidak valid; data input salah"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/suspect/create/{id_case} [post]
func (c *Suspect_Controller) Create(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", request.Date_Of_Birth)

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: &request.ID_Card_Number,
		Full_Name:      &request.Full_Name,
		Gender:         &request.Gender,
		Date_Of_Birth:  dateParse,
		Address:        &request.Address,
		Phone:          &request.Phone,
		Occupation:     &request.Occupation,
		Alibi:          &request.Alibi,
		Status:         &request.Status,
	}

	suspect, errCreate := c.service.Create(ID_Case, input)

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
		"Messaage": "Success Create Suspect",
		"Data":     suspect,
	})

}

// @Summary Mendapatkan seluruh data tersangka pada kasus
// @Description Endpoint ini mengambil seluruh data tersangka untuk kasus tertentu dengan dukungan paginasi. Cocok untuk monitoring investigasi. Asumsi: ID kasus valid.
// @Tags Suspects
// @Produce json
// @Param id_case path string true "UUID kasus, wajib diisi"
// @Param page query int false "Nomor halaman, default 1"
// @Param limit query int false "Jumlah item per halaman, default 10"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan daftar tersangka beserta metadata paginasi"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/suspect/get-all/{id_case} [get]
func (c *Suspect_Controller) GetAll(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	suspect, meta, errGet := c.service.GetAll(ID_Case, page, limit)

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
		"Message":    "Success Get Suspect Data",
		"Data":       suspect,
		"Pagination": meta,
	})

}

// @Summary Mendapatkan tersangka berdasarkan ID pada kasus
// @Description Endpoint ini mengambil data tersangka spesifik berdasarkan UUID pada kasus tertentu. Cocok untuk detail view investigasi. Asumsi: ID kasus dan ID tersangka valid.
// @Tags Suspects
// @Produce json
// @Param id_case path string true "UUID kasus, wajib diisi"
// @Param id path string true "UUID tersangka, wajib diisi"
// @Success 200 {object} map[string]interface{} "Berhasil mendapatkan data tersangka"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/suspect/get-id/{id}/{id_case} [get]
func (c *Suspect_Controller) GetById(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	suspect, errGet := c.service.GetById(ID, ID_Case)

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
		"Message": "Success Get Suspect Data",
		"Data":    suspect,
	})

}

// @Summary Memperbarui data tersangka
// @Description Endpoint ini memperbarui data tersangka pada kasus tertentu berdasarkan ID, mendukung perubahan seluruh atribut utama. Validasi dan error handling dilakukan secara ketat. Asumsi: ID kasus dan ID tersangka valid.
// @Tags Suspects
// @Accept application/json
// @Produce json
// @Param id_case path string true "UUID kasus, wajib diisi"
// @Param id path string true "UUID tersangka, wajib diisi"
// @Param request body suspects_request.Suspects_Request true "Payload tersangka dalam format JSON.\n\nField 'gender' hanya dapat bernilai: \n- 'Male': Menunjukkan jenis kelamin laki-laki. Dalam perspektif kriminologi dan psikologi forensik, identifikasi gender ini penting untuk analisis pola perilaku, kecenderungan kriminal, serta pendekatan interogasi dan perlindungan hukum yang relevan. Gender laki-laki sering dikaitkan dengan karakteristik fisik dan psikososial tertentu yang dapat mempengaruhi jalannya investigasi.\n- 'Female': Menunjukkan jenis kelamin perempuan. Gender ini memiliki implikasi signifikan dalam proses investigasi, terutama terkait perlindungan saksi, pendekatan psikologis, dan sensitivitas terhadap isu gender-based violence. Penanganan tersangka atau saksi perempuan menuntut penerapan prinsip non-diskriminasi dan perlakuan khusus sesuai standar hak asasi manusia.\n\nField 'status' hanya dapat bernilai: \n- 'Arrested': Status ini menandakan bahwa tersangka telah resmi ditahan oleh aparat penegak hukum. Dalam sistem peradilan pidana, status ini menuntut pemenuhan hak-hak hukum tersangka, termasuk akses bantuan hukum dan perlindungan dari perlakuan sewenang-wenang. Penahanan merupakan fase kritis yang mempengaruhi kelanjutan proses investigasi dan peradilan.\n- 'Released': Menunjukkan bahwa tersangka telah dibebaskan, baik karena tidak cukup bukti, alasan hukum, atau pertimbangan kemanusiaan. Status ini penting dalam evaluasi efektivitas investigasi dan menjadi indikator penerapan asas praduga tak bersalah.\n- 'Wanted': Status ini berarti tersangka sedang dalam pencarian aktif oleh aparat penegak hukum. Penetapan status 'wanted' memerlukan koordinasi lintas lembaga, publikasi identitas, dan penggunaan sumber daya intelijen untuk penangkapan.\n- 'Under Investigation': Menandakan bahwa individu masih dalam proses penyelidikan dan status hukumnya belum final. Status ini menuntut kehati-hatian dalam publikasi data dan perlakuan terhadap individu, guna menjaga asas praduga tak bersalah dan integritas proses hukum.\n- 'Eyewitness': Status ini diberikan kepada individu yang berperan sebagai saksi mata dalam kasus. Saksi mata memiliki nilai probatif tinggi dalam pembuktian perkara, sehingga perlindungan dan penanganan profesional sangat diperlukan untuk menjaga keamanan dan keabsahan kesaksian."
// @Success 200 {object} map[string]interface{} "Berhasil memperbarui data tersangka"
// @Failure 400 {object} map[string]interface{} "Permintaan tidak valid; data input salah"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/suspect/update/{id}/{id_case} [put]
func (c *Suspect_Controller) Update(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", request.Date_Of_Birth)

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: &request.ID_Card_Number,
		Full_Name:      &request.Full_Name,
		Gender:         &request.Gender,
		Date_Of_Birth:  dateParse,
		Address:        &request.Address,
		Phone:          &request.Phone,
		Occupation:     &request.Occupation,
		Alibi:          &request.Alibi,
		Status:         &request.Status,
	}

	suspect, errUpdate := c.service.Update(ID, ID_Case, input)

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
		"Message": "Success Update Suspect Data",
		"Data":    suspect,
	})

}

// @Summary Menghapus data tersangka dari kasus
// @Description Endpoint ini menghapus data tersangka berdasarkan ID pada kasus tertentu. Proses ini tidak dapat dibatalkan dan berdampak pada data investigasi terkait. Asumsi: ID kasus dan ID tersangka valid.
// @Tags Suspects
// @Produce json
// @Param id_case path string true "UUID kasus, wajib diisi"
// @Param id path string true "UUID tersangka, wajib diisi"
// @Success 200 {object} map[string]interface{} "Berhasil menghapus tersangka"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan sistem internal"
// @Router /api/suspect/delete/{id}/{id_case} [delete]
func (c *Suspect_Controller) Delete(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID, ID_Case)

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
		"Message": "Success Delete Suspect Data",
	})

}
