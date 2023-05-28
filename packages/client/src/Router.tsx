import React, { FC, ReactNode, useEffect, useState } from "react";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage/HomePage";
import ToolbarLayout from "./components/ToolbarLayout/ToolbarLayout";
import PostsPage from "./pages/PostsPage/PostsPage";
import SettingsPage from "./pages/SettingsPage/SettingsPage";
import LoginPage from "./pages/LoginPage/LoginPage";
import RegisterPage from "./pages/RegisterPage/RegisterPage";
import ApiService from "./services/ApiService";
import AccountContext from "./context/AccountContext";
import { Account } from "./types";
import CreatePostPage from "./pages/CreatePostPage/CreatePostPage";

interface ProtectedProps {
  children: ReactNode;
}

const Protected: FC<ProtectedProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoaded, setIsLoaded] = useState(false);
  const [account, setAccount] = useState<Account>();

  useEffect(() => {
    (async () => {
      const { ok } = await ApiService.get("/auth/validate");

      if (!ok) {
        const { data, ok: refreshOk } = await ApiService.get<{
          accessToken: string;
        }>("/auth/refresh");
        if (!data || !refreshOk) return setIsLoaded(true);
        localStorage.setItem("accessToken", data.accessToken);
      }

      const { data } = await ApiService.get<Account>("/accounts");
      setAccount(data);
      setIsAuthenticated(true);
      setIsLoaded(true);
    })();
  }, []);

  if (!isLoaded) return <></>;

  if (!isAuthenticated || !account) {
    return <Navigate to="/login" />;
  }

  return (
    <AccountContext.Provider value={account}>
      {children}
    </AccountContext.Provider>
  );
};

const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <Protected>
              <ToolbarLayout>
                <HomePage />
              </ToolbarLayout>
            </Protected>
          }
        />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route
          path="/posts"
          element={
            <Protected>
              <ToolbarLayout>
                <PostsPage />
              </ToolbarLayout>
            </Protected>
          }
        />
        <Route
          path="/posts/create"
          element={
            <Protected>
              <ToolbarLayout>
                <CreatePostPage />
              </ToolbarLayout>
            </Protected>
          }
        />
        <Route
          path="/posts/:id"
          element={
            <Protected>
              <ToolbarLayout>Post</ToolbarLayout>
            </Protected>
          }
        />
        <Route
          path="/settings"
          element={
            <Protected>
              <ToolbarLayout>
                <SettingsPage />
              </ToolbarLayout>
            </Protected>
          }
        />

        <Route
          path="*"
          element={<ToolbarLayout>404 Not found page</ToolbarLayout>}
        />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
