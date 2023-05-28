import {
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import React, { FC, ReactNode } from "react";
import HomeIcon from "@mui/icons-material/Home";
import ChatIcon from "@mui/icons-material/Chat";
import SettingsIcon from "@mui/icons-material/Settings";
import { useNavigate } from "react-router-dom";

interface DrawerProps {
  isOpen: boolean;
  onClose: () => void;
}

const appDrawerFixtures: { text: string; icon: ReactNode; link: string }[] = [
  { text: "Home", icon: <HomeIcon />, link: "/" },
  { text: "Posts", icon: <ChatIcon />, link: "/posts" },
  { text: "Settings", icon: <SettingsIcon />, link: "/settings" },
];

const AppDrawer: FC<DrawerProps> = ({ isOpen, onClose }) => {
  const navigate = useNavigate();

  return (
    <Drawer anchor="left" open={isOpen} onClose={onClose}>
      <List
        sx={{
          width: 250,
        }}
      >
        {appDrawerFixtures.map((drawerItem) => (
          <ListItem key={drawerItem.text} disablePadding>
            <ListItemButton
              onClick={() => {
                navigate(drawerItem.link);
                onClose();
              }}
            >
              <ListItemIcon>{drawerItem.icon}</ListItemIcon>
              <ListItemText primary={drawerItem.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
};

export default AppDrawer;
