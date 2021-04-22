import React from 'react';
import {Card} from "react-bootstrap";
import {Button} from "react-bootstrap";
import { Route } from 'react-router-dom/cjs/react-router-dom.min';
import { NavLink } from 'react-router-dom';
const Home = () => {
    
    return (

        <div className="center">   
       <NavLink to="/checkin"><Button variant="primary" size="lg">
          เข้า
        </Button></NavLink> 
        <NavLink to="/checkout"><Button variant="secondary" size="lg">
         ออก
        </Button></NavLink> 
        </div>
    );
}
 
export default Home;