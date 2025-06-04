
import { useEffect, useState } from 'react';
import '../CSS/MyCart.css'
import NavBar from '../component/NavBar.jsx';
import Login from './Login.jsx';
import { useNavigate } from 'react-router-dom';
function MyCart(){

    const [isLogin,setLoginIn]=useState(null)

    useEffect(()=>{
        async function checkLogin(){
            try {
                const response = await fetch("/glow-zaar/MyCart", {
                                        credentials: "include"
                                            })

                const data =await response.json()
                setLoginIn(data)
            } catch (error) {
                console.log("Error while fetching Cookie Found details")
            }
        }
        checkLogin()
    }
    ,isLogin)
    if(!isLogin){
        return(
            <h2>Loading....</h2>
        );
    }
    else{
        if(isLogin.isFound==false){
            return(
                <div>
                    <NavBar/>
                    <Login/>
                </div>
            );
        }
        else{
            return(
                <h2>Now you can see your cart</h2>
            )
        }
    }
       
}

export default MyCart