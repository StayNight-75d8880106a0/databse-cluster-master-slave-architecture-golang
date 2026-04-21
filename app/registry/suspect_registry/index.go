package suspect_registry

import (
	"databse-cluster-master-slave-architecture-golang/app/ai/vector_store"
	"databse-cluster-master-slave-architecture-golang/app/controller/suspect_controller"
	"databse-cluster-master-slave-architecture-golang/app/repository/cases_repository"
	"databse-cluster-master-slave-architecture-golang/app/repository/suspect_repository"
	"databse-cluster-master-slave-architecture-golang/app/service/cases_service"
	"databse-cluster-master-slave-architecture-golang/app/service/suspect_service"
)

type Suspect_Module struct {
	Suspect_Controller *suspect_controller.Suspect_Controller
}

func Suspect_Registry(vs *vector_store.AI_VectorStore) *Suspect_Module {

	repository := suspect_repository.NewSuspectRepositoryRegistry(vs)

	cases_repository := cases_repository.NewCasesRepositoryRegistry(vs)

	cases_service := cases_service.NewCasesServiceRegistry(cases_repository)

	service := suspect_service.NewSuspectServiceRegistry(repository, cases_service)

	controller := suspect_controller.NewSuspectControllerRegistry(service)

	return &Suspect_Module{
		Suspect_Controller: controller,
	}

}
