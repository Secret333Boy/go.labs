import { AppBar, Toolbar, IconButton, Typography, Button } from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import React, { FC, ReactNode, useState } from "react";
import AppDrawer from "../AppDrawer/AppDrawer";
import { useLocation, useNavigate } from "react-router-dom";

interface ToolbarLayoutProps {
  children: ReactNode;
}

const ToolbarLayout: FC<ToolbarLayoutProps> = ({ children }) => {
  const location = useLocation();

  const [isDrawerOpen, setIsDrawerOpen] = useState(false);

  return (
    <>
      <div>
        <AppDrawer
          isOpen={isDrawerOpen}
          onClose={() => setIsDrawerOpen(false)}
        />
        <AppBar>
          <Toolbar>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
              onClick={() => setIsDrawerOpen(true)}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              {location.pathname}
            </Typography>
          </Toolbar>
        </AppBar>
      </div>
      <div className="pt-[64px]">{children}</div>
    </>
  );
};

export default ToolbarLayout;
