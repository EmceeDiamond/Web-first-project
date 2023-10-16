import {
  BrowserRouter,
  Routes,
  Route
} from 'react-router-dom'

import Home from "./pages/Home"
import AddEmployee from './pages/AddEmployee'
import UpdateEmployee from './pages/UpdateEmployee'
import GetProduct  from  './pages/GetProducts'
import AddProduct  from  './pages/AddProduct'

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home/>} />
          <Route path='/add' element={<AddEmployee />} />
          <Route path='/update/:id' element={ <UpdateEmployee /> } />
          <Route path='/get_product' element={ <GetProduct /> } />
          <Route path='/add_product' element={ <AddProduct /> } />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App