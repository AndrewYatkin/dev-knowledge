package userRepoModel

import profileEntity "dev-knowledge/domain/entity/user/profile"

type Profile struct {
	ID         string `bson:"profile_id"`
	FirstName  string `bson:"first_name"`
	LastName   string `bson:"last_name"`
	Patronymic string `bson:"patronymic"`
}

func ProfileToEntity(repoProfile *Profile) (*profileEntity.Profile, error) {
	profileID, err := profileEntity.ProfileIDFrom(repoProfile.ID)
	if err != nil {
		return nil, err
	}

	profileBuilder := profileEntity.NewBuilder().
		ID(profileID).
		LastName(repoProfile.LastName).
		Patronymic(repoProfile.Patronymic)

	if repoProfile.FirstName != "" {
		profileBuilder.FirstName(repoProfile.FirstName)
	}
	return profileBuilder.Build(), nil
}

func ProfileToModel(profile *profileEntity.Profile) *Profile {
	return &Profile{
		ID:         profile.ID().String(),
		FirstName:  profile.FirstName(),
		LastName:   profile.LastName(),
		Patronymic: profile.Patronymic(),
	}
}
