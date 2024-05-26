from locust import HttpUser, task, between
import random

class WebsiteUser(HttpUser):
    wait_time = between(1, 5)  # Adjust the wait time as needed
    host = "http://localhost:3000"  # Set the base URL for the frontend

    @task
    def load_homepage(self):
        self.client.get("/")  # Simulates loading the frontend homepage

    @task
    def simulate_map_clicks(self):
        # Simulate 5 random map clicks
        for _ in range(5):
            coordinates = self.random_coordinates()
            url = f"http://localhost:4000/api/max-strength?lat={coordinates['lat']}&lon={coordinates['lon']}"
            self.client.get(url)  # Simulates the request sent to the backend server

    def random_coordinates(self):
        min_lat = 44.5
        max_lat = 47.0
        min_lon = 13.0
        max_lon = 17.0
        lat = random.uniform(min_lat, max_lat)
        lon = random.uniform(min_lon, max_lon)
        return {"lat": lat, "lon": lon}

# To simulate 500 users for 1 minute, use the following locust command:
# locust -f locustfile.py --headless -u 500 -r 500 -t 15s --host=http://localhost --csv=results
