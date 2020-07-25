import tornado.ioloop
import tornado.web

users = []

class Users(object):
    def __init__(self, id):
        self.id = id
        self.username = f"TestUser--{id}"
        self.email = f"t{id}@mail.org"

    def __repr__(self):
        return f"id = {self.id}, username = {self.username}, email = {self.email}"

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")

class UserHandler(tornado.web.RequestHandler):
    def get(self, id):
        for u in users:
            if u.id == int(id):
                return self.write(f"the user is {u.__dict__}")

        self.write("There is no user!")

def make_app():
    return tornado.web.Application([
        (r"/", MainHandler),
        (r"/users/(.*)", UserHandler)
    ])

def create_users():
    for i in range(100):
        newUser = Users(i)
        print(newUser)
        users.append(newUser)

if __name__ == "__main__":
    create_users()
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()