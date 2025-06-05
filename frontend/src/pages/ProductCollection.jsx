import ProductCard from "../component/Collection_ProductCard"
import '../CSS/Home.css'
import NavBar from "../component/NavBar"
import 'react-router-dom'
import { Link, useNavigate, useParams } from "react-router-dom"
import { useEffect, useRef, useState } from "react"
import ProductDetail from "./ProductDetails"

const cache_product={}

function ProductCollection(){

    const {product_type_id}=useParams()
    console.log("product_type_id - ",product_type_id)

    // if(product_type_id!=global_ProductTypeId){
    //     loadedOnce=false

    // }else{
    //     global_ProductTypeId=product_type_id
    // }

    const [collection,setCollection]=useState([])
        useEffect(() => {
        async function LoadProducts() {
            if(cache_product[product_type_id]){
                setCollection(cache_product[product_type_id])
            }else
            {

                try {
                        console.log("running")
                        const response = await fetch(`/glow-zaar/load-products/${product_type_id}`);
                        const data = await response.json();
                        cache_product[product_type_id]=data
                        setCollection(data); // this updates the state and causes a re-render
                        
                } catch (err) {
                    console.error("Failed to fetch products:", err);
                }
                
            }
            

        }

        LoadProducts();
    }, []);

    
    const navigate=useNavigate()


    const categoryCollection=collection.map(
        (selected_product) => <ProductCard Product={selected_product} key={selected_product.pid} 
                                     />
    )
    return (
        <>
            <NavBar/>
            <br></br><br></br><br></br>
            <div className="product-grid">
                    {categoryCollection}
            </div>
        </>
    )
}



export default ProductCollection