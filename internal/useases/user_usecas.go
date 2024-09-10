package useases

type IUserUsecase interface {
	GetSalary()
}

type userUsecase struct {
}

func UserUsecase() IUserUsecase {
	return &userUsecase{}
}

func (h *userUsecase) GetSalary() {

}
