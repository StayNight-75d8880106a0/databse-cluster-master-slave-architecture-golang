package detective_registry

import (
	"databse-cluster-master-slave-architecture-golang/app/ai/vector_store"
	"databse-cluster-master-slave-architecture-golang/app/controller/detective_controller"
	"databse-cluster-master-slave-architecture-golang/app/repository/detective_repository"
	"databse-cluster-master-slave-architecture-golang/app/service/detective_service"
)

type Detective_Module struct {
	Detective_Controller *detective_controller.Detective_Controller
}

func Detective_Registry(vs *vector_store.AI_VectorStore) *Detective_Module {

	repository := detective_repository.NewDetectiveRepositoryRegistry(vs)

	service := detective_service.NewDetectiveServiceRegistry(repository)

	controller := detective_controller.NewDetectiveControllerRegistry(service)

	return &Detective_Module{
		Detective_Controller: controller,
	}

}
