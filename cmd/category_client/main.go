package main

import (
	"context"
	"fmt"
	desc "github.com/ilyasoloviev99/category-service/pkg/category"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := desc.NewCategoryServiceClient(connection)

	res, err := client.GetCategory(context.Background(), &desc.GetCategoryRequest{Name: "apple"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res.GetName())
}
