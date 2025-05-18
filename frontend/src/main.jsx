import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import {createBrowserRouter,RouterProvider} from 'react-router-dom'
import HomePage from './pages/Home.jsx'
import PageNotFound from './component/PageNotFound.jsx'
import ProductCollection from './pages/ProductCollection.jsx'
import ProductDetails from './pages/ProductDetails.jsx'



const my_router=createBrowserRouter([
  {
    path : '/',
    element : <HomePage/>,
    errorElement :<PageNotFound/>
  },
  {
    path : '/product-collection',
    element : <ProductCollection/>,
    errorElement : <PageNotFound/>
  },
  {
    path : '/product-details',
    element : <ProductDetails/>
  },
  {
    path : '/product-details/:pid',
    element : <ProductDetails/>,
    errorElement : <PageNotFound/> 
  }
])

createRoot(document.getElementById('root')).render(
    <RouterProvider router={my_router}/>
)
