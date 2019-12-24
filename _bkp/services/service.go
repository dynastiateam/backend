package services

//type Service interface {
//	user.Service
//}
//
//type services struct {
//	user.Service
//	//repo repository.Repository
//}
//
//func New(userSrv user.Service) Service {
//	return &services{userSrv}
//}

//const (
//	EventCreated   = iota //new event
//	EventAccepted         //accepted by guard
//	EventCompleted        //competed by guard
//	EventRejected         //rejected by guard
//	EventMissed           //wasn't completed before ETA
//)

//func (s *services) AddRequest(req *repository.Request) (*repository.Request, error) {
//	if err := validator.New().Struct(req); err != nil {
//		return nil, err
//	}
//
//	req.Status = EventCreated
//
//	return s.repo.AddRequest(req)
//}
