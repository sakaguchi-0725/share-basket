package registry

import "backend/usecase"

type UseCase struct {
	SignUp                 usecase.SignUp
	ResendConfirmationCode usecase.ResendConfirmationCode
	VerifyToken            usecase.VerifyToken
	GetJoinedGroups        usecase.GetJoinedGroups
	CreateGroup            usecase.CreateGroup
	GetGroup               usecase.GetGroup
	GetGroupMembers        usecase.GetGroupMembers
	GetShoppingItems       usecase.GetShoppingItems
}

func NewUseCase(repo *Repository) UseCase {
	return UseCase{
		SignUp:                 usecase.NewSignUp(repo.Account, repo.User, repo.Transaction),
		ResendConfirmationCode: usecase.NewResendConfirmationCode(),
		VerifyToken:            usecase.NewVerifyToken(repo.Authenticator),
		GetJoinedGroups:        usecase.NewGetJoinedGroups(repo.Account, repo.GroupMember, repo.Group),
		CreateGroup:            usecase.NewCreateGroup(),
		GetGroup:               usecase.NewGetGroup(),
		GetGroupMembers:        usecase.NewGetGroupMembers(),
		GetShoppingItems:       usecase.NewGetShoppingItems(),
	}
}
