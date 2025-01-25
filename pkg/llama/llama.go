package llama

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/ollama/ollama/api"
)

func llama() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	truepointer := true
	req := &api.GenerateRequest{
		Model: "llama3.2",
		Prompt: "Tu es une IA assitant trader spécialisé dans le day trading de momentum, tes réponses sont clair et courtes," +
			"Tu répondra en JSON uniquement (sans les backtick), un objet avec deux champ : un champ pour la note en number, et un autre champ en string pour ton analyse," +
			" l'action NVDA a un gros momentum aujourd'hui grace a un resultat trimestriel exceptionnel." +
			" note sur 5 le catalyseur fondamental de ce momentum.",

		// set streaming to true
		Stream: &truepointer,
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Créer un contexte avec annulation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Appelé à la fin pour s'assurer de tout nettoyer

	// Lancer un processus (exemple : "sleep 100")
	cmd := exec.CommandContext(ctx, "./ollama", "serve")

	// Démarrer le processus
	if err := cmd.Start(); err != nil {
		log.Fatalf("Erreur lors du démarrage du processus : %v", err)
	}
	log.Printf("Processus démarré avec PID %d", cmd.Process.Pid)

	// Capturer les signaux système (exemple : Ctrl+C ou arrêt du programme)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Attendre que le signal ou la fin du contexte soit déclenché
	go func() {
		<-signalChan
		log.Println("Signal reçu, arrêt du processus...")

		cmd := exec.CommandContext(ctx, "./ollama", "stop", "llama3.2:latest")
		if err := cmd.Run(); err != nil {
			log.Printf("Erreur lors de l'arrêt du processus : %v", err)
		}

		cancel() // Annuler le contexte pour arrêter le processus
	}()

	llama()
}
