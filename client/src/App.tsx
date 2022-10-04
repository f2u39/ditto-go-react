import React from 'react';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { Route, Routes } from 'react-router-dom';

import Nav from './Components/Nav/Nav';
import SignIn from './Components/SignIn/SignIn';
import Act from './Components/Act/Act';
import Word from './Components/Word/Word';
import Game from './Components/Game/Game';
import Stopwatch from './Components/Stopwatch/Stopwatch';
import useToken from './useToken';

import './App.css';
import { Box, Grid } from '@mui/material';



export default function App() {
  // const { token, setToken } = useToken();

  // if (!token) {
  //   return <SignIn setToken={setToken} />
  // }

  return (
    // <ThemeProvider theme={darkTheme}>
    //   <CssBaseline />
    //   <div className="App">
    //     <Grid xs={12} sx={{ height: 80 }}>
    //       <Nav />
    //     </Grid>

    //     <BrowserRouter>
          <Routes>
            {/* <Route path="/" element={<Home />} /> */}
            <Route path="/act" element={<Act />} />
            <Route path="/word" element={<Word />} />
            <Route path="/game" element={<Game />} />
          </Routes>
    //     </BrowserRouter>
    //   </div>
    // </ThemeProvider>
  );
}
