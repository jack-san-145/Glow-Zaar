import ProductCard from "../compartment/Collection_ProductCard"

import jumkha_2 from '../assets/products/jumkha_2.jpg'
import jumkha_3 from '../assets/products/jumkha_3.jpg'
import jumkha_4 from '../assets/products/jumkha_4.jpg'
import '../CSS/Home.css'
import NavBar from "../compartment/NavBar"
import 'react-router-dom'
import { Link } from "react-router-dom"
import { useEffect, useState } from "react"


function ProductCollection(){

    const [collection,setCollection]=useState([])
   
    useEffect(() => {
        async function LoadProducts() {
            try {
                const response = await fetch("http://localhost:8989/load-cosmetics/");
                const data = await response.json();
                setCollection(data); // this updates the state and causes a re-render
            } catch (err) {
                console.error("Failed to fetch products:", err);
            }
        }

        LoadProducts();
    }, []);

    const categoryCollection=collection.map(
        (selected_product) => <ProductCard Product={selected_product}/>
    )
    return (
        <>
            <NavBar/>
            <div className="product-grid">
                    {categoryCollection}
            </div>
        </>
    )
}



export default ProductCollection