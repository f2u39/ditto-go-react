package handler

import (
	act_srv "ditto/service/act"
	game_srv "ditto/service/game"
	inc_srv "ditto/service/inc"
	user_srv "ditto/service/user"
)

var (
	ActService  act_srv.Service
	GameService game_srv.GameService
	IncService  inc_srv.IncService
	UserService user_srv.UserService
)

func Init() {
	ActService = act_srv.NewService()
	GameService = game_srv.NewService()
	IncService = inc_srv.NewIncService()
	UserService = user_srv.NewUserService()
}
