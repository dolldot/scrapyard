from flask import Flask, request, render_template
from rag import rag

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/query', methods=['POST'])
def query():
    user_query = request.form['query']
    response = rag(user_query)
    return response

if __name__ == "__main__":
    app.run(debug=True)