import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_URL } from './defineUrl';

function Main() {
	const kibanaUrl: string = KIBANA_URL;

	return (
		<div className="embed-responsive embed-responsive-16by9"
		style={{position:"absolute", width:"100%" , height:"80%", left :"0px"}} >	
			<Iframe url={kibanaUrl} className="embed-responsive-item"  
	               height = "50%"
			scrolling="yes" >
			</Iframe>
		</div>
		
	);
}

export default Main;
