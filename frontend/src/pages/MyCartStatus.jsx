// MyCart.jsx

import { useEffect, useState } from 'react';
import '../CSS/MyCart.css';
import NavBar from '../component/NavBar.jsx';
import Login from './Login.jsx';
import { useNavigate } from 'react-router-dom';

export function MyCartStatus() {
  const [isLogin, setLoginIn] = useState(null);
  let product
  useEffect(() => {
    async function checkLogin() {
      try {
        const response = await fetch("/glow-zaar/my-profile-status", {
          credentials: "include"
        });

        const data = await response.json();
        setLoginIn(data);
      } catch (error) {
        console.log("Error while fetching Cookie Found details");
      }
    }
    checkLogin();
  }, []);

  if (!isLogin) {
    return(
      <div>
        <NavBar/>
        <h2>Loading login status....</h2>;
      </div>
    ) 
  }

  if (!isLogin.isFound) {
    return (
      <div>
        <NavBar />
        <Login />
      </div>
    );
  }

  return <MyCartProducts />;
}

export function MyCartProducts() {
  const [cartItems, setCartItems] = useState([]);
  const [removalStatus, setRemovalStatus] = useState(null);

  async function removeProduct(pid) {
    try {
      const response = await fetch(`/glow-zaar/remove/${pid}`, {
        method: "DELETE"
      });
      const data = await response.json();
      setRemovalStatus(data);
      setCartItems(prevItems => prevItems.filter(item => item.pid !== pid));
    } catch (error) {
      console.log("Error while fetching product removal");
    }
  }
  async function placeOrder(product){
    try {
      const response = await fetch("/glow-zaar/place-order", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'  
          },
        body : JSON.stringify({"pid":product.pid,"price":product.price,"quantity":product.quantity})
      });
      const data = await response.json();
      console.log("data - ",data)
      setRemovalStatus(data);
      setCartItems(prevItems => prevItems.filter(item => item.pid !== product.pid));
    } catch (error) {
      console.log("Error while fetching place order")
    }
  }

  async function orderAllCartItems(){
    try {
      const response = await fetch("/glow-zaar/order-all-cartProducts",{
        credentials : 'include'
      })
      const data=await response.json()
      console.log("data from all the cart products - ",data)
      setRemovalStatus(data);
      setCartItems([]);
    } catch (error) {
      console.log("error occured while fetching the order all cart products")
    }
  }

  useEffect(() => {
    async function loadCartProducts() {
      try {
        const response = await fetch("/glow-zaar/GetMyCardProducts", {
          credentials: "include"
        });
        const data = await response.json();
        if (data.Error) {
          console.log("Error occurred = ", data.Error);
        } else {
          setCartItems(data);
        }
      } catch (error) {
        console.log("Error while fetching cart products");
      }
    }

    loadCartProducts();
  }, []);

  if (!cartItems) {
    return <h2>Loading Products....</h2>;
  }

  if (cartItems.length === 0) {
    return (
      <div>
        <NavBar />
        <h2>Your Cart is empty, add more products</h2>
      </div>
    );
  }

  const totalItems = cartItems.length;

  return (
    <>
      <NavBar />
      <div className="cart-wrapper">
        <div className="cart-header">
          <h2 className="cart-title"> </h2>
          <button className="btn proceed-pay" onClick={()=>{orderAllCartItems()}}>Proceed to Buy ({totalItems} items)</button>
        </div>
        {cartItems.map(product => (
          <div key={product.pid} className="cart-row-full">
            <img src={product.poster} alt={product.name} className="cart-img-left" />
            <div className="cart-details">
              <h3>{product.name}</h3>
              <p>No. of items: {product.quantity}</p>
              <p>Price: ₹{product.price}</p>
              <div className="cart-buttons">
                <button className="btn delete" onClick={() => removeProduct(product.pid)}>Remove</button>
                <button className="btn order" onClick={()=>{placeOrder(product)}}>Place Order</button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
}
