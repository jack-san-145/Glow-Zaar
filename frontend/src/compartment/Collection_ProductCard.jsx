import { Link } from 'react-router-dom';
import '../CSS/ProductCard.css' 



function ProductCard({Product,whenClicked}){
    return(
        
    <div className='product-row'>
        <div className="product-card" onClick={whenClicked}>
            <Link to='/product-details' >
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