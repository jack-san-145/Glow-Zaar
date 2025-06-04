import React, { useState } from 'react';
import '../CSS/Auth.css';
import NavBar from '../component/NavBar';
import { useNavigate } from 'react-router-dom';


function Login() {
  const [formData, setFormData] = useState({
    user_email: '',
    user_password: ''
  });

  const [message, setMessage] = useState("");
  const navigate = useNavigate(); 
  const [status, setStatus] = useState(null); 

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const encoded = new URLSearchParams(formData);
      const response = await fetch("/glow-zaar/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: encoded.toString()
      });

      const result = await response.json();

      if (response.ok) {
        setMessage(result.message);
        setStatus(result.status);
          setTimeout(() => {
            if(result.status=="true"){
              navigate("/"); 
            }
            else{
              navigate("/login")
            }
          }, 2000);
      } 
    } catch (err) {
      setMessage("❌ Something went wrong");
    }
  };

  return (
    <>
      <NavBar />
      <div className="auth-container">
        <div className="auth-box">
          <div className="auth-header">Login</div>
          <form onSubmit={handleSubmit}>
            <div className="auth-content">
              <input
                className="auth-input"
                type="email"
                name="user_email"
                placeholder="Email"
                value={formData.user_email}
                onChange={handleChange}
                required
              />
              <input
                className="auth-input"
                type="password"
                name="user_password"
                placeholder="Password"
                value={formData.user_password}
                onChange={handleChange}
                required
              />
              <button className="auth-button" type="submit">Login</button>
              <div className="auth-footer">
                Don't have an account? <span><a href="/register">Register</a></span>
              </div>
            </div>
          </form>

          {/* ✅ Show message inside box */}
        {/* ✅ Show message inside box */}
          {status && (
            <div className={`auth-message ${status === 'true' ? 'success' : 'error'}`}>
              {message}
            </div>
          )}


        </div>
      </div>
    </>
  );
}

export default Login;
