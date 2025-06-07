
import { useEffect, useState } from 'react';
import '../CSS/MyCart.css'
import NavBar from '../component/NavBar.jsx';
import Login from './Login.jsx';
import { useNavigate } from 'react-router-dom';
import jumkha_3 from '../assets/products/jumkha_3.jpg'
function MyCartProducts(){

function removeProduct()
{
  // const response=await fetch(`/glow-zaar/remove/{pid}${}`)
  console.log("hello")
  console.log(event)
}
const [cartItems,setCartItems]=useState([])

    useEffect(()=>{
        async function loadcartProducts(){
              try {
                   const response = await fetch("/glow-zaar/GetMyCardProducts", {
                                                  credentials: "include"
                                                      })

                   const data =await response.json()
                  if(data.Error)
                  {
                      console.log("Error occuresd = ",data.Error)
                  }else{
                    setCartItems(data)
                  }
                  
                  } catch (error) {
                      console.log("Error while fetching cart products")
                  }
          }
        loadcartProducts()
    }
    ,[cartItems])

  if(!cartItems)
  {
      return(
            <h2>Loading Products....</h2>
          );
      
  }
  else{
    if(cartItems.length==0){
      console.log("it is running")
      return (
        <div>
          <NavBar/>
          <h2>Your Cart is empty , Add more Products</h2>
        </div>
      )
    }
    return (
      <>
      <NavBar/>
      
        <div className="cart-wrapper">
            <br></br>
            <br></br>
          {/* <h2 className="cart-title">Proceed to Buy ({totalItems} items)</h2> */}
          <div className="cart-header">
            <h2 className="cart-title">   </h2>
            <button className="btn proceed-pay">Proceed to Buy  (totalItems items)</button>
          </div>
          {cartItems.map(product => (
            <div key={product.pid} className="cart-row-full">
              <img src={product.poster} alt={product.name} className="cart-img-left" />

              <div className="cart-details">
                <h3>{product.name}</h3>
                <p>No. of items: {product.quantity}</p>
                <p>Price: ₹{product.price}</p>
                <div className="cart-buttons">
                  <button className="btn delete" onClick={()=>{removeProduct()}}>Remove</button>
                  <button className="btn order">Place Order</button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </>
    );
  }
        
}

export default MyCartProducts