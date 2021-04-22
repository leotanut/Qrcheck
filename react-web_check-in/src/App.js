import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
 
import Home from './components/Home';
import CheckIn from './components/CheckIn';
import CheckOut from './components/CheckOut';
import Navigation from './components/Navigation';
//import Error from './components/Error';
 
class App extends Component {
  render() {
    return (      
       <BrowserRouter>

        <div>
            
             <Route path="/" component={Home} exact/>
             <Route path="/checkin" component={CheckIn}/>
             <Route path="/checkout" component={CheckOut}/>
          
        </div> 
      </BrowserRouter>
    );
  }
}
 
export default App;