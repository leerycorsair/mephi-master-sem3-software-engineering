package model

type ResourceStatus string

const (
	Free     ResourceStatus = "Free"
	Occupied ResourceStatus = "Occupied"
)

type Resource struct {
	status ResourceStatus
}

func NewResource(status ResourceStatus) Resource {
	return Resource{status: status}
}

func (r *Resource) IsFree() bool {
	return r.status == Free
}

func (r *Resource) Free() {
	r.status = Free
}

func (r *Resource) Occupy() error {
	if r.IsFree() {
		r.status = Occupied
		return nil
	} else {
		return ErrResourceIsBusy
	}
}
