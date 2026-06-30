const API_BASE = "http://localhost:8080/asteroids";
const gridElement = document.getElementById("asteroidGrid");
const formElement = document.getElementById("asteroidForm");

// 1. READ: Fetch all tracking entries from the Go server and draw them to screen
async function loadAsteroids() {
    try {
        const response = await fetch(`${API_BASE}?limit=50`);
        if (!response.ok) throw new Error("Network query failed");
        
        const data = await response.json();
        
        // Wipe old data cards clear
        gridElement.innerHTML = "";

        if (!data || data.length === 0) {
            gridElement.innerHTML = `<p style="color: var(--text-muted)">No tracking assets found in PostgreSQL memory.</p>`;
            return;
        }

        // Loop through each item in the JSON payload array
        data.forEach(asteroid => {
            const isHazard = asteroid.is_potentially_hazardous_asteroid;
            const cardHtml = `
                <div class="card">
                    <span class="badge ${isHazard ? 'hazard' : 'safe'}">
                        ${isHazard ? '⚠️ Hazardous' : '🟢 Safe Orbit'}
                    </span>
                    <div class="card-title">${asteroid.name}</div>
                    <div class="card-detail"><strong>Database ID:</strong> ${asteroid.id}</div>
                    <div class="card-detail"><strong>Magnitude (H):</strong> ${asteroid.absolute_magnitude_h}</div>
                    <button class="delete-btn" onclick="deleteAsteroid('${asteroid.id}')">Drop Reference</button>
                </div>
            `;
            gridElement.insertAdjacentHTML("beforeend", cardHtml);
        });
    } catch (error) {
        console.error("Fetch Error:", error);
        gridElement.innerHTML = `<p style="color: var(--accent-hazard)">Failed loading data. Ensure Go server is active on port 8080.</p>`;
    }
}

// 2. CREATE: Take form details and POST it as a JSON block
formElement.addEventListener("submit", async (e) => {
    e.preventDefault();

    const payload = {
        id: document.getElementById("astId").value,
        name: document.getElementById("astName").value,
        absolute_magnitude_h: parseFloat(document.getElementById("astMag").value),
        is_potentially_hazardous_asteroid: document.getElementById("astHazard").value === "true"
    };

    try {
        const response = await fetch(`${API_BASE}/create`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload)
        });

        if (response.ok) {
            formElement.reset(); // clear input rows
            loadAsteroids();    // instantly re-draw the inventory list out of the database!
        } else {
            alert("Database rejected entry creation.");
        }
    } catch (error) {
        console.error("Submission error:", error);
    }
});

// 3. DELETE: Target a specific asset identifier string and call the removal route
async function deleteAsteroid(id) {
    if (!confirm(`Are you sure you want to delete asteroid ID: ${id}?`)) return;

    try {
        const response = await fetch(`${API_BASE}/delete?id=${id}`, {
            method: "DELETE"
        });

        if (response.ok) {
            loadAsteroids(); // Refresh grid instantly
        } else {
            alert("Failed to drop record from target index.");
        }
    } catch (error) {
        console.error("Deletion error:", error);
    }
}

// Kick off the interface rendering on initial engine startup
loadAsteroids();