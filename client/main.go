package main

import (
	"context"
	"log"
	pb "root/proto"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HelloWorldRequest struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/", func(c *fiber.Ctx) error {
		var req HelloWorldRequest
		if err := c.BodyParser(&req); err != nil {
			log.Fatalf("не удалось распарсить запрос: %v", err)
		}

		conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatalf("не удалось подключиться: %v", err)
		}
		defer conn.Close()
		client := pb.NewGetHelloClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.HelloWorld(ctx, &pb.HelloWorldResponse{
			Message: req.Message,
		})
		if err != nil {
			log.Fatalf("не удалось обработать запрос: %v", err)
		}

		return c.SendString(r.GetMessage())
	})

	log.Fatal(app.Listen(":3000"))
}
