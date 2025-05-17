import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import {createBrowserRouter,RouterProvider} from 'react-router-dom'
import HomePage from './pages/Home.jsx'
import PageNotFound from './compartment/PageNotFound.jsx'
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
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={my_router}/>
  </StrictMode>,
)
