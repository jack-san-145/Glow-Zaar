import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import {createBrowserRouter,RouterProvider} from 'react-router-dom'
import HomePage from './pages/Home.jsx'
import Register from './pages/Register.jsx'
import Login from './pages/Login.jsx'
import PageNotFound from './component/PageNotFound.jsx'
import ProductCollection from './pages/ProductCollection.jsx'
import ProductDetails from './pages/ProductDetails.jsx'
import MyOrders from './pages/MyOrders.jsx'
import MyCartProducts from './pages/MyCartProducts.jsx'
import MyCartStatus from './pages/MyCartStatus.jsx'




const my_router=createBrowserRouter([
  {
    path : '/',
    element : <HomePage/>,
    errorElement :<PageNotFound/>
  },
  {
    path : '/product-collection/:product_type_id',
    element : <ProductCollection/>,
    errorElement : <PageNotFound/>
  },
  {
    path : '/product-details',
    element : <ProductDetails/>
  },
  {
    path : '/product-details/:pid/:product_type_id',
    element : <ProductDetails/>,
    errorElement : <PageNotFound/> 
  },
  {
    path:'/register',
    element : <Register/>,
    errorElement :<PageNotFound/>
  },
  {
    path:'/login',
    element : <Login/>,
    errorElement :<PageNotFound/>
  },
  {
    path:'/MyOrders',
    element : <MyOrders/>,
    errorElement :<PageNotFound/>
  },
  {
    path:'/MyCartProducts',
    element :<MyCartProducts/>,
    errorElement :<PageNotFound/>
  },
  {
    path:'/myCartStatus',
    element :<MyCartStatus/>,
    errorElement :<PageNotFound/>
  }
])

createRoot(document.getElementById('root')).render(
    <RouterProvider router={my_router}/>
)
