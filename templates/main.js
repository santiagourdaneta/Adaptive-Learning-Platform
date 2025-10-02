document.addEventListener('DOMContentLoaded', () => {
    const contentContainer = document.getElementById('content-container');
    const searchForm = document.getElementById('search-form');
    const username = 'learner123'; // Este valor debe ser dinámico en una aplicación real.

    // Fetch the personalized learning path from the Go API.
    async function fetchLearningPath() {
        try {
            const response = await fetch(`http://localhost:8080/path/${username}`);
        if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
        const data = await response.json();
        
        // Log the IDs returned by the API
        console.log('Learning Path IDs:', data.path); 

        if (data.path && data.path.length > 0) {
            // Pass the first ID (integer) to the next function
            fetchContentByTopic(data.path[0]); 
        } else {
            fetchContentByTopic(data.path[0]); // Pass the first ID
        }
        } catch (error) {
            console.error('Failed to fetch learning path:', error);
            contentContainer.innerHTML = '<p>Error loading learning path. The server might be down or unreachable.</p>';
        }
    }

    // Fetch content for a specific topic.
    async function fetchContentByTopic(topicID) {
        try {
            const response = await fetch(`http://localhost:8080/content/${topicID}`);
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const contents = await response.json();
            contentContainer.innerHTML = ''; // Clear previous content.
            
            if (contents.length > 0) {
                contents.forEach(content => {
                    const element = document.createElement('div');
                    element.innerHTML = `<p><strong>Topic ID: ${content.topic_id}</strong> - ${content.data}</p>`;
                    contentContainer.appendChild(element);
                });
            } else {
                contentContainer.innerHTML = '<p>No content available for this topic.</p>';
            }
        } catch (error) {
            console.error('Failed to fetch content:', error);
            contentContainer.innerHTML = '<p>Error loading content. Check if the Go API is running.</p>';
        }
    }

    // Search for content based on a user query.
    async function search(query) {
        try {
            const response = await fetch(`http://localhost:8080/search?q=${encodeURIComponent(query)}`);
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const results = await response.json();
            contentContainer.innerHTML = ''; // Clear previous content.

            if (results.length > 0) {
                results.forEach(result => {
                    const element = document.createElement('div');
                    element.innerHTML = `<p><strong>Search Result:</strong> ${result.data}</p>`;
                    contentContainer.appendChild(element);
                });
            } else {
                contentContainer.innerHTML = '<p>No search results found.</p>';
            }
        } catch (error) {
            console.error('Search failed:', error);
            contentContainer.innerHTML = '<p>Error performing search. Please check your network connection.</p>';
        }
    }

    // Add event listener for the search form.
    searchForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const query = e.target.elements.query.value;
        if (query.trim() !== '') {
            search(query);
        }
    });

    // Initial call to fetch the learning path when the page loads.
    fetchLearningPath();
});