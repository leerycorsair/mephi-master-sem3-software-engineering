package manager

type RequestStatus string

const (
	Waiting RequestStatus = "Waiting"
	Active  RequestStatus = "Active"
)

type Request struct {
	resourceId string
	time       int
	status     RequestStatus
}

func NewRequest(rId string, t int) *Request {
	return &Request{
		resourceId: rId,
		time:       t,
		status:     Waiting,
	}
}

func (r *Request) RequestProcess() {
	if r.status == Active {
		r.time--
	}
}

func (r Request) IsOver() bool {
	return r.time == 0
}

func (r Request) IsActive() bool {
	return r.status == Active
}

func (r Request) IsWaiting() bool {
	return r.status == Waiting
}

func (r *Request) MakeActive() {
	r.status = Active
}

func (r Request) GetResId() string {
	return r.resourceId
}
