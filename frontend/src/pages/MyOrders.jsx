
import { useEffect, useState } from 'react';
import '../CSS/MyCart.css'
import NavBar from '../component/NavBar.jsx';
import Login from './Login.jsx';
import { useNavigate } from 'react-router-dom';
import jumkha_3 from '../assets/products/jumkha_3.jpg'
import '../CSS/MyOrders.css';
export function MyOrderStatus(){

const [isLogin,setLoginIn]=useState(null)
const navigate=useNavigate()

useEffect(()=>{
    async function checkLogin(){
        try {
            const response = await fetch("/glow-zaar/my-profile-status", {
                                    credentials: "include"
                                        })

            const data =await response.json()
            setLoginIn(data)
        } catch (error) {
            console.log("Error while fetching Cookie Found details")
        }
    }
    checkLogin()
},[])

if(!isLogin){
    return(
        <h2>Loading login status....</h2>
    );
}
if(!isLogin.isFound){
        return(
            <div>
                <NavBar/>
                <Login/>
            </div>
        );
    }
    return <MyOrders/>
}

export function MyOrders() {

  const [productItems,setProductItems]=useState([])
  const [profile,setProfile]=useState(null)
  useEffect(()=>{
      async function loadorderedProducts(){
                    try {
                        const response = await fetch("/glow-zaar/GetMyOrderProducts", {
                                                        credentials: "include"
                                                            })
  
                        const data =await response.json()
                        if(data.Error)
                        {
                            console.log("Error occured = ",data.Error)
                        }else{
                          setProductItems(data)
                        }
                        
                        } catch (error) {
                            console.log("Error while fetching ordered products")
                        }
                }
              loadorderedProducts()
           async function getProfileDetail(){
            try {
                const response=await fetch("/glow-zaar/my-profile",{
                    credentials : 'include'
                })
                const data=await response.json()
                console.log("profile - ",data)
                setProfile(data)
            } catch (error) {
                console.log("profile error ")
            }
        }
        getProfileDetail()
      },[])
      

      if(!productItems)
      {
          return(
                <h2>Loading Ordered Products....</h2>
              );
          
      }
      else if(productItems && profile ){
              if(productItems.length==0){
                    console.log("it is running")
                    return (
                      <div>
                        <NavBar/>
                        <h2>Your Orders is empty , Order more Products</h2>
                      </div>
                    )
              }
              return (
                <>
                  <NavBar/>
                  <br></br><br></br><br></br>
                  <div className="myorders-main-container">
                  <div className="orders-section">
                      {productItems.map(product => (
                      <div key={product.id} className="order-card">
                          <div className="image-sec">
                          <img src={product.poster} alt={product.name} className="order-image" />
                          </div>
                          <div className="order-info">
                          <h3>{product.name}</h3>
                          <p className='p-tag'><strong >No. of items:</strong> {product.quantity}</p>
                          <p className='p-tag'><strong>Price:</strong> ₹{product.price}</p>
                          <p className='p-tag'><strong>🕒 Ordered On:</strong> <span>{product.orderedDate}</span></p>
                          <p className='delivery-date'>📅 <strong>Estimated Delivery:</strong>{product.deliveryDate}</p>
                          <div className="badge">✔ Ordered Successfully</div>
                          </div>
                      </div>
                      ))}
                  </div>

                  <div className="delivery-info-box">
                      <h2>🚚 Delivery Address</h2>
                      <p>📍 <strong>{profile.name}</strong></p>
                      <p>📍 <strong>{profile.address}</strong></p>
                      <p>📅 <strong>Email : </strong> {profile.email}</p>
                      <p>📦 <strong>Total Orders : </strong> {productItems.length}</p>
                  </div>
                  </div>
                </>
              );
          }
      

  
}

