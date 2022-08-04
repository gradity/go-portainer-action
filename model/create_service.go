package model

type CreateService struct {
	Name         string
	TaskTemplate TaskTemplate
}

type TaskTemplate struct {
	ContainerSpec ContainerSpec
}

type ContainerSpec struct {
	Image string
	Env   []string
	// Mounts
}
