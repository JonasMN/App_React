import React from 'react';
import './App.css';
import ListOfGifs from './components/ListOfGifs';
import {Link ,Route} from "wouter"



 function App() {
return (
 <div className="App">
  <section className="App-content">
  <Link  to='/gif/panda'>Gifs de pandas</Link> 
  <Link  to='/gif/ecuador'>Gifs de ecuador</Link> 
  <Link  to='/gif/venezuela'>Gifs de venezuela</Link> 
  <Link  to='/gif/mario'>Gifs de mario</Link>
  <Route   
    path="/gif/:keyword"
    component={ListOfGifs} />
  </section>
 </div>
 )
}

export default App;
