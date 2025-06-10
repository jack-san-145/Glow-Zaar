import '../CSS/ProductDetails.css';
import body_spray from '../assets/products/body_spray.jpg'
import ProductCard from '../component/Collection_ProductCard';
import NavBar from '../component/NavBar';
import jumkha_2 from '../assets/products/jumkha_2.jpg'
import cosmetics from '../assets/products/cosmetics.jpg'
import ProductCollection from './ProductCollection';
import { useLocation, useParams } from 'react-router-dom';
import { useEffect, useRef, useState } from 'react';

const cache_product_details={}
let quantity=0
const discountOnce={}
let ordered=false

function ProductDetail() {
  const [statusMessage, setStatusMessage] = useState("");
  
  async function buyNow(){
      try {
        const response=await fetch("/glow-zaar/buyNow",{
          method : "POST",
          headers: {
            'Content-Type': 'application/json'  
          },
          body : JSON.stringify({"pid":product.pid,"price":product.price,"quantity":itemCount})

        })
        const data=await response.json()
        console.log("data from buy now - ",data)
        setStatusMessage("Ordered Successfully")
        
      } catch (error) {
        console.log("error occured while fetching buy it now ",err)
      }
  }
  async function addToCart(){
      try {
        const response=await fetch("/glow-zaar/addToCart",{
          method : "POST",
          headers: {
            'Content-Type': 'application/json'  
          },
          body : JSON.stringify({"pid":product.pid,"price":product.price,"quantity":itemCount})

        })
        const data=await response.json()
        console.log("data from addToCart - ",data)
        setStatusMessage("Added to cart")
        
      } catch (error) {
        console.log("error occured while fetching add to cart",err)
      }
  }
  useEffect(() => {
  if (statusMessage) {
    const timer = setTimeout(() => setStatusMessage(""), 2000);
    return () => clearTimeout(timer);
  }
}, [statusMessage]);

  let product;
  const {pid,product_type_id}=useParams()
  console.log("pid value - ",pid)
  console.log("Product type - ",product_type_id)
  const [product_collection,setProduct]=useState([])
  const [itemCount,setItemCount]=useState(1)
  const[applyDiscount,setDiscount]=useState(0)
  let loadedOnce=useRef(false)

  useEffect(()=>{
    console.log("Use effect triggered - ",pid)
    console.log("use trigger - ",product_type_id)
      async function fetching(){
      console.log("fetching calling")

      if(cache_product_details[product_type_id]){
        setProduct(cache_product_details[product_type_id])
      }
      else{
          try {
            const response=await fetch(`/glow-zaar/load-products/${product_type_id}`)
            const data=await response.json()
            cache_product_details[product_type_id]=data
            setProduct(data)
          } catch (error) {
            console.log("Failed to fetch products : ",error)
          }
      }

          
    }
      fetching()

    },[])
  
  if(product_collection.length==0){
    return(
      <h2>loading....</h2>
    );
  }

let filtered_products=[];
function parsing(){
  let count=0
  for(let i=product_collection.length-1;i>=0;i--){
    if(pid==product_collection[i].pid)
    {
      product=product_collection[i]
    }

      if([0,4,7,9].includes(i)){
        filtered_products.push(product_collection[i])
        count+=1
      }
  }
}
parsing()

console.log("product",product)
console.log("filtered_products",filtered_products)

const SimilarProduct=filtered_products.map(
  (selected_product) => <ProductCard Product={selected_product} key={selected_product.pid} 
                                     />
    
)


return (
<>  
    <NavBar/>
    <br/><br/><br/><br/><br/>
    <div className="product-detail-container">
      <div className="image-section">
        <img src={product.poster} alt={product.poster} />
      </div>

      <div className="info-section">
        <p className="sku">{product.sku}</p>
        <h1 className="product-name">{product.name}</h1>

        <div className="price-section">
          {product.originalPrice && (
            <span className="original-price">₹{product.originalPrice}</span>
          )}
          <span className="current-price">₹{product.price}</span>
          {product.sale && <span className="sale-tag">Sale</span>}
        </div>

        <div className="quantity-section">
          <button onClick={()=>{setItemCount(itemCount-1);console.log(itemCount)}}>-</button>
          {
            itemCount<=1 ?(<span>1</span>):(<span>{itemCount}</span>)
          }
          <button onClick={()=>{
            let temp=itemCount
             if(temp+1>product.quantity)
            {
              setItemCount(itemCount)
            }else{
              setItemCount(itemCount+1)
            }
            console.log(itemCount)}}>+</button>
        </div>
        <div className='action-buttons'>
            <button className='discount' onClick={()=>{
              console.log("discount clicking")
              if(!discountOnce[product.pid]){
                let discount=product.discount
                let discountAmount=(discount/100)*product.price
                product.price=Math.round(product.price-discountAmount)
                console.log("discount - ",discount)
                console.log("DiscountAmount - ",discountAmount)
                console.log("product price - ",product.price)
                setDiscount(product.price)
                discountOnce[product.pid]=true
              }
              
            }}>Apply Discount  ({product.discount}%) </button>
        </div>          
        <div className="action-buttons">
        
          <button className="add-to-cart" onClick={()=>{addToCart()}} >Add to cart</button>
          <button className="buy-now" onClick={()=>{buyNow()}}>Buy it now</button>
        </div>
          {statusMessage && <h2 className='status'>{statusMessage}</h2>}
          
        <div className="product-specs">
         <h2>Specifications</h2>
          <p><strong>Product Name:</strong> {product.name}</p>
          <p><strong>Brand:</strong> {product.brand}</p>
          <p><strong>Category:</strong> {product.category}</p>
          <p><strong>Color:</strong> {product.color}</p>
          <p><strong>Material:</strong> {product.material}</p>
          <p><strong>Weight:</strong> {product.weight}</p>
          <p><strong>Size (L×B×H):</strong> {product.size}</p>
        </div>
      </div>
    </div>
    <hr/>
    <div className='recommends'>
        <h2>Similar Products</h2>
        <div className='product-grid'>
        
        {SimilarProduct}

        </div>
    </div>
</>
  );
}

export default ProductDetail;
