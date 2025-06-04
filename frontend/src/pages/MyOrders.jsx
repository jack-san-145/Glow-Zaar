import { useEffect, useState } from "react";


function MyOrders()
{
    const [order,setorder]=useState(null)
    useEffect(()=>{
        async function fetchingOrder(){
           const response=await fetch("/glow-zaar/MyOrders",{
            method:"GET",
            credentials : "include",
           })
            data=await response.json()
            setorder(data)
        }
        fetchingOrder()
    },order)

    if(!order){
        return (
            <div>
                <h2>Loading</h2>
            </div>
        )
    }
    else{
        return(
            <div>
                <h2>Myorder loaded</h2>
            </div>
        )
    }
}

export default MyOrders