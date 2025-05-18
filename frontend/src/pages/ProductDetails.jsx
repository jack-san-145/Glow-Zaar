import '../CSS/ProductDetails.css';
import body_spray from '../assets/products/body_spray.jpg'
import ProductCard from '../component/Collection_ProductCard';
import NavBar from '../component/NavBar';
import jumkha_2 from '../assets/products/jumkha_2.jpg'
import cosmetics from '../assets/products/cosmetics.jpg'
import ProductCollection from './ProductCollection';
import { useLocation, useParams } from 'react-router-dom';

 let default_product = {
    poster: jumkha_2,
    sku: "SKU12345",
    name: "Fan Air Cooler USB Electric Fan",
    brand: "Generic",
    category: "Electronics",
    color: "White",
    material: "Plastic",
    weight: "120 grams",
    size: "29D x 21W x 26H cm",
    originalPrice: 867,
    price: 1178.82,
    sale: true,
    discount:20,
    ordered:true,
    addToCart:true
  };

function ProductDetail() {
  const {state}=useLocation()
  let product=state?.clickedProduct
  console.log(product)
  if(!product){
    product=default_product
  }

  const {pid}=useParams()
  // const product=

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
          <button>-</button>
          <span>1</span>
          <button>+</button>
        </div>
        <div className='action-buttons'>
            <button className='discount'>Apply Discount  ({product.discount}%) </button>
        </div>          
        <div className="action-buttons">
        
          <button className="add-to-cart" >Add to cart</button>
          <button className="buy-now" >Buy it now</button>
        </div>
          {product.ordered ?(<h2 className='status'>Ordered Successfully</h2>)
                            :product.addToCart ?(<h2 className='status'>Added to cart</h2>)
                            :(<h2></h2>)
          }
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
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        {/* <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/>
        <ProductCard Product={{pid:3,poster:cosmetics,name:"Cosmetics",price:999}}/> */}
        

        </div>
    </div>
</>
  );
}

export default ProductDetail;
