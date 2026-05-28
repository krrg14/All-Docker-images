from flask import Flask, render_template

app = Flask(__name__)

products = [
    {"name": "Organic Bananas", "category": "Fruit", "price": "$2.49", "description": "Sweet and ripe bunch"},
    {"name": "Whole Milk", "category": "Dairy", "price": "$4.19", "description": "Fresh local milk"},
    {"name": "Whole Wheat Bread", "category": "Bakery", "price": "$3.29", "description": "Soft loaf with grains"},
    {"name": "Spinach Pack", "category": "Vegetables", "price": "$2.89", "description": "Crisp baby spinach"},
    {"name": "Greek Yogurt", "category": "Dairy", "price": "$5.49", "description": "High-protein yogurt"},
    {"name": "Tomato Sauce", "category": "Pantry", "price": "$2.99", "description": "Rich tomato blend"},
]

@app.route('/')
def index():
    return render_template('index.html', products=products)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8501)
