import ProductCard from '../component/Home_ProductCard'
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

function HomePage(){

    const Products=[
        {product_type_id:"cloth",poster:baby_cloth,name:"Baby Cloth",price:599},
        {product_type_id:"Mobile",poster:body_spray,name:"Body Spray",price:459},
        {product_type_id:"Cosmetics",poster:cosmetics,name:"Cosmetics",price:999},
        {product_type_id:"jewels",poster:jumkha_2,name:"Jumkha",price:250},
        {product_type_id:"Fashion",poster:jumkha_3,name:"Jumkha",price:260},
        {product_type_id:"accessoriers",poster:jumkha_4,name:"Jumkha",price:299}
    ]

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

export default HomePage