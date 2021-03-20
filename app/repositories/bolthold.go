package repositories

import (
	"fmt"
	"os"

	_core "github.com/kiketordera/advanced-performance/app/core"
	a "github.com/kiketordera/advanced-performance/app/core"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	"github.com/timshannon/bolthold"
	"gopkg.in/mgo.v2/bson"
)

// BoltRepository implements models.Repository with Bolt DataBse
type BoltRepository struct {
	DataBase *bolthold.Store
}

// InitDatabase creates the instance of the BoltHold Database
func InitBoltDatabase() _domain.IRepository {
	fmt.Print("Init database exe3cuted")
	db, err := bolthold.Open("../data/"+a.KNameProyect+".db", 0666, nil)
	if err != nil {
		os.MkdirAll("../data/", os.ModePerm)
		db, err = bolthold.Open("../data/"+a.KNameProyect+".db", 0666, nil)
		if err != nil {
			fmt.Print("An error happened creating the DataBase: ", err)
		}
		// We create the frirst admin to controll the system
		db.Upsert(a.KAdminUser.ID, a.KAdminUser)
	}
	return BoltRepository{
		DataBase: db,
	}
}

// GetAdminUser finds and returns the Admin user
func (r BoltRepository) GetAdminUser() _domain.User {
	u, _ := r.GetUserByMail(_core.KAdminUser.Email)
	return u
}

// SaveObject saves the element in the DataBase
func (r BoltRepository) SaveObject(object interface{}, id bson.ObjectId) error {
	err := r.DataBase.Upsert(id, object)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// getUserByUsername finds the User in database by his username and returns it.
func (r BoltRepository) GetUserByMail(email string) (_domain.User, bool) {
	var users []_domain.User
	err := r.DataBase.Find(&users, bolthold.Where("Email").Eq(email))
	if err != nil || users == nil {
		if err != nil {
			fmt.Print(err)
		}
		return _domain.User{}, false
	}
	return users[0], true
}
