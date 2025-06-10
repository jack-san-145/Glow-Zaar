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
import { MyOrderStatus } from './pages/MyOrders.jsx'
import { MyCartStatus } from './pages/MyCartStatus.jsx'; 
import { ProfileStatus } from './pages/ProfilePage.jsx'



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
    path:'/myCartStatus',
    element :<MyCartStatus/>,
    errorElement :<PageNotFound/>
  },
  {
    path : '/myOrderStatus',
    element : <MyOrderStatus/>,
    errorElement : <PageNotFound/>
  },
  {
    path : 'myprofileStatus',
    element : <ProfileStatus/>,
    errorElement : <PageNotFound/>
  }
])

createRoot(document.getElementById('root')).render(
    <RouterProvider router={my_router}/>
)
