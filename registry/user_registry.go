package registry

import (
	"clean_arc/interface/controller"
	ip "clean_arc/interface/presenter"
	ir "clean_arc/interface/repository"

	"clean_arc/usecase/interactor"
	up "clean_arc/usecase/presenter"
	ur "clean_arc/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(
		r.NewUserInteractor(),
		r.NewAuthInteractor(),
		r.NewTokenInteractor(),
	)
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

func (r *registry) NewTokenInteractor() interactor.TokenInteractor {
	return interactor.NewTokenInteractor(r.NewTokenRepository())
}

func (r *registry) NewTokenRepository() ur.TokenRepository {
	return ir.NewTokenRepository()
}

func (r *registry) NewAuthInteractor() interactor.AuthInteractor {
	return interactor.NewAuthInteractor(r.NewAuthRepository())
}

func (r *registry) NewAuthRepository() ur.AuthRepository {
	return ir.NewAuthRepository(r.redisDB)
}
