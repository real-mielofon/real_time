import React from 'react';
import ReactDom from "react-dom";

import 'jquery';
import 'bootstrap/dist/js/bootstrap';


import App from './components/App.jsx';

ReactDom.render(
    <App />, 
    document.getElementById('root')
    )