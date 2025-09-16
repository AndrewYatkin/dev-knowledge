package userRepoModel

import passwordEntity "dev-knowledge/domain/entity/user/password"

type Password struct {
	Hash string `bson:"hash"`
	Salt string `bson:"salt"`
}

func PasswordToEntity(password *Password) (*passwordEntity.Password, error) {
	entityHash, err := passwordEntity.HashFrom(password.Hash)
	if err != nil {
		return nil, err
	}

	entitySalt, err := passwordEntity.SaltFrom(password.Salt)
	if err != nil {
		return nil, err
	}

	entityPassword, err := passwordEntity.NewBuilder().
		Hash(entityHash).
		Salt(entitySalt).
		Build()
	if err != nil {
		return nil, err
	}

	return entityPassword, nil
}

func PasswordToModel(password *passwordEntity.Password) *Password {
	return &Password{
		Hash: password.Hash(),
		Salt: password.Salt(),
	}
}
