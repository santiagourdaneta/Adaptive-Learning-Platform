package main

import (
    "log"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)


// Get a personalized learning path with caching
func getLearningPath(c *gin.Context) {
    username := c.Param("username")

    // Check cache first
    if val, ok := cache.Load(username); ok {
        log.Println("Serving from cache:", username)
        c.JSON(http.StatusOK, gin.H{"path": val})
        return
    }

   // Placeholder for the AI service. Returns IDs for now.
    // Make sure these IDs exist in your `topics` table.
	learningPathIDs := []int{1, 2, 3} 

    // Store in cache for a limited time (e.g., 5 minutes)
	cache.Store(username, learningPathIDs)
	go func() {
		time.Sleep(5 * time.Minute)
		cache.Delete(username)
	}()

	c.JSON(http.StatusOK, gin.H{"path": learningPathIDs})
}

// Register a new user with validation
func registerUserHandler(c *gin.Context) {
	var profile Profile
	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Server-side validation
	if err := validate.Struct(profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	stmt, err := db.Prepare("INSERT INTO profiles(username, learning_style) VALUES(?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(profile.Username, profile.LearningStyle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}


// Get content for a specific topic
// In go-api/handlers.go

func getContentByTopic(c *gin.Context) {
	topicID := c.Param("topicID")
	
    // Initialize contents as an empty slice to ensure it's always a valid JSON array.
    var contents []Content 

	rows, err := db.Query("SELECT id, topic_id, type, data FROM content WHERE topic_id = ?", topicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve content"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ID, &content.TopicID, &content.Type, &content.Data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan content"})
			return
		}
		contents = append(contents, content)
	}
    // Check if the database query returned an error
	if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during row iteration"})
		return
    }

	// The `contents` variable is now guaranteed to be a valid slice (even if empty)
    c.JSON(http.StatusOK, contents)
}

// Submit a user's answer and save it to history
func submitAnswer(c *gin.Context) {
	var history History
	if err := c.BindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("INSERT INTO history(user_id, content_id, is_correct, time_taken) VALUES(?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(history.UserID, history.ContentID, history.IsCorrect, history.TimeTaken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer submitted successfully"})
}

// Handler for search at the speed of light
func searchContent(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
        return
    }

    var contents []Content
    // Using a prepared statement for security and speed
    rows, err := db.Query("SELECT id, topic_id, type, data FROM content WHERE data LIKE ?", "%"+query+"%")
    if err != nil {
        log.Printf("Search query failed: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var content Content
        if err := rows.Scan(&content.ID, &content.TopicID, &content.Type, &content.Data); err != nil {
            log.Printf("Error scanning content row: %v", err)
            continue
        }
        contents = append(contents, content)
    }

    c.JSON(http.StatusOK, contents)
}