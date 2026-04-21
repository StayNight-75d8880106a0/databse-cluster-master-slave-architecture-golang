package suspect_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/ai/vector_store"
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"

	"gorm.io/gorm"
)

type Suspect_Repository struct {
	master *gorm.DB
	slave2 *gorm.DB
	vector *vector_store.AI_VectorStore
}

func NewSuspectRepositoryRegistry(vs *vector_store.AI_VectorStore) *Suspect_Repository {
	dbCluster := database.GetInstanceDbCluster()
	return &Suspect_Repository{
		master: dbCluster.Master,
		slave2: dbCluster.SlaveSuspects,
		vector: vs,
	}
}

func (repo *Suspect_Repository) Create(suspect *models.Suspects) error {

	errCreate := repo.master.Table("suspects").Create(suspect).Error

	go repo.vector.LoadDatabaseSnapshot(repo.master)

	return errCreate

}

func (repo *Suspect_Repository) GetAll(ID_Case string, limit int, offset int) ([]models.Suspects, int64, error) {

	var suspects []models.Suspects
	var totalData int64

	errCount := repo.slave2.Table("suspects").Where("case_id = ?", ID_Case).Count(&totalData).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	errGet := repo.slave2.Table("suspects").Where("case_id = ?", ID_Case).Limit(limit).Offset(offset).Find(&suspects).Error

	return suspects, totalData, errGet

}

func (repo *Suspect_Repository) GetById(ID string, ID_Case string) (*models.Suspects, error) {

	var suspect *models.Suspects

	errGet := repo.slave2.Table("suspects").Where("id = ? AND case_id = ?", ID, ID_Case).First(&suspect).Error

	return suspect, errGet

}

func (repo *Suspect_Repository) Update(ID string, ID_Case string, suspect *models.Suspects) error {

	errUpdate := repo.master.Table("suspects").Where("id = ? AND case_id = ?", ID, ID_Case).Updates(suspect).Error

	go repo.vector.LoadDatabaseSnapshot(repo.master)

	return errUpdate

}

func (repo *Suspect_Repository) Delete(ID string, ID_Case string) error {

	var suspect *models.Suspects

	errDelete := repo.master.Table("suspects").Unscoped().Where("id = ? AND case_id = ?", ID, ID_Case).Delete(&suspect).Error

	go repo.vector.LoadDatabaseSnapshot(repo.master)

	return errDelete

}
