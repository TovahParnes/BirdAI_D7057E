package repositories

type RepositoryEndpoints struct {
	User  UserRepository
	Post  PostRepository
	Bird  BirdRepository
	Media MediaRepository
	Admin AdminRepository
}
