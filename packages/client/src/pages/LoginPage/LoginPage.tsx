import { Button, TextField } from "@mui/material";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import ApiService from "../../services/ApiService";

const LoginPage = () => {
  const navigate = useNavigate();

  const [loginData, setLoginData] = useState({ Email: "", Password: "" });

  const handleLogin = async () => {
    const res = await ApiService.post<string>("/auth/login", { body: loginData });

    localStorage.setItem("accessToken", res.data || "");

    if (res.ok) navigate("/")
  };

  return (
    <div className="w-full min-h-screen h-screen flex justify-center items-center flex-col">
      <TextField
        label="Email"
        variant="outlined"
        value={loginData.Email}
        onChange={(e) => setLoginData({ ...loginData, Email: e.target.value })}
      />
      <TextField
        label="Password"
        variant="outlined"
        type="password"
        value={loginData.Password}
        onChange={(e) =>
          setLoginData({ ...loginData, Password: e.target.value })
        }
      />
      <Button onClick={handleLogin}>Login</Button>
      <div>
        Don't have account now?{" "}
        <Link to="/register" className="text-blue-500">
          Register now
        </Link>
      </div>
    </div>
  );
};

export default LoginPage;
