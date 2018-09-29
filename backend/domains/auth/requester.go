package auth

import (
	"fmt"

	"github.com/avtal.chalmers.it/backend/domains/util"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func upsertUser(newUser user) {
	c, session := util.FetchCollection("users")
	defer session.Close()

	if doesUserExist(newUser) {
		updateUserTerms(newUser, c)
	} else {
		c.Insert(&newUser)
	}

}

func updateUserTerms(usr user, c *mgo.Collection) {
	mergedTerms := mergeTerms(getUser(usr.Uid).AcceptedTerms, usr.AcceptedTerms)

	updateQuery := bson.M{
		"$set": bson.M{
			"acceptedterms": mergedTerms,
		},
	}

	err := c.Update(bson.M{"uid": usr.Uid}, updateQuery)
	if err != nil {
		fmt.Println(err)
	}
}

func mergeTerms(curTerms, termsDiff terms) terms {
	for k, v := range termsDiff {
		curTerms[k] = v
	}

	return curTerms
}

func doesUserExist(usr user) bool {
	return getUser(usr.Uid).Uid != ""
}

func getUser(cid string) user {
	c, session := util.FetchCollection("users")
	defer session.Close()

	result := user{}
	if err := c.Find(bson.M{"uid": cid}).One(&result); err != nil {
		util.LogCorruptData(err)
	}

	return result
}
