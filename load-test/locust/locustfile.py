from locust import HttpUser, between, task
from datetime import datetime
import json
class WebsiteTest(HttpUser):
    wait_time = between(1,5)
    counter= 0

    @task
    def create_user(self):
        d = datetime.now()
        post_body = dict(ID=0, Username=f"username--{datetime.now}--{self.counter}", Email=f"ama-{self.counter}@email.com")
        print(json.dumps(post_body))
        self.counter+=1
        self.client.post("/users",data=json.dumps(post_body))
    
    @task
    def get_user(self):
        self.client.get("/users")
