package v8

import (
	"github.com/v8platform/designer"
)

func NewRepository(path string, userAndPassword ...string) designer.Repository {

	repo := designer.Repository{
		Path: path,
	}

	if len(userAndPassword) > 0 && len(userAndPassword[0]) > 0 {
		repo.User = userAndPassword[0]
		if len(userAndPassword) == 2 && len(userAndPassword[1]) > 0 {
			repo.Password = userAndPassword[1]
		}
	}

	return repo
}

// RepositoryUpdateCfg получает команду обновления конфигурации из хранилища конфигурации
// Подробнее в пакете designer.UpdateCfgOptions
func RepositoryUpdateCfg(repo designer.Repository, updateDBCfg ...designer.UpdateDBCfgOptions) designer.RepositoryUpdateCfgOptions {

	command := designer.RepositoryUpdateCfgOptions{
		Repository: repo,
		Designer:   designer.NewDesigner(),
	}

	//if len(updateDBCfg) > 0 {
	//	command.UpdateDBCfg = &updateDBCfg[0]
	//}

	return command
}
