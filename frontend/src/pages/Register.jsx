import React, { useState } from 'react';
import "../CSS/Auth.css";
import NavBar from '../component/NavBar';

function Register() {
  const [formData, setFormData] = useState({
    user_name: '',
    user_email: '',
    user_password: '',
    user_age: '',
    user_address: ''
  });

  const [message, setMessage] = useState(""); // to show success or error

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const encoded = new URLSearchParams(formData);
      const response = await fetch("http://localhost:8989/glow-zaar/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: encoded.toString()
      });

      const result = await response.json();

      if (response.ok) {
        setMessage(result.message);
      } 
    } catch (err) {
      setMessage("❌ Something went wrong");
    }
  };

  return (
    <>
        <div className="auth-container">
          <NavBar />
          <div className="auth-box">
            <div className="auth-header">Register</div>

            <form onSubmit={handleSubmit}>
              <div className="auth-content">
                <input className="auth-input" type="text" name="user_name" placeholder="Name" onChange={handleChange} required />
                <input className="auth-input" type="email" name="user_email" placeholder="Email" onChange={handleChange} required />
                <input className="auth-input" type="password" name="user_password" placeholder="Password" onChange={handleChange} required />
                <input className="auth-input" type="text" name="user_age" placeholder="Age" onChange={handleChange} required />
                <textarea className="auth-input" name="user_address" placeholder="Address" onChange={handleChange} rows="3" required></textarea>

                <button className="auth-button" type="submit">Register</button>

                <div className="auth-footer">
                  Already have an account? <a href='/login'><span>Login</span></a>
                </div>
              </div>
            </form>
            {message && <div className="auth-message">{message}</div>}
          </div>
          
        </div>
        
    </>
  );
}

export default Register;
