import '../CSS/NavBar.css'
import { Link } from 'react-router-dom';

function NavBar(){
  return (
      <nav className="navbar">
        <div className="brand">Glow Zaar</div>
        <div className="nav-items">
          <Link to="/">Home</Link>
          <Link to="/MyCart">My Cart</Link>
          <Link to="something">Profile</Link>
          <Link to="/MyOrders">My Orders</Link>
          <div className="search-bar">
            <i className="fas fa-search" />
            <input type="text" placeholder="Search..." />
          </div>
        </div>
      </nav>

  );
}


export default NavBar