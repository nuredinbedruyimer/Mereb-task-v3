package services

import (
	"Mereb3/constants"
	"Mereb3/database"
	"Mereb3/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, _ = database.CreateMongoClient()
var PersonCollecion = database.OpenCollection(client, "person")

func CreatePersonService(person *models.Person) error {

	person.PersonID = primitive.NewObjectID().Hex()
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()

	ctx, cancell := context.WithTimeout(context.Background(), constants.TIME_OUT)

	defer cancell()

	//  If We Just Need Insertion Add WE can use it instead of ignoring

	_, insertionError := PersonCollecion.InsertOne(ctx, person)

	if insertionError != nil {
		return insertionError

	}
	return nil

}

func GetAllPersonsService(perPage int, page int) ([]models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()
	opt := options.Find().SetLimit(int64(perPage)).SetSkip(int64(page))

	cursor, err := PersonCollecion.Find(ctx, bson.M{}, opt)
	if err != nil {
		return nil, err
	}

	var persons []models.Person
	if err = cursor.All(ctx, &persons); err != nil {
		return nil, err
	}
	return persons, nil
}
func GetPersonService(id string) (models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()
	filter := bson.M{"person_id": id}

	var person models.Person
	err := PersonCollecion.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		return person, err
	}
	return person, nil
}

func UpdatePersonService(id string, updatedPerson models.Person) (*models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":       updatedPerson.Name,
			"age":        updatedPerson.Age,
			"updated_at": time.Now(), // Set updated_at to current time
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	filter := bson.M{"person_id": id}
	var newPerson models.Person
	if err := PersonCollecion.FindOneAndUpdate(ctx, filter, update, opts).Decode(&newPerson); err != nil {
		return nil, err
	}
	return &newPerson, nil
}

func DeletePersonService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()
	filter := bson.M{"person_id": id}

	_, err := PersonCollecion.DeleteOne(ctx, filter)
	return err
}
