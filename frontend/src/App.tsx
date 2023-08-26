import React from 'react'

import { BrowserRouter, Route, Routes } from 'react-router-dom'

import CardList from './components/CardList'
import FormLead from './components/FormLead'

function App() {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <Routes>
        <Route path='/' element={<CardList />} />
          <Route path='/contact' element={<FormLead />} />        
        </Routes>  
      </BrowserRouter>
    </React.StrictMode>
  )
  {/* <div className='App'>
      <CardList
        } />
    </div> */}
}

export default App
