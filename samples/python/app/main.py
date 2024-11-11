from flask import *

app = Flask(__name__)
app.json.enxure_ascii = False

@app.get("/")
def message():
  return jsonify({"message": "hello python"})

@app.get("/<name>")
def greeting(name):
  return jsonify({"message": "hello! " + name})

if __name__ == '__main__':
  app.run(host='0.0.0.0', port=9000, debug=True)