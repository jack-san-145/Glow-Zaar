import ProductCard from "../component/Collection_ProductCard"
import '../CSS/Home.css'
import NavBar from "../component/NavBar"
import 'react-router-dom'
import { Link, useNavigate, useParams } from "react-router-dom"
import { useEffect, useRef, useState } from "react"
import ProductDetail from "./ProductDetails"

 

function ProductCollection(){

    const {product_type_id}=useParams()
    console.log("product_type_id - ",product_type_id)

    const [collection,setCollection]=useState([])
    let loadedOnce=useRef(false)
        useEffect(() => {
        async function LoadProducts() {
            try {
                if(!loadedOnce.current){
                    const response = await fetch(`http://localhost:8989/load-products/${product_type_id}`);
                    const data = await response.json();
                    setCollection(data); // this updates the state and causes a re-render
                    loadedOnce.current=true
                }
                
            } catch (err) {
                console.error("Failed to fetch products:", err);
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