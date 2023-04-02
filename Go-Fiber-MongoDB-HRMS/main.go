package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mongoInstance MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017"

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	Salary float64 `json:"salary,omitempty" bson:"salary,omitempty"`
	Age    int     `json:"age,omitempty" bson:"age,omitempty"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mongoInstance = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		var employees []Employee
		cursor, err := mongoInstance.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer cursor.Close(c.Context())

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		return c.JSON(employees)
	})
	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mongoInstance.Db.Collection("employees")
		var employee Employee
		if err := c.BodyParser(&employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}

		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)

	})
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		collection := mongoInstance.Db.Collection("employees")

		idParam := c.Params("id")
		employeeID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employee Employee

		if err := c.BodyParser(&employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeID}}

		updateQuery := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "name",
						Value: employee.Name,
					},
					{
						Key:   "salary",
						Value: employee.Salary,
					},
					{
						Key:   "age",
						Value: employee.Age,
					},
				},
			},
		}

		err = collection.FindOneAndUpdate(c.Context(), query, updateQuery).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)

			}
			return c.SendStatus(500)

		}

		employee.ID = idParam

		return c.Status(200).JSON(employee)

	})
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		collection := mongoInstance.Db.Collection("employees")

		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			return c.Status(500).SendString(err.Error())

		}

		query := bson.D{{Key: "_id", Value: employeeID}}

		_, err = collection.DeleteOne(c.Context(), query)

		if err != nil {
			return c.Status(500).SendString(err.Error())

		}

		return c.Status(200).JSON("employee deleted")
	})

	log.Fatal(app.Listen(":8000"))

}
