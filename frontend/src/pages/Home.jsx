import ProductCard from '../component/ProductCard'
import baby_cloth from '../assets/products/baby_cloth.jpg'
import body_spray from '../assets/products/body_spray.jpg'
import cosmetics from '../assets/products/cosmetics.jpg'
import jumkha_2 from '../assets/products/jumkha_2.jpg'
import jumkha_3 from '../assets/products/jumkha_3.jpg'
import jumkha_4 from '../assets/products/jumkha_4.jpg'
import perfume from '../assets/products/perfume.jpg'
import shirt from '../assets/products/shirt.jpg'
import watch_1 from '../assets/products/watch_1.jpg'
import watch_2 from '../assets/products/watch_2.jpg'
import women_cloth from '../assets/products/women_cloth.jpg'
import '../CSS/Home.css'
import NavBar from '../component/NavBar'
import Login from './Login'
import Register from './Register'
import { useEffect, useState } from 'react'

function HomePage(){

    const [Products,setProducts]=useState(null)

    useEffect(()=>{
       async function loadHomeCard()
       {

        const response=await fetch("/glow-zaar/home",{
            credentials :"include"
        })
        const data =await response.json()
        setProducts(data)
        }
        loadHomeCard()
    },[])

    if(!Products){
        return(
            <div>
                <NavBar/>
                <h2>Loading Home Page</h2>
            </div>
        )
    }
    else{
            function cardClickHandler(id){
                console.log(id)
            }

         const ProductList=Products.map(
            (product) =>  <ProductCard Product={product} key={product.product_type_id} whenClicked={
                                        ()=>{cardClickHandler(product.product_type_id)}
                                    }/> 
            )

        return (
            <div className='home'>
                <NavBar/>
                <br/><br/><br/>
                <div className='product-grid'>
                    
                    {ProductList}
                </div>
            </div>
        );
    }
}

export default HomePage