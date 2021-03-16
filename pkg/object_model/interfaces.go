package object_model

type LinkManager interface {
	GetLinks(request GetLinksRequest) (GetLinksResult, error)
	AddLink(request AddLinkRequest) error
	UpdateLink(request UpdateLinkRequest) error
	DeleteLink(username string, url string) error
}

type UserManager interface {
	Register(user User) error
	Login(username string, authToken string) (session string, err error)
	Logout(username string, session string) error
}

type SocialGraphManager interface {
	Follow(followed string, follower string) error
	Unfollow(followed string, follower string) error

	GetFollowing(username string) (map[string]bool, error)
	GetFollowers(username string) (map[string]bool, error)
}

type LinkManagerEvent interface {
	OnLinkAdded(username string, link *Link)
	onLinkUpdated(username string, link *Link)
	onLinkDeleted(username string, url string)
}
