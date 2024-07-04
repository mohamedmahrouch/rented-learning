# Utilisez une image Alpine basée sur Go 1.22.4
FROM golang:1.22.4-alpine

# Définissez le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copiez les fichiers de dépendances et téléchargez-les
COPY go.mod go.sum ./
RUN go mod tidy

# Copiez tout le reste des fichiers de votre application
COPY . .

# Exécutez votre application avec go run main.go
CMD ["go", "run", "cmd/server/main.go"]
