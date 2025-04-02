import { Box, Button, Paper, TextField, Typography } from "@mui/material";
import React, { useState, useEffect, useMemo, useCallback } from "react";
import { loadSlim } from "@tsparticles/slim";
import Particles, { initParticlesEngine } from "@tsparticles/react";
import { useNavigate } from 'react-router-dom';
import {particleFunc} from './particleOptions'

function Login() {
  const [init, setInit] = useState(false);
  const [name, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [language, setLanguage] = useState('');
  const [dob, setDob] = useState('');
  const [loggedIn, setLoggedIn] = useState(false)
  const [passwordError, setPasswordError] = useState(false);
  const [passwordErrorDesc, setPasswordErrorDesc] = useState('');
  const [signupError, setSignupError] = useState('')
  const [signupErrorDesc, setSignupErrorDesc] = useState('')

  const handleSignIn = async (type, e) => {
    e.preventDefault(); // Prevent form from submitting normally

    try {
      if (type === "signup" && 
        (name === '' || password === '' || dob === '' || language === '')) {
          setSignupError(true)
          setSignupErrorDesc("Please fill out all fields")
          throw new Error('Please fill out all fields');
      }
      console.log(JSON.stringify({ name, password, language, dob }));
      const response = await fetch(`http://localhost:9888/user/${type}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, password, language, dob }),
      });

      if (!response.ok) {
        throw new Error('Request failed');
      }

      const data = await response.json();

      if (data.error) {
        throw new Error(data.status);
      }
      console.log('Login successful:', data.data);
      setLoggedIn(true)
    } catch (error) {
      console.error('Login error:', error);
      setPasswordError(true);
      setPasswordErrorDesc("Invalid username or password");
    }
  };

  const navigate = useNavigate();

  const goToChat = useCallback(() => {
    navigate('/movie');
  }, [navigate]);

  // this should be run only once per application lifetime
  useEffect(() => {
    initParticlesEngine(async (engine) => {
      await loadSlim(engine);
    }).then(() => {
      setInit(true);
    });

    if (loggedIn) {
      goToChat();
    }
  }, [goToChat, loggedIn]);

  const options = useMemo(
    particleFunc, []
  );

  useEffect(() => {
    if (loggedIn) {
      goToChat();
    }
  }, [goToChat, loggedIn]);

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleLanguageChange = (event) => {
    setLanguage(event.target.value)
  }

  const handleDobChange = (event) => {
    setDob(event.target.value)
  }

  return (
    <>
      <Box display="flex" gap={3} marginTop={5} justifyContent="space-around">
        <Button variant="contained" color='primary' onClick={goToChat}>Go to Chat</Button>
      </Box>
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        {init && <Particles id="tsparticles" options={options} />}
        <Paper
          style={{
            display: "flex",
            flexDirection: "column",
            gap: 20,
            padding: 50,
            zIndex: 99,
          }}
        >
          <Typography variant="h4" fontWeight="bold" marginBottom={2}>
            Welcome to MovieDB
          </Typography>
          <TextField
            label="Username"
            variant="filled"
            placeholder="Enter username here.."
            value={name}
            onChange={handleUsernameChange}
            error={signupError}
            helperText={signupErrorDesc}
          />
          <TextField
            label="Password"
            type="password"
            variant="filled"
            placeholder="Enter password here.."
            value={password}
            onChange={handlePasswordChange}
            error={passwordError}
            helperText={passwordErrorDesc}
          />
          <TextField
            label="Language"
            variant="filled"
            placeholder="Enter language here.."
            value={language}
            onChange={handleLanguageChange}
            error={signupError}
            helperText={signupErrorDesc}
          />
          <TextField
            label="Date of Birth (YYYY-MM-DD)"
            variant="filled"
            placeholder="Enter DOB here.."
            value={dob}
            onChange={handleDobChange}
            error={signupError}
            helperText={signupErrorDesc}
          />
          <Box display="flex" gap={3} marginTop={5} justifyContent="space-around">
            <Button variant="contained" color='primary' onClick={(e) => handleSignIn("signin", e)}>Login</Button>
            <Button variant="contained" color='secondary' onClick={(e) => handleSignIn("signup", e)}>Sign up</Button>
          </Box>
        </Paper>
      </Box>
    </>
  );
}

export default Login;
