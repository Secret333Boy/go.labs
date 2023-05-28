import { Button, TextField } from "@mui/material";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import ApiService from "../../services/ApiService";

const RegisterPage = () => {
  const navigate = useNavigate();

  const [registerData, setRegisterData] = useState({ Email: "", Password: "" });

  const handleRegister = async () => {
    const res = await ApiService.post<string>("/auth/register", { body: registerData });

    localStorage.setItem("accessToken", res.data || "");

    if (res.ok) navigate("/");
  };

  return (
    <div className="w-full min-h-screen h-screen flex justify-center items-center flex-col">
      <TextField
        label="Email"
        variant="outlined"
        value={registerData.Email}
        onChange={(e) =>
          setRegisterData({ ...registerData, Email: e.target.value })
        }
      />
      <TextField
        label="Password"
        variant="outlined"
        type="password"
        value={registerData.Password}
        onChange={(e) =>
          setRegisterData({ ...registerData, Password: e.target.value })
        }
      />
      <Button onClick={handleRegister}>Register</Button>
      <div>
        Already have an account?{" "}
        <Link to="/login" className="text-blue-500">
          Login now
        </Link>
      </div>
    </div>
  );
};

export default RegisterPage;
