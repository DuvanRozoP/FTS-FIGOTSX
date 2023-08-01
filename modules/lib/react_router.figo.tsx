import React from 'react';
import { Route, Routes } from 'react-router-dom';

import About from '@pages/About';
import Home from '@pages/Home';
import Index from '@pages/index';


function RouterApp() {
  return (
	<Routes>
		<Route path="/about" Component={About} />
		<Route path="/home" Component={Home} />
		<Route path="/" Component={Index} />
		
	</Routes>
  );
}

export default RouterApp;