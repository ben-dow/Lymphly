import { Routes, Route } from 'react-router'
import Search from '../../pages/search'

export default function Body(){
    return(
      <div className=''>
        <Routes>
          <Route path="/" element={<Search/>}/>
        </Routes>
      </div>
    )
}

