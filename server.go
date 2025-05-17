package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	mentalpb "github.com/ramyasreetejo/mental_wellness_bot/proto"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

type server struct {
	mentalpb.UnimplementedMentalWellnessBotServer
	model *genai.GenerativeModel
}

func testGenAIModel(model *genai.GenerativeModel) error {
	resp, err := model.GenerateContent(context.Background(), genai.Text("Test"))
	if err != nil {
		return fmt.Errorf("Gemini test failed: %v", err)
	}
	log.Println("Gemini test successful:", resp)
	return nil
}

func NewServer(apiKey string) (*server, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("genai client error: %v", err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	return &server{model: model}, nil
}

func (s *server) GenerateResponse(ctx context.Context, req *mentalpb.UserInput) (*mentalpb.BotResponse, error) {
	resp, err := s.model.GenerateContent(ctx, genai.Text(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	reply := ""
	for _, cand := range resp.Candidates {
		if len(cand.Content.Parts) > 0 {
			if part, ok := cand.Content.Parts[0].(genai.Text); ok {
				reply = string(part)
			}
		}
	}
	return &mentalpb.BotResponse{Reply: reply}, nil
}

func startRestAPI(srv *server) {
	http.HandleFunc("/api/chat", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Message string `json:"message"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		res, err := srv.GenerateResponse(r.Context(), &mentalpb.UserInput{Message: req.Message})
		if err != nil {
			http.Error(w, "AI error: "+err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"reply": res.Reply})
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))
	go func() {
		log.Println("REST server at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	srv, err := NewServer(apiKey)
	if err != nil {
		log.Fatalf("server init error: %v", err)
	}

	// Run test before starting servers
	if err := testGenAIModel(srv.model); err != nil {
		log.Fatalf("%v", err)
	}

	go startRestAPI(srv)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mentalpb.RegisterMentalWellnessBotServer(grpcServer, srv)

	log.Printf("gRPC server listening on :%s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
