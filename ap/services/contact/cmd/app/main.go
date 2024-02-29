package main

import (
	"ap/pkg/store/postgres"
	"ap/services/contact/internal/domain"
	"ap/services/contact/internal/repository"
	"context"
	"log"
)

func main() {
	log.Println("Starting the application...")
	conn, err := postgres.Connect("localhost", 5432, "postgres", "1112", "clean-arch-go")
	log.Println("Connected to the database...")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())
	contactRepo := repository.NewContactRepository(conn)
	contactRepo.CreateContact(domain.Contact{FirstName: "Aidana", MiddleName: "Gabduluahitovna", LastName: "Ibrayeva", PhoneNumber: "+77028518880"})

	contact, err := contactRepo.GetContact(1)
	log.Println("Retrieving contact...")
	if err != nil {
		log.Fatalf("Failed to retrieve contact: %v", err)
	}
	log.Printf("Retrieved contact: %v", contact)
}
