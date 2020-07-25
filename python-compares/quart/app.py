from quart import Quart, websocket

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

app = Quart(__name__)

@app.route('/')
async def hello():
    return 'hello'

@app.route('/users/<id>')
async def getUser(id):
    for u in users:
        if u.id == int(id):
            return f"the User is {u.__dict__}"
    return 'There is no User!'

@app.websocket('/ws')
async def ws():
    while True:
        await websocket.send('hello')




app.run(host="0.0.0.0")
