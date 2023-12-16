from fastapi import FastAPI

app = FastAPI()

users = {
    1: {
        'username': 'user1',
        'email': 'user1@email.com',
        'password': 'password1'
    },
    2: {
        'username': 'user2',
        'email': 'user2@email.com',
        'password': 'password2'
    },
    3: {
        'username': 'user3',
        'email': 'user3@email.com',
        'password': 'password3'
    }
}

@app.get('/users')
def get_users():
    return users

@app.post('/users')
def create_user(username: str, email: str, password: str):
    users[len(users) + 1] = {
        'username': username,
        'email': email,
        'password': password
    }
    return True
    
@app.post('/login')
def login(email: str, password: str):
    for user in users:
        data = user[user]
        if data['email'] == email and data['password'] == password:
            return 'Login successfully'
        
    return 'Login failed'


if __name__ == '__main__':
    import uvicorn
    uvicorn.run(app)