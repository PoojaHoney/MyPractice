import React, { useState } from 'react';
import logo from './logo.svg';
import { Input, Button } from '@ui5/webcomponents-react';
import './App.css';
import axios from 'axios';

function App() {
  const [sValue, setInputValue] = useState<string>()
  return (
    <div className="App">
      <Input value={sValue}></Input>
      <Button onClick={((oEvent) => {
        debugger;
        axios.get("/users/me/", {
          headers : {
            Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJwb29qYSIsImV4cCI6MTY0ODgyNDE2M30.5Z-2Eygcvy4z2fcaxgYkAg76aPjLhSKT1gtq9VteyPo`
          } 
        }).then((oResponse: any) => {
          debugger
        });
      })} >Click me </Button>
    </div>
  );
}

export default App;
