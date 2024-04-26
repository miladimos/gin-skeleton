package repository

// log "github.com/sirupsen/logrus"

// type UserRepository interface {
// 	CommonBehaviorRepository(models.User)

// 	FindAllUser() ([]dao.User, error)
// 	FindUserById(id int) (dao.User, error)
// 	Save(user *dao.User) (dao.User, error)
// 	DeleteUserById(id int) error
// }

// // type usersRepository struct {
// // 	CommonBehaviourRepository[entity.User]
// // }

// // func NewUsersRepository(db *gorm.DB) UserRepository {
// // 	return &usersRepository{
// // 		CommonBehaviourRepository: NewCommonBehaviour[entity.User](db),
// // 	}
// // }

// type UserRepositoryImpl struct {
// 	db *gorm.DB
// }

// func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
// 	var users []dao.User

// 	var err = u.db.Preload("Role").Find(&users).Error
// 	if err != nil {
// 		log.Error("Got an error finding all couples. Error: ", err)
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
// 	user := dao.User{
// 		ID: id,
// 	}
// 	err := u.db.Preload("Role").First(&user).Error
// 	if err != nil {
// 		log.Error("Got and error when find user by id. Error: ", err)
// 		return dao.User{}, err
// 	}
// 	return user, nil
// }

// func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
// 	var err = u.db.Save(user).Error
// 	if err != nil {
// 		log.Error("Got an error when save user. Error: ", err)
// 		return dao.User{}, err
// 	}
// 	return *user, nil
// }

// func (u UserRepositoryImpl) DeleteUserById(id int) error {
// 	err := u.db.Delete(&dao.User{}, id).Error
// 	if err != nil {
// 		log.Error("Got an error when delete user. Error: ", err)
// 		return err
// 	}
// 	return nil
// }

// func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
// 	db.AutoMigrate(&dao.User{})
// 	return &UserRepositoryImpl{
// 		db: db,
// 	}
// }
