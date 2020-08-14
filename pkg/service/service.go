package service


type Service interface {
	Search()

}

type service struct {

}

func (s *service) Search() {
	search("text")
}

