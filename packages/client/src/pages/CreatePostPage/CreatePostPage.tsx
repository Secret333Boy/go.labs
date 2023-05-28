import { Button, TextField } from "@mui/material";
import React, { useState } from "react";
import ApiService from "../../services/ApiService";
import { useNavigate } from "react-router-dom";

const CreatePostPage = () => {
  const navigate = useNavigate();

  const [createPostData, setCreatePostData] = useState({
    Title: "",
    Description: "",
  });

  const handleCreatePost = async () => {
    await ApiService.post("/posts", { body: createPostData });
    navigate("/posts");
  };

  return (
    <div className="w-full min-h-screen h-screen flex justify-center items-center flex-col">
      <TextField
        label="Title"
        variant="outlined"
        value={createPostData.Title}
        onChange={(e) =>
          setCreatePostData({ ...createPostData, Title: e.target.value })
        }
      />
      <TextField
        label="Description"
        variant="outlined"
        value={createPostData.Description}
        onChange={(e) =>
          setCreatePostData({ ...createPostData, Description: e.target.value })
        }
      />
      <Button onClick={handleCreatePost}>Create a post</Button>
    </div>
  );
};

export default CreatePostPage;
