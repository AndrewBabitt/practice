from flask import Flask

users = []

class Users(object):
    def __init__(self, id):
        self.id = id
        self.username = f"TestUser--{id}"
        self.email = f"t{id}@mail.org"

    def __repr__(self):
        return f"id = {self.id}, username = {self.username}, email = {self.email}"

def create_users():
    for i in range(100):
        newUser = Users(i)
        print(newUser)
        users.append(newUser)

create_users()

app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello, World!"

@app.route('/users/<id>')
def getUser(id):
    for u in users:
        if u.id == int(id):
            return f"the User is {u.__dict__}"
    return 'There is no User!'


if __name__ == "__main__":
    app.run()
