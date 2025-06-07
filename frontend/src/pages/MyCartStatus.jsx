
import { useEffect, useState } from 'react';
import '../CSS/MyCart.css'
import NavBar from '../component/NavBar.jsx';
import Login from './Login.jsx';
import { useNavigate } from 'react-router-dom';
import jumkha_3 from '../assets/products/jumkha_3.jpg'
function MyCartStatus(){

const [isLogin,setLoginIn]=useState(null)
const navigate=useNavigate()

useEffect(()=>{
    async function checkLogin(){
        try {
            const response = await fetch("/glow-zaar/GetMyCardStatus", {
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
else
{
    if(isLogin.isFound==false){
        return(
            <div>
                <NavBar/>
                <Login/>
            </div>
        );
    }
    else
    {
      navigate("/MyCartProducts")
    }
}
       
}

export default MyCartStatus