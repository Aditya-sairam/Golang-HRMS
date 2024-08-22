import React, { useState } from 'react';
import axios from 'axios';
import './styling/signup.css'


const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleEmail = (e) => {
    setEmail(e.target.value)
  }

  const handlePassword = (e) => {
    setPassword(e.target.value)
  }

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await axios.post('http://localhost:9000/users/login', {
        email,       
        password,
      });
      const {token} = response.token
      if(token){
        localStorage.setItem('token',token)
      }
      // Handle successful login (e.g., save token, redirect)
      console.log(response.data);
    } catch (error) {
      // Handle error (e.g., display error message)
      setError('Login failed. Please check your email and password.');
    }
  };

  return (
    <div>
      
      <form onSubmit={handleSubmit} className='form-container'>
        <h2>Login</h2>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={handleEmail}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={handlePassword}
            required
          />
        </div>
        <button type="submit">Login</button>
        {error && <p>{error}</p>}
      </form>
    </div>
  );
};

export default Login;
