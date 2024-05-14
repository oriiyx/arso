from locust import HttpUser, task

class MaxHealthUser(HttpUser):
    @task
    def max_strength(self):
        self.client.get("/api/max-strength?lat=45.939824718792245&lon=15.848949414784824")