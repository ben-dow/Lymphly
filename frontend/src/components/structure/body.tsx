import { Routes, Route, Navigate } from 'react-router'
import Search from '../../pages/search/search'

export default function Body(){
    return(
      <div className='h-full w-full'>
        <Routes>
        <Route index element={ <Navigate to="/search" /> } />
          <Route path="/search/*" element={<Search/>}/>
        </Routes>
      </div>
    )
}

