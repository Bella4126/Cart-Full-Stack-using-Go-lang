import { useState, useEffect } from 'react'
import './App.css'

const API_BASE = 'https://cart-full-stack-using-go-lang.onrender.com'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [token, setToken] = useState('')
  const [user, setUser] = useState(null)
  const [items, setItems] = useState([
    { ID: 1, name: 'Laptop', price: 999.99 },
    { ID: 2, name: 'Wireless Mouse', price: 29.99 },
    { ID: 3, name: 'Mechanical Keyboard', price: 79.99 },
    { ID: 4, name: 'USB-C Hub', price: 49.99 },
    { ID: 5, name: 'Webcam HD', price: 89.99 },
    { ID: 6, name: 'Bluetooth Headphones', price: 159.99 }
  ])
  const [cartItems, setCartItems] = useState([])
  const [orders, setOrders] = useState([])

  useEffect(() => {
    const savedToken = localStorage.getItem('token')
    const savedUser = localStorage.getItem('user')
    if (savedToken && savedUser) {
      setToken(savedToken)
      setUser(JSON.parse(savedUser))
      setIsLoggedIn(true)
    }
  }, [])

  useEffect(() => {
    if (isLoggedIn) {
      // fetchItems() // Commented out to use dummy data
    }
  }, [isLoggedIn])

  const fetchItems = async () => {
    try {
      const response = await fetch(`${API_BASE}/items`)
      const data = await response.json()
      setItems(data)
    } catch (error) {
      console.error('Error fetching items:', error)
    }
  }

  const fetchCartItems = async () => {
    try {
      const response = await fetch(`${API_BASE}/carts`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      const data = await response.json()
      const userCartItems = data.filter(cart => cart.user_id === user.id)
      setCartItems(userCartItems)
      return userCartItems
    } catch (error) {
      console.error('Error fetching cart items:', error)
      return []
    }
  }

  const fetchOrders = async () => {
    try {
      const response = await fetch(`${API_BASE}/orders`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      const data = await response.json()
      const userOrders = data.filter(order => order.user_id === user.id)
      setOrders(userOrders)
      return userOrders
    } catch (error) {
      console.error('Error fetching orders:', error)
      return []
    }
  }

  const handleLogin = async (username, password) => {
    try {
      const response = await fetch(`${API_BASE}/users/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      })

      if (response.ok) {
        const data = await response.json()
        setToken(data.token)
        setUser(data.user)
        setIsLoggedIn(true)
        localStorage.setItem('token', data.token)
        localStorage.setItem('user', JSON.stringify(data.user))
      } else {
        const error = await response.json()
        alert(`Login failed: ${error.error}`)
      }
    } catch (error) {
      alert('Login failed: ' + error.message)
    }
  }

  const handleRegister = async (username, password) => {
    try {
      const response = await fetch(`${API_BASE}/users`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      })

      if (response.ok) {
        alert('Registration successful! Please login.')
      } else {
        const error = await response.json()
        alert(`Registration failed: ${error.error}`)
      }
    } catch (error) {
      alert('Registration failed: ' + error.message)
    }
  }

  const handleAddToCart = async (itemId) => {
    const item = items.find(item => item.ID === itemId)
    const existingCartItem = cartItems.find(cartItem => cartItem.item_id === itemId)
    
    if (existingCartItem) {
      setCartItems(cartItems.map(cartItem => 
        cartItem.item_id === itemId 
          ? { ...cartItem, quantity: cartItem.quantity + 1 }
          : cartItem
      ))
    } else {
      setCartItems([...cartItems, { 
        item_id: itemId, 
        quantity: 1, 
        item: item 
      }])
    }
    
    alert('Item added to cart!')
  }

  const handleShowCart = async () => {
    if (cartItems.length === 0) {
      alert('Your cart is empty!')
    } else {
      const cartInfo = cartItems.map(item => 
        `${item.item.name} - $${item.item.price} x ${item.quantity}`
      ).join('\n')
      alert(`Cart Items:\n${cartInfo}`)
    }
  }

  const handleShowOrders = async () => {
    if (orders.length === 0) {
      alert('No orders found!')
    } else {
      const orderInfo = orders.map(order => 
        `Order #${order.ID} - Total: $${order.total}`
      ).join('\n')
      alert(`Order History:\n${orderInfo}`)
    }
  }

  const handleCheckout = async () => {
    if (cartItems.length === 0) {
      alert('Your cart is empty!')
      return
    }
    
    const total = cartItems.reduce((sum, cartItem) => 
      sum + (cartItem.item.price * cartItem.quantity), 0
    )
    
    const newOrder = {
      ID: orders.length + 1,
      total: total.toFixed(2),
      items: [...cartItems]
    }
    
    setOrders([...orders, newOrder])
    setCartItems([])
    alert('Order placed successfully!')
  }

  const handleLogout = () => {
    setIsLoggedIn(false)
    setToken('')
    setUser(null)
    setItems([])
    setCartItems([])
    setOrders([])
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  if (!isLoggedIn) {
    return <Login onLogin={handleLogin} onRegister={handleRegister} />
  }

  return (
    <div className="app">
      <header className="header">
        <h1>Shopping Cart</h1>
        <div className="user-info">
          <span>Welcome, {user?.username}!</span>
          <button onClick={handleLogout} className="logout-btn">Logout</button>
        </div>
      </header>

      <div className="controls">
        <button onClick={handleShowCart} className="cart-btn">Cart</button>
        <button onClick={handleShowOrders} className="orders-btn">Order History</button>
        <button onClick={handleCheckout} className="checkout-btn">Checkout</button>
      </div>

      <div className="items-grid">
        {items.map(item => (
          <div key={item.ID} className="item-card">
            <h3>{item.name}</h3>
            <p className="price">${item.price}</p>
            <button 
              onClick={() => handleAddToCart(item.ID)}
              className="add-to-cart-btn"
            >
              Add to Cart
            </button>
          </div>
        ))}
      </div>
    </div>
  )
}

function Login({ onLogin, onRegister }) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [isRegister, setIsRegister] = useState(false)

  const handleSubmit = (e) => {
    e.preventDefault()
    if (!username || !password) {
      alert('Please fill in all fields')
      return
    }

    if (isRegister) {
      onRegister(username, password)
    } else {
      onLogin(username, password)
    }
  }

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit} className="login-form">
        <h2>{isRegister ? 'Register' : 'Login'}</h2>
        
        <div className="form-group">
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>

        <div className="form-group">
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        <button type="submit" className="submit-btn">
          {isRegister ? 'Register' : 'Login'}
        </button>

        <button 
          type="button" 
          onClick={() => setIsRegister(!isRegister)}
          className="toggle-btn"
        >
          {isRegister ? 'Already have an account? Login' : 'Need an account? Register'}
        </button>
      </form>
    </div>
  )
}

export default App