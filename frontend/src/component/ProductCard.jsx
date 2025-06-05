import { Link } from 'react-router-dom';
import '../CSS/ProductCard.css' 
import '../pages/ProductCollection.jsx'
function ProductCard({Product,whenClicked}){

    if(Product.isproduct)
    {
        return(
        
            <div className='product-row'>
                <div className="product-card" onClick={whenClicked}>
                    <Link to={`/product-collection/${Product.product_type_id}`}>
                        <img src={Product.poster} alt="product-img" className="product-image"/>
                        <div className="product-details">
                            <h3>{Product.name}</h3>
                            <p className="highlight">₹{Product.price}</p>
                        </div>
                    </Link>
                </div>
            </div>
         );
        
    }
    else{
        console.log("product from home page - ",Product)
        return(
        
            <div className='product-row'>
                <div className="product-card" onClick={whenClicked}>
                    <Link to={`/product-collection/${Product.product_type_id}`}>
                        <img src={Product.poster} alt="product-img" className="product-image"/>
                        <div className="product-details">
                            <h3>{Product.product_type_id}</h3>
                            {/* <p className="highlight">₹{Product.price}</p> */}
                        </div>
                    </Link>
                </div>
            </div>

        );
    }
    
    
}

export default ProductCard