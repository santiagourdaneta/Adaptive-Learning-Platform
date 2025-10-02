package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database and creates the necessary tables.
func InitDB(databasePath string) {
	var err error
	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}

	createTablesSQL := `
	CREATE TABLE IF NOT EXISTS profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		learning_style TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS topics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);
	CREATE TABLE IF NOT EXISTS content (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		topic_id INTEGER NOT NULL,
		type TEXT NOT NULL,
		data TEXT NOT NULL,
		FOREIGN KEY(topic_id) REFERENCES topics(id)
	);
	CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		content_id INTEGER NOT NULL,
		is_correct BOOLEAN NOT NULL,
		time_taken INTEGER NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES profiles(id),
		FOREIGN KEY(content_id) REFERENCES content(id)
	);
	`
	_, err = db.Exec(createTablesSQL)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// Add indexes for lightning-fast search
	createIndexesSQL := `
	CREATE INDEX IF NOT EXISTS idx_content_topic_id ON content(topic_id);
	CREATE INDEX IF NOT EXISTS idx_content_data ON content(data);
	CREATE INDEX IF NOT EXISTS idx_history_user_id ON history(user_id);
	`
	_, err = db.Exec(createIndexesSQL)
	if err != nil {
		log.Fatalf("Error creating indexes: %v", err)
	}

	seedDatabase()
}

// Structs to represent our data
type Profile struct {
	ID           int    `json:"id"`
	Username     string `json:"username" validate:"required,min=3,max=30"`
	LearningStyle string `json:"learning_style" validate:"required"`
}

type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Content struct {
	ID      int    `json:"id"`
	TopicID int    `json:"topic_id"`
	Type    string `json:"type"`
	Data    string `json:"data"`
}

type History struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ContentID int    `json:"content_id"`
	IsCorrect bool   `json:"is_correct"`
	TimeTaken int    `json:"time_taken"`
	Timestamp string `json:"timestamp"`
}

func seedDatabase() {
	log.Println("Seeding database with initial data...")

	// 1. Poblar la tabla 'topics'
	// Poblar la tabla 'topics'
    topics := []string{"Go", "Python", "Rust"}
    for _, topic := range topics {
        _, err := db.Exec("INSERT OR IGNORE INTO topics (name) VALUES (?)", topic)
		if err != nil {
			log.Printf("Error seeding topic %s: %v", topic, err)
		}
	}

	// 2. Poblar la tabla 'profiles'
	profiles := []string{"user1", "user2", "user3"}
	learningStyles := []string{"visual", "auditory", "kinesthetic"}
	for i, username := range profiles {
		_, err := db.Exec("INSERT OR IGNORE INTO profiles (username, learning_style) VALUES (?, ?)", username, learningStyles[i%len(learningStyles)])
		if err != nil {
			log.Printf("Error seeding user %s: %v", username, err)
		}
	}
	
	// 3. Poblar la tabla 'content' (depende de topics)
	// Primero obtenemos los IDs de los topics
	var topicIDs map[string]int64 = make(map[string]int64)
	rows, err := db.Query("SELECT id, name FROM topics")
	if err != nil {
		log.Printf("Error getting topic IDs: %v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Println("Error scanning topic row:", err)
			continue
		}
		topicIDs[name] = id
	}
	
	// Ahora insertamos el contenido usando los IDs de los topics
	contentData := []struct {
		TopicName string
		Type      string
		Data      string
	}{
		{"Go", "text", "Go is a statically typed, compiled programming language..."},
		{"Go", "quiz", "What is Go's mascot?"},
		{"Python", "text", "Python is an interpreted, high-level, general-purpose programming language."},
		{"Python", "quiz", "Which keyword is used for a function in Python?"},
		{"Rust", "text", "Rust is a multi-paradigm, general-purpose programming language."},
		{"Rust", "quiz", "What does 'ownership' mean in Rust?"},
	}
	for _, c := range contentData {
		topicID, ok := topicIDs[c.TopicName]
		if !ok {
			log.Printf("Topic ID for '%s' not found.", c.TopicName)
			continue
		}
		_, err := db.Exec("INSERT OR IGNORE INTO content (topic_id, type, data) VALUES (?, ?, ?)", topicID, c.Type, c.Data)
		if err != nil {
			log.Printf("Error seeding content: %v", err)
		}
	}
	
	log.Println("Database seeding complete.")
}

// GetTopicIDByName fetches a topic's ID from the database using its name.
func GetTopicIDByName(name string) (int, error) {
    var id int
    err := db.QueryRow("SELECT id FROM topics WHERE name = ?", name).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}