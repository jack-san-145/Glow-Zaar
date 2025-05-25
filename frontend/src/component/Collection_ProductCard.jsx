import { Link } from 'react-router-dom';
import '../CSS/ProductCard.css' 



function ProductCard({Product}){
    return(
        
    <div className='product-row'>
        <div className="product-card">
            <Link to={`/product-details/${Product.pid}/${Product.product_type_id}`} >
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

export default ProductCard