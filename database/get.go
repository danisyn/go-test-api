package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Filter struct {
	User string
	Password string
	Token string
}

type Download struct {
	Token string
}

type File struct {
	File string
}

func GetUser(user string, password string, token string) (bool){
	var status bool
	uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	usersCollection := client.Database("app").Collection("users")
	filter := Filter{User:user, Password: password, Token: token}
	var result Filter
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if result.User == "" {
		status = false
	} else {
		status = true
	}
	return  status
}

func GetConfig(token string, exists bool) (bool, string){
	fmt.Println("Called")
	var status bool
	var response string
	if exists == true {
		
		uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
				panic(err)
		}

		usersCollection := client.Database("app").Collection("kubeconfigs")
		filter := Download{Token: token}
		var result File

		err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
		if result.File == "" {
			status = false
			response = "There is no Kubeconfig file configured for this user"
		} else {
			status = true
			
			fmt.Println(result.File)
			response = result.File
			
		}
		
	}
	
	
return status, response


}